package adaptive_card

import "testing"

// Append must raise the parent's version to the highest child version and never lower it.
func TestAdaptiveCardAppendVersion(t *testing.T) {
	cases := []struct {
		name     string
		append   func(versions ...float64) float64
		versions []float64
		expected float64
	}{
		{"AdaptiveCard/higher", appendToAdaptiveCard, []float64{1.2}, 1.2},
		{"AdaptiveCard/lower", appendToAdaptiveCard, []float64{1.2, 1.0}, 1.2},
		{"AdaptiveCard/same", appendToAdaptiveCard, []float64{1.0}, 1.0},
		{"Container/higher", appendToContainer, []float64{1.3}, 1.3},
		{"Container/lower", appendToContainer, []float64{1.3, 1.1}, 1.3},
		{"Column/higher", appendToColumn, []float64{1.2}, 1.2},
		{"Column/lower", appendToColumn, []float64{1.2, 1.0}, 1.2},
		{"ColumnSet/higher", appendToColumnSet, []float64{1.2}, 1.2},
		{"ColumnSet/lower", appendToColumnSet, []float64{1.2, 1.0}, 1.2},
		{"ImageSet/higher", appendToImageSet, []float64{1.1}, 1.1},
		{"ImageSet/lower", appendToImageSet, []float64{1.1, 1.0}, 1.1},
	}
	for i, c := range cases {
		if r := c.append(c.versions...); r != c.expected {
			t.Errorf("[Case%d %s] Expected: %v, Result: %v", i+1, c.name, c.expected, r)
		}
	}
}

// Versions propagate through nested elements up to the card.
func TestAdaptiveCardAppendVersionNested(t *testing.T) {
	col := NewColumn()
	col.Append(&TextBlock{Type: "TextBlock", version: 1.2})
	cs := NewColumnSet()
	cs.Append(col)
	cr := NewContainer()
	cr.Append(cs)
	ac := New()
	ac.Append(cr)
	if ac.GetVersion() != 1.2 {
		t.Errorf("Expected: %v, Result: %v", 1.2, ac.GetVersion())
	}
}

func appendToAdaptiveCard(versions ...float64) float64 {
	ac := New()
	for _, v := range versions {
		ac.Append(&TextBlock{version: v})
	}
	return ac.GetVersion()
}

func appendToContainer(versions ...float64) float64 {
	c := NewContainer()
	for _, v := range versions {
		c.Append(&TextBlock{version: v})
	}
	return c.GetVersion()
}

func appendToColumn(versions ...float64) float64 {
	c := NewColumn()
	for _, v := range versions {
		c.Append(&TextBlock{version: v})
	}
	return c.GetVersion()
}

func appendToColumnSet(versions ...float64) float64 {
	cs := NewColumnSet()
	for _, v := range versions {
		cs.Append(&Column{version: v})
	}
	return cs.GetVersion()
}

func appendToImageSet(versions ...float64) float64 {
	is := NewImageSet()
	for _, v := range versions {
		is.Append(&Image{version: v})
	}
	return is.GetVersion()
}
