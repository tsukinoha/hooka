package adaptive_card

import (
	"encoding/json"
	"strconv"
)

type (
	Column struct {
		version   float64
		Type      string    `json:"type"`
		Id        string    `json:"id,omitempty"`
		Items     []Element `json:"items"`
		Separator bool      `json:"separator,omitempty"`
		Spacing   string    `json:"spacing,omitempty"`
		Style     string    `json:"style,omitempty"`
		width     string
	}
	ColumnStr struct {
		Column
		Width string `json:"width,omitempty"`
	}
	ColumnInt struct {
		Column
		Width int `json:"width,omitempty"`
	}
)

func NewColumn() *Column {
	c := &Column{
		version:   1.0,
		Type:      "Column",
		Id:        "",
		Items:     []Element{},
		Separator: false,
		Spacing:   "",
		Style:     "",
		width:     "",
	}
	return c
}

func (c *Column) GetVersion() float64 {
	return c.version
}

func (c *Column) GetType() string {
	return c.Type
}

func (c *Column) SetId(id string) {
	c.Id = id
}

func (c *Column) Append(item Element) {
	c.Items = append(c.Items, item)
	if item.GetVersion() > c.GetVersion() {
		c.version = item.GetVersion()
	}
}

func (c *Column) SetSeparator(separator bool) {
	c.Separator = separator
}

func (c *Column) SetSpacing(spacing string) {
	c.Spacing = normalize(spacing, spacingValues)
}

func (c *Column) SetStyle(style string) {
	c.Style = normalize(style, containerStyleValues)
}

func (c *Column) SetWidth(width string) {
	widthInt, err := strconv.Atoi(width)
	if err != nil {
		c.width = normalize(width, columnWidthValues)
	} else {
		if widthInt > 0 {
			c.width = width
		} else {
			c.width = "0"
		}
	}
}

func (c *Column) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error
	var widthInt int
	widthInt, err = strconv.Atoi(c.width)
	if err != nil {
		col := ColumnStr{Column: *c, Width: c.width}
		data, err = json.Marshal(col)
	} else {
		col := ColumnInt{Column: *c, Width: widthInt}
		data, err = json.Marshal(col)
	}
	return data, err
}
