package adaptive_card

import "testing"

func TestAdaptiveCardNewColumnSet(t *testing.T) {
	cs := NewColumnSet()
	if len(cs.Columns) != 0 {
		t.Errorf("[Case1] Expected: ColumnSet.[]Column{}, Result: %v", cs.Columns)
	}
}

func TestAdaptiveCardColumnSetGetVersion(t *testing.T) {
	cases := []struct {
		version float64
	}{
		{
			1.0,
		},
	}
	for i, c := range cases {
		cs := NewColumnSet()
		if cs.GetVersion() != c.version {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.version, cs.GetVersion())
		}
	}
}

func TestAdaptiveCardColumnSetGetType(t *testing.T) {
	cases := []struct {
		typ string
	}{
		{
			"ColumnSet",
		},
	}
	for i, c := range cases {
		cs := NewColumnSet()
		if cs.GetType() != c.typ {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.typ, cs.GetType())
		}
	}
}

func TestAdaptiveCardColumnSetSetId(t *testing.T) {
	cases := []struct {
		id string
	}{
		{
			"TestColumnSetId",
		},
	}
	for i, c := range cases {
		cs := NewColumnSet()
		cs.SetId(c.id)
		if cs.Id != c.id {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.id, cs.Id)
		}
	}
}

func TestAdaptiveCardColumnSetAppend(t *testing.T) {
	cases := []struct {
		columns []*Column
	}{
		{
			[]*Column{&Column{}},
		},
		{
			[]*Column{&Column{}, &Column{}},
		},
	}
	for i, c := range cases {
		cs := NewColumnSet()
		for _, column := range c.columns {
			cs.Append(column)
		}
		if len(cs.Columns) != len(c.columns) {
			t.Errorf("[Case%d] Column Count Expected: %v, Result: %v", i+1, len(c.columns), len(cs.Columns))
		}
	}
}

func TestAdaptiveCardColumnSetSetSpacing(t *testing.T) {
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
		cs := NewColumnSet()
		cs.SetSpacing(c.spacing)
		if cs.Spacing != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, cs.Spacing)
		}
	}
}

func TestAdaptiveCardColumnSetSetSeparator(t *testing.T) {
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
		cs := NewColumnSet()
		cs.SetSeparator(c.separator)
		if cs.Separator != c.separator {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.separator, cs.Separator)
		}
	}
}

func TestAdaptiveCardColumnSetSetHorizontalAlignment(t *testing.T) {
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
		cs := NewColumnSet()
		cs.SetHorizontalAlignment(c.ha)
		if cs.HorizontalAlignment != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, cs.HorizontalAlignment)
		}
	}
}
