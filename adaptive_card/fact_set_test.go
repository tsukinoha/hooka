package adaptive_card

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestAdaptiveCardFactSetNewFactSet(t *testing.T) {
	fs := NewFactSet()
	if len(fs.Facts) != 0 {
		t.Errorf("[Case1] Expected: FactSet.[]Fact{}, Result: %v", fs.Facts)
	}
}

func TestAdaptiveCardFactSetGetVersion(t *testing.T) {
	cases := []struct {
		version float64
	}{
		{
			1.0,
		},
	}
	for i, c := range cases {
		fs := NewFactSet()
		if fs.GetVersion() != c.version {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.version, fs.GetVersion())
		}
	}
}

func TestAdaptiveCardFactSetGetType(t *testing.T) {
	cases := []struct {
		typ string
	}{
		{
			"FactSet",
		},
	}
	for i, c := range cases {
		fs := NewFactSet()
		if fs.GetType() != c.typ {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.typ, fs.GetType())
		}
	}
}

func TestAdaptiveCardFactSetSetId(t *testing.T) {
	cases := []struct {
		id string
	}{
		{
			"TestFactSetId",
		},
	}
	for i, c := range cases {
		fs := NewFactSet()
		fs.SetId(c.id)
		if fs.Id != c.id {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.id, fs.Id)
		}
	}
}

func TestAdaptiveCardFactSetAppend(t *testing.T) {
	cases := []struct {
		facts []*Fact
	}{
		{
			[]*Fact{&Fact{Title: "title1", Value: "value1"}},
		},
		{
			[]*Fact{&Fact{Title: "title1", Value: "value1"}, &Fact{Title: "title2", Value: "value2"}},
		},
	}
	for i, c := range cases {
		fs := NewFactSet()
		for _, fact := range c.facts {
			fs.Append(fact)
		}
		if len(fs.Facts) != len(c.facts) {
			t.Errorf("[Case%d] Fact Count Expected: %v, Result: %v", i+1, len(c.facts), len(fs.Facts))
		}
		for j, fact := range fs.Facts {
			if fact.Title != c.facts[j].Title {
				t.Errorf("[Case%d] Fact Title Expected: %v, Result: %v", i+1, c.facts[j].Title, fact.Title)
			}
			if fact.Value != c.facts[j].Value {
				t.Errorf("[Case%d] Fact Value Expected: %v, Result: %v", i+1, c.facts[j].Value, fact.Value)
			}
		}
	}
}

func TestAdaptiveCardFactSetSetSeparator(t *testing.T) {
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
		fs := NewFactSet()
		fs.SetSeparator(c.separator)
		if fs.Separator != c.separator {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.separator, fs.Separator)
		}
	}
}

func TestAdaptiveCardFactSetSetSpacing(t *testing.T) {
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
		fs := NewFactSet()
		fs.SetSpacing(c.spacing)
		if fs.Spacing != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, fs.Spacing)
		}
	}
}

func TestAdaptiveCardFactSetMarshalEmptyFacts(t *testing.T) {
	data, err := json.Marshal(NewFactSet())
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"facts":[]`) {
		t.Errorf("Expected: \"facts\":[], Result: %s", data)
	}
}
