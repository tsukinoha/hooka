package adaptive_card

type (
	Image struct {
		version             float64
		Type                string `json:"type"`
		Url                 string `json:"url"`
		AltText             string `json:"altText,omitempty"`
		HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
		Size                string `json:"size,omitempty"`
		Style               string `json:"style,omitempty"`
		Separator           bool   `json:"separator,omitempty"`
		Spacing             string `json:"spacing,omitempty"`
		Id                  string `json:"id,omitempty"`
	}
)

func NewImage(url string) *Image {
	i := &Image{
		version:             1.0,
		Type:                "Image",
		Url:                 url,
		AltText:             "",
		HorizontalAlignment: "",
		Size:                "",
		Style:               "",
		Separator:           false,
		Spacing:             "",
		Id:                  "",
	}
	return i
}

func (i *Image) GetVersion() float64 {
	return i.version
}

func (i *Image) GetType() string {
	return i.Type
}

func (i *Image) SetId(id string) {
	i.Id = id
}

func (i *Image) SetAltText(altText string) {
	i.AltText = altText
}

func (i *Image) SetSeparator(separator bool) {
	i.Separator = separator
}

func (i *Image) SetHorizontalAlignment(horizontalAlignment string) {
	i.HorizontalAlignment = normalize(horizontalAlignment, horizontalAlignmentValues)
}

func (i *Image) SetSize(size string) {
	i.Size = normalize(size, imageSizeValues)
}

func (i *Image) SetStyle(style string) {
	i.Style = normalize(style, imageStyleValues)
}

func (i *Image) SetSpacing(spacing string) {
	i.Spacing = normalize(spacing, spacingValues)
}
