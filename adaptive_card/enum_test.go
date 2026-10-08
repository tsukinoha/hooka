package adaptive_card

import "testing"

func TestAdaptiveCardNormalize(t *testing.T) {
	allowed := []string{"default", "extraLarge"}
	cases := []struct {
		value    string
		expected string
	}{
		{"default", "default"},
		{"DEFAULT", "default"},
		{"extraLarge", "extraLarge"},
		{"extralarge", "extraLarge"},
		{"EXTRALARGE", "extraLarge"},
		{"extra large", ""},
		{" default", ""},
		{"", ""},
	}
	for i, c := range cases {
		if r := normalize(c.value, allowed); r != c.expected {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.expected, r)
		}
	}
	if r := normalize("default", nil); r != "" {
		t.Errorf("[Case%d] Expected: %v, Result: %v", len(cases)+1, "", r)
	}
}
