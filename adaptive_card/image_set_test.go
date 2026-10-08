package adaptive_card

import (
	"testing"
)

func TestAdaptiveCardNewImageSet(t *testing.T) {
	is := NewImageSet()
	if len(is.Images) != 0 {
		t.Errorf("[Case1] Expected: ImageSet.[]Image{}, Result: %v", is.Images)
	}
}

func TestAdaptiveCardImageSetGetVersion(t *testing.T) {
	cases := []struct {
		version float64
	}{
		{
			1.0,
		},
	}
	for i, c := range cases {
		is := NewImageSet()
		if is.GetVersion() != c.version {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.version, is.GetVersion())
		}
	}
}

func TestAdaptiveCardImageSetGetType(t *testing.T) {
	cases := []struct {
		typ string
	}{
		{
			"ImageSet",
		},
	}
	for i, c := range cases {
		is := NewImageSet()
		if is.GetType() != c.typ {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.typ, is.GetType())
		}
	}
}

func TestAdaptiveCardImageSetSetId(t *testing.T) {
	cases := []struct {
		id string
	}{
		{
			"TestImageSetId",
		},
	}
	for i, c := range cases {
		is := NewImageSet()
		is.SetId(c.id)
		if is.Id != c.id {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.id, is.Id)
		}
	}
}

func TestAdaptiveCardImageSetAppend(t *testing.T) {
	cases := []struct {
		images []*Image
	}{
		{
			[]*Image{&Image{Url: "https://example.com/1.png"}},
		},
		{
			[]*Image{&Image{Url: "https://example.com/1.png"}, &Image{Url: "https://example.com/2.png"}},
		},
	}
	for i, c := range cases {
		is := NewImageSet()
		for _, im := range c.images {
			is.Append(im)
		}
		if len(is.Images) != len(c.images) {
			t.Errorf("[Case%d] Image Count Expected: %v, Result: %v", i+1, len(c.images), len(is.Images))
		}
		for j, im := range is.Images {
			if im.Url != c.images[j].Url {
				t.Errorf("[Case%d] Image URL Expected: %v, Result: %v", i+1, c.images[j].Url, im.Url)
			}
		}
	}
}

func TestAdaptiveCardImageSetSetImageSize(t *testing.T) {
	cases := []struct {
		imageSize string
		expected  string
	}{
		{"auto", "auto"},
		{"AUTO", "auto"},
		{"stretch", "stretch"},
		{"STRETCH", "stretch"},
		{"small", "small"},
		{"SMALL", "small"},
		{"medium", "medium"},
		{"MEDIUM", "medium"},
		{"large", "large"},
		{"LARGE", "large"},
		{"", ""},
		{"TestSize", ""},
	}
	for i, c := range cases {
		is := NewImageSet()
		is.SetImageSize(c.imageSize)
		if is.ImageSize != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, is.ImageSize)
		}
	}
}

func TestAdaptiveCardImageSetSetSpacing(t *testing.T) {
	cases := []struct {
		spacing  string
		expected string
	}{
		{"default", "default"},
		{"DEFAULT", "default"},
		{"none", "none"},
		{"NONE", "none"},
		{"small", "small"},
		{"SMALL", "small"},
		{"medium", "medium"},
		{"MEDIUM", "medium"},
		{"large", "large"},
		{"LARGE", "large"},
		{"extraLarge", "extraLarge"},
		{"EXTRALARGE", "extraLarge"},
		{"padding", "padding"},
		{"PADDING", "padding"},
		{"test spacing", ""},
		{"", ""},
	}
	for i, c := range cases {
		is := NewImageSet()
		is.SetSpacing(c.spacing)
		if is.Spacing != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, is.Spacing)
		}
	}
}

func TestAdaptiveCardImageSetSetSeparator(t *testing.T) {
	cases := []struct {
		separator bool
	}{
		{
			true,
		},
		{
			false,
		},
	}
	for i, c := range cases {
		is := NewImageSet()
		is.SetSeparator(c.separator)
		if is.Separator != c.separator {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.separator, is.Separator)
		}
	}
}
