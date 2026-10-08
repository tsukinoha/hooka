package adaptive_card

import (
	"fmt"
	"testing"
)

func TestAdaptiveCardNewContainer(t *testing.T) {
	c := NewContainer()
	if len(c.Items) != 0 {
		t.Errorf("[Case1] Expected: Container.[]Item{}, Result: %v", c.Items)
	}
}

func TestAdaptiveCardContainerGetVersion(t *testing.T) {
	cases := []struct {
		version float64
	}{
		{
			1.0,
		},
	}
	for i, c := range cases {
		cr := NewContainer()
		if cr.GetVersion() != c.version {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.version, cr.GetVersion())
		}
	}
}

func TestAdaptiveCardContainerGetType(t *testing.T) {
	cases := []struct {
		typ string
	}{
		{
			"Container",
		},
	}
	for i, c := range cases {
		cr := NewContainer()
		if cr.GetType() != c.typ {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.typ, cr.GetType())
		}
	}
}

func TestAdaptiveCardContainerSetId(t *testing.T) {
	cases := []struct {
		id string
	}{
		{
			"TestContainerId",
		},
	}
	for i, c := range cases {
		cr := NewContainer()
		cr.SetId(c.id)
		if cr.Id != c.id {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.id, cr.Id)
		}
	}
}

func TestAdaptiveCardContainerAppend(t *testing.T) {
	cases := []struct {
		items []Element
	}{
		{
			[]Element{&TextBlock{Text: "text1"}},
		},
		{
			[]Element{&TextBlock{Text: "test1"}, &TextBlock{Text: "test2"}},
		},
		{
			[]Element{&Image{Url: "https://example.com/1.png"}},
		},
		{
			[]Element{&Image{Url: "https://example.com/1.png"}, &Image{Url: "https://example.com/2.png"}},
		},
		{
			[]Element{&ImageSet{}},
		},
		{
			[]Element{&ImageSet{}, &ImageSet{}},
		},
		{
			[]Element{&FactSet{}},
		},
		{
			[]Element{&FactSet{}, &FactSet{}},
		},
		{
			[]Element{&Container{}},
		},
		{
			[]Element{&Container{}, &Container{}},
		},
		{
			[]Element{&Column{}},
		},
		{
			[]Element{&Column{}, &Column{}},
		},
		{
			[]Element{&ColumnSet{}},
		},
		{
			[]Element{&ColumnSet{}, &ColumnSet{}},
		},
	}
	for i, c := range cases {
		cr := NewContainer()
		for _, it := range c.items {
			cr.Append(it)
		}
		if len(cr.Items) != len(c.items) {
			t.Errorf("[Case%d] Item Count Expected: %v, Result: %v", i+1, len(c.items), len(cr.Items))
		}
		for j, it := range cr.Items {
			if fmt.Sprintf("%T", it) != fmt.Sprintf("%T", c.items[j]) {
				t.Errorf("[Case%d] Element Expected: %v, Result: %v", i+1, fmt.Sprintf("%T", c.items[j]), fmt.Sprintf("%T", it))
			}
		}
	}
}

func TestAdaptiveCardContainerSetStyle(t *testing.T) {
	cases := []struct {
		style    string
		expected string
	}{
		{"default", "default"},
		{"DEFAULT", "default"},
		{"emphasis", "emphasis"},
		{"EMPHASIS", "emphasis"},
		{"", ""},
		{"TestStyle", ""},
	}
	for i, c := range cases {
		cr := NewContainer()
		cr.SetStyle(c.style)
		if cr.Style != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, cr.Style)
		}
	}
}

func TestAdaptiveCardContainerSetSeparator(t *testing.T) {
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
		cr := NewContainer()
		cr.SetSeparator(c.separator)
		if cr.Separator != c.separator {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.separator, cr.Separator)
		}
	}
}

func TestAdaptiveCardContainerSetSpacing(t *testing.T) {
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
		cr := NewContainer()
		cr.SetSpacing(c.spacing)
		if cr.Spacing != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, cr.Spacing)
		}
	}
}
