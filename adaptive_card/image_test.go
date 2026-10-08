package adaptive_card

import (
	"testing"
)

func TestAdaptiveCardNewImage(t *testing.T) {
	cases := []struct {
		url string
	}{
		{
			"https://example.com/example.png",
		},
	}
	for i, c := range cases {
		im := NewImage(c.url)
		if im.Url != c.url {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.url, im.Url)
		}
	}
}

func TestAdaptiveCardImageGetVersion(t *testing.T) {
	cases := []struct {
		version float64
	}{
		{
			1.0,
		},
	}
	for i, c := range cases {
		im := NewImage("https://example.com/example.png")
		if im.GetVersion() != c.version {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.version, im.GetVersion())
		}
	}
}

func TestAdaptiveCardImageGetType(t *testing.T) {
	cases := []struct {
		typ string
	}{
		{
			"Image",
		},
	}
	for i, c := range cases {
		im := NewImage("https://example.com/example.png")
		if im.GetType() != c.typ {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.typ, im.GetType())
		}
	}
}

func TestAdaptiveCardImageSetId(t *testing.T) {
	cases := []struct {
		id string
	}{
		{
			"TestImageId",
		},
	}
	for i, c := range cases {
		im := NewImage("https://example.com/example.png")
		im.SetId(c.id)
		if im.Id != c.id {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.id, im.Id)
		}
	}
}

func TestAdaptiveCardImageSetAltText(t *testing.T) {
	cases := []struct {
		altText string
	}{
		{
			"TestAltText",
		},
	}
	for i, c := range cases {
		im := NewImage("https://example.com/example.png")
		im.SetAltText(c.altText)
		if im.AltText != c.altText {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.altText, im.AltText)
		}
	}
}

func TestAdaptiveCardImageSetSeparator(t *testing.T) {
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
		im := NewImage("test string")
		im.SetSeparator(c.separator)
		if im.Separator != c.separator {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.separator, im.Separator)
		}
	}
}

func TestAdaptiveCardImageSetHorizontalAlignment(t *testing.T) {
	cases := []struct {
		ha       string
		expected string
	}{
		{"left", "left"},
		{"LEFT", "left"},
		{"center", "center"},
		{"CENTER", "center"},
		{"right", "right"},
		{"RIGHT", "right"},
		{"", ""},
		{"TestAlignment", ""},
	}
	for i, c := range cases {
		im := NewImage("test string")
		im.SetHorizontalAlignment(c.ha)
		if im.HorizontalAlignment != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, im.HorizontalAlignment)
		}
	}
}

func TestAdaptiveCardImageSetSize(t *testing.T) {
	cases := []struct {
		size     string
		expected string
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
		im := NewImage("test string")
		im.SetSize(c.size)
		if im.Size != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, im.Size)
		}
	}
}

func TestAdaptiveCardImageSetStyle(t *testing.T) {
	cases := []struct {
		style    string
		expected string
	}{
		{"default", "default"},
		{"DEFAULT", "default"},
		{"person", "person"},
		{"PERSON", "person"},
		{"", ""},
		{"TestStyle", ""},
	}
	for i, c := range cases {
		im := NewImage("test string")
		im.SetStyle(c.style)
		if im.Style != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, im.Style)
		}
	}
}

func TestAdaptiveCardImageSetSpacing(t *testing.T) {
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
		im := NewImage("test string")
		im.SetSpacing(c.spacing)
		if im.Spacing != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, im.Spacing)
		}
	}
}
