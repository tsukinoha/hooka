package adaptive_card

type (
	Container struct {
		version   float64
		Type      string    `json:"type"`
		Id        string    `json:"id,omitempty"`
		Items     []Element `json:"items"`
		Style     string    `json:"style,omitempty"`
		Separator bool      `json:"separator,omitempty"`
		Spacing   string    `json:"spacing,omitempty"`
	}
)

func NewContainer() *Container {
	c := &Container{
		version:   1.0,
		Type:      "Container",
		Id:        "",
		Items:     []Element{},
		Style:     "",
		Separator: false,
		Spacing:   "",
	}
	return c
}

func (c *Container) GetVersion() float64 {
	return c.version
}

func (c *Container) GetType() string {
	return c.Type
}

func (c *Container) SetId(id string) {
	c.Id = id
}

func (c *Container) Append(item Element) {
	c.Items = append(c.Items, item)
	if item.GetVersion() > c.GetVersion() {
		c.version = item.GetVersion()
	}
}

func (c *Container) SetStyle(style string) {
	c.Style = normalize(style, containerStyleValues)
}

func (c *Container) SetSeparator(separator bool) {
	c.Separator = separator
}

func (c *Container) SetSpacing(spacing string) {
	c.Spacing = normalize(spacing, spacingValues)
}
