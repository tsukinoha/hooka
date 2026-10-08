package adaptive_card

type (
	ImageSet struct {
		version             float64
		Type                string   `json:"type"`
		Id                  string   `json:"id,omitempty"`
		Images              []*Image `json:"images"`
		ImageSize           string   `json:"imageSize,omitempty"`
		Spacing             string   `json:"spacing,omitempty"`
		Separator           bool     `json:"separator,omitempty"`
		HorizontalAlignment string   `json:"horizontalAlignment,omitempty"`
	}
)

func NewImageSet() *ImageSet {
	is := &ImageSet{
		version:   1.0,
		Type:      "ImageSet",
		Id:        "",
		Images:    []*Image{},
		ImageSize: "",
		Separator: false,
		Spacing:   "",
	}
	return is
}

func (is *ImageSet) GetVersion() float64 {
	return is.version
}

func (is *ImageSet) GetType() string {
	return is.Type
}

func (is *ImageSet) SetId(id string) {
	is.Id = id
}

func (is *ImageSet) Append(image *Image) {
	is.Images = append(is.Images, image)
	if image.GetVersion() > is.GetVersion() {
		is.version = image.GetVersion()
	}
}

func (is *ImageSet) SetImageSize(imageSize string) {
	is.ImageSize = normalize(imageSize, imageSizeValues)
}

func (is *ImageSet) SetSpacing(spacing string) {
	is.Spacing = normalize(spacing, spacingValues)
}

func (is *ImageSet) SetSeparator(separator bool) {
	is.Separator = separator
}
