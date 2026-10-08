package adaptive_card

import (
	"fmt"
	"testing"
)

func TestAdaptiveCardNewColumn(t *testing.T) {
	cl := NewColumn()
	if len(cl.Items) != 0 {
		t.Errorf("[Case1] Expected: Column.[]Item{}, Result: %v", cl.Items)
	}
}

func TestAdaptiveCardColumnGetVersion(t *testing.T) {
	cases := []struct {
		version float64
	}{
		{
			1.0,
		},
	}
	for i, c := range cases {
		cl := NewColumn()
		if cl.GetVersion() != c.version {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.version, cl.GetVersion())
		}
	}
}

func TestAdaptiveCardColumnGetType(t *testing.T) {
	cases := []struct {
		typ string
	}{
		{
			"Column",
		},
	}
	for i, c := range cases {
		cl := NewColumn()
		if cl.GetType() != c.typ {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.typ, cl.GetType())
		}
	}
}

func TestAdaptiveCardColumnSetId(t *testing.T) {
	cases := []struct {
		id string
	}{
		{
			"TestColumnId",
		},
	}
	for i, c := range cases {
		cl := NewColumn()
		cl.SetId(c.id)
		if cl.Id != c.id {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.id, cl.Id)
		}
	}
}

func TestAdaptiveCardColumnAppend(t *testing.T) {
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
		cl := NewColumn()
		for _, it := range c.items {
			cl.Append(it)
		}
		if len(cl.Items) != len(c.items) {
			t.Errorf("[Case%d] Item Count Expected: %v, Result: %v", i+1, len(c.items), len(cl.Items))
		}
		for j, it := range cl.Items {
			if fmt.Sprintf("%T", it) != fmt.Sprintf("%T", c.items[j]) {
				t.Errorf("[Case%d] Element Expected: %v, Result: %v", i+1, fmt.Sprintf("%T", c.items[j]), fmt.Sprintf("%T", it))
			}
		}
	}
}

func TestAdaptiveCardColumnSetSpacing(t *testing.T) {
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
		cl := NewColumn()
		cl.SetSpacing(c.spacing)
		if cl.Spacing != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, cl.Spacing)
		}
	}
}

func TestAdaptiveCardColumnSetSeparator(t *testing.T) {
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
		cl := NewColumn()
		cl.SetSeparator(c.separator)
		if cl.Separator != c.separator {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.separator, cl.Separator)
		}
	}
}

func TestAdaptiveCardColumnSetStyle(t *testing.T) {
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
		cl := NewColumn()
		cl.SetStyle(c.style)
		if cl.Style != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, cl.Style)
		}
	}
}

func TestAdaptiveCardColumnSetWidth(t *testing.T) {
	cases := []struct {
		width    string
		expected string
	}{
		{"auto", "auto"},
		{"AUTO", "auto"},
		{"stretch", "stretch"},
		{"STRETCH", "stretch"},
		{"", ""},
		{"TestWidth", ""},
		{"0", "0"},
		{"-1", "0"},
		{"1", "1"},
		{"100", "100"},
		{"Auto", "auto"},
	}
	for i, c := range cases {
		cl := NewColumn()
		cl.SetWidth(c.width)
		if cl.width != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, cl.width)
		}
	}
}
