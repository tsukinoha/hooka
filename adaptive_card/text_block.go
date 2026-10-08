package adaptive_card

type (
	TextBlock struct {
		version             float64
		Type                string `json:"type"`
		Text                string `json:"text"`
		Color               string `json:"color,omitempty"`
		HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
		Subtle              bool   `json:"isSubtle,omitempty"`
		MaxLines            int    `json:"maxLines,omitempty"`
		Size                string `json:"size,omitempty"`
		Weight              string `json:"weight,omitempty"`
		Wrap                bool   `json:"wrap,omitempty"`
		Separator           bool   `json:"separator,omitempty"`
		Spacing             string `json:"spacing,omitempty"`
		Id                  string `json:"id,omitempty"`
	}
)

func NewTextBlock(text string) *TextBlock {
	tb := &TextBlock{
		version:             1.0,
		Type:                "TextBlock",
		Text:                text,
		Color:               "",
		HorizontalAlignment: "",
		Subtle:              false,
		MaxLines:            0,
		Size:                "",
		Weight:              "",
		Wrap:                false,
		Separator:           false,
		Spacing:             "",
		Id:                  "",
	}
	return tb
}

func (tb *TextBlock) GetVersion() float64 {
	return tb.version
}

func (tb *TextBlock) GetType() string {
	return tb.Type
}

func (tb *TextBlock) SetId(id string) {
	tb.Id = id
}

func (tb *TextBlock) SetSpacing(spacing string) {
	tb.Spacing = normalize(spacing, spacingValues)
}

func (tb *TextBlock) SetSeparator(separator bool) {
	tb.Separator = separator
}

func (tb *TextBlock) SetHorizontalAlignment(horizontalAlignment string) {
	tb.HorizontalAlignment = normalize(horizontalAlignment, horizontalAlignmentValues)
}

func (tb *TextBlock) SetWrap(wrap bool) {
	tb.Wrap = wrap
}

func (tb *TextBlock) SetMaxLines(maxLines int) {
	if maxLines < 0 {
		maxLines = 0
	}
	tb.MaxLines = maxLines
}

func (tb *TextBlock) SetSize(size string) {
	tb.Size = normalize(size, textSizeValues)
}

func (tb *TextBlock) SetWeight(weight string) {
	tb.Weight = normalize(weight, textWeightValues)
}

func (tb *TextBlock) SetColor(color string) {
	tb.Color = normalize(color, textColorValues)
}

func (tb *TextBlock) SetSubtle(subtle bool) {
	tb.Subtle = subtle
}
