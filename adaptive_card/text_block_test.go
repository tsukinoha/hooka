package adaptive_card

import (
	"testing"
)

func TestAdaptiveCardNewTextBlock(t *testing.T) {
	cases := []struct {
		text string
	}{
		{
			"test string",
		},
	}
	for i, c := range cases {
		tb := NewTextBlock(c.text)
		if tb.Text != c.text {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.text, tb.Text)
		}
	}
}

func TestAdaptiveCardTextBlockGetVersion(t *testing.T) {
	cases := []struct {
		version float64
	}{
		{
			1.0,
		},
	}
	for i, c := range cases {
		tb := NewTextBlock("test string")
		if tb.GetVersion() != c.version {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.version, tb.GetVersion())
		}
	}
}

func TestAdaptiveCardTextBlockGetType(t *testing.T) {
	cases := []struct {
		typ string
	}{
		{
			"TextBlock",
		},
	}
	for i, c := range cases {
		tb := NewTextBlock("test string")
		if tb.GetType() != c.typ {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.typ, tb.GetType())
		}
	}
}

func TestAdaptiveCardTextBlockSetId(t *testing.T) {
	cases := []struct {
		id string
	}{
		{
			"TestTextBlockId",
		},
	}
	for i, c := range cases {
		tb := NewTextBlock("test string")
		tb.SetId(c.id)
		if tb.Id != c.id {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.id, tb.Id)
		}
	}
}

func TestAdaptiveCardTextBlockSetSpacing(t *testing.T) {
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
		tb := NewTextBlock("test string")
		tb.SetSpacing(c.spacing)
		if tb.Spacing != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, tb.Spacing)
		}
	}
}

func TestAdaptiveCardTextBlockSetSeparator(t *testing.T) {
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
		tb := NewTextBlock("test string")
		tb.SetSeparator(c.separator)
		if tb.Separator != c.separator {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.separator, tb.Separator)
		}
	}
}

func TestAdaptiveCardTextBlockSetHorizontalAlignment(t *testing.T) {
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
		tb := NewTextBlock("test string")
		tb.SetHorizontalAlignment(c.ha)
		if tb.HorizontalAlignment != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, tb.HorizontalAlignment)
		}
	}
}

func TestAdaptiveCardTextBlockSetWrap(t *testing.T) {
	cases := []struct {
		wrap bool
	}{
		{
			true,
		},
		{
			false,
		},
	}
	for i, c := range cases {
		tb := NewTextBlock("test string")
		tb.SetWrap(c.wrap)
		if tb.Wrap != c.wrap {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.wrap, tb.Wrap)
		}
	}
}

func TestAdaptiveCardTextBlockSetMaxLines(t *testing.T) {
	cases := []struct {
		maxLines int
		expected int
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
	}
	for i, c := range cases {
		tb := NewTextBlock("test string")
		tb.SetMaxLines(c.maxLines)
		if tb.MaxLines != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, tb.MaxLines)
		}
	}
}

func TestAdaptiveCardTextBlockSetSize(t *testing.T) {
	cases := []struct {
		size     string
		expected string
	}{
		{"default", "default"},
		{"DEFAULT", "default"},
		{"small", "small"},
		{"SMALL", "small"},
		{"medium", "medium"},
		{"MEDIUM", "medium"},
		{"large", "large"},
		{"LARGE", "large"},
		{"extraLarge", "extraLarge"},
		{"EXTRALARGE", "extraLarge"},
		{"", ""},
		{"TestSize", ""},
	}
	for i, c := range cases {
		tb := NewTextBlock("test string")
		tb.SetSize(c.size)
		if tb.Size != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, tb.Size)
		}
	}
}

func TestAdaptiveCardTextBlockSetWeight(t *testing.T) {
	cases := []struct {
		weight   string
		expected string
	}{
		{"default", "default"},
		{"DEFAULT", "default"},
		{"lighter", "lighter"},
		{"LIGHTER", "lighter"},
		{"bolder", "bolder"},
		{"BOLDER", "bolder"},
		{"", ""},
		{"TestWeight", ""},
	}
	for i, c := range cases {
		tb := NewTextBlock("test string")
		tb.SetWeight(c.weight)
		if tb.Weight != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, tb.Weight)
		}
	}
}

func TestAdaptiveCardTextBlockSetColor(t *testing.T) {
	cases := []struct {
		color    string
		expected string
	}{
		{"default", "default"},
		{"DEFAULT", "default"},
		{"dark", "dark"},
		{"DARK", "dark"},
		{"light", "light"},
		{"LIGHT", "light"},
		{"accent", "accent"},
		{"ACCENT", "accent"},
		{"good", "good"},
		{"GOOD", "good"},
		{"warning", "warning"},
		{"WARNING", "warning"},
		{"attention", "attention"},
		{"ATTENTION", "attention"},
		{"", ""},
		{"TestWeight", ""},
	}
	for i, c := range cases {
		tb := NewTextBlock("test string")
		tb.SetColor(c.color)
		if tb.Color != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, tb.Color)
		}
	}
}

func TestAdaptiveCardTextBlockSetSubtle(t *testing.T) {
	cases := []struct {
		subtle bool
	}{
		{
			true,
		},
		{
			false,
		},
	}
	for i, c := range cases {
		tb := NewTextBlock("test string")
		tb.SetSubtle(c.subtle)
		if tb.Subtle != c.subtle {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.subtle, tb.Subtle)
		}
	}
}
