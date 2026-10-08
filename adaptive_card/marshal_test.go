package adaptive_card

import (
	"encoding/json"
	"testing"
)

func TestAdaptiveCardElementMarshalJSON(t *testing.T) {
	cases := []struct {
		name     string
		elem     func() any
		expected string
	}{
		{
			"TextBlock/minimum",
			func() any { return NewTextBlock("hello") },
			`{"type":"TextBlock","text":"hello"}`,
		},
		{
			"TextBlock/full",
			func() any {
				tb := NewTextBlock("hello")
				tb.SetId("tb1")
				tb.SetColor("accent")
				tb.SetHorizontalAlignment("center")
				tb.SetSubtle(true)
				tb.SetMaxLines(2)
				tb.SetSize("extraLarge")
				tb.SetWeight("bolder")
				tb.SetWrap(true)
				tb.SetSeparator(true)
				tb.SetSpacing("small")
				return tb
			},
			`{"type":"TextBlock","text":"hello","color":"accent","horizontalAlignment":"center","isSubtle":true,"maxLines":2,"size":"extraLarge","weight":"bolder","wrap":true,"separator":true,"spacing":"small","id":"tb1"}`,
		},
		{
			"Image/minimum",
			func() any { return NewImage("https://example.com/1.png") },
			`{"type":"Image","url":"https://example.com/1.png"}`,
		},
		{
			"Image/full",
			func() any {
				im := NewImage("https://example.com/1.png")
				im.SetId("im1")
				im.SetAltText("alt")
				im.SetHorizontalAlignment("right")
				im.SetSize("small")
				im.SetStyle("person")
				im.SetSeparator(true)
				im.SetSpacing("none")
				return im
			},
			`{"type":"Image","url":"https://example.com/1.png","altText":"alt","horizontalAlignment":"right","size":"small","style":"person","separator":true,"spacing":"none","id":"im1"}`,
		},
		{
			"ImageSet/empty",
			func() any { return NewImageSet() },
			`{"type":"ImageSet","images":[]}`,
		},
		{
			"ImageSet/full",
			func() any {
				is := NewImageSet()
				is.SetId("is1")
				is.Append(NewImage("https://example.com/1.png"))
				is.SetImageSize("medium")
				is.SetSpacing("large")
				is.SetSeparator(true)
				return is
			},
			`{"type":"ImageSet","id":"is1","images":[{"type":"Image","url":"https://example.com/1.png"}],"imageSize":"medium","spacing":"large","separator":true}`,
		},
		{
			"FactSet/full",
			func() any {
				fs := NewFactSet()
				fs.SetId("fs1")
				fs.Append(NewFact("t", "v"))
				fs.SetSeparator(true)
				fs.SetSpacing("medium")
				return fs
			},
			`{"type":"FactSet","id":"fs1","facts":[{"title":"t","value":"v"}],"separator":true,"spacing":"medium"}`,
		},
		{
			"Container/empty",
			func() any { return NewContainer() },
			`{"type":"Container","items":[]}`,
		},
		{
			"Container/full",
			func() any {
				c := NewContainer()
				c.SetId("c1")
				c.Append(NewTextBlock("a"))
				c.SetStyle("emphasis")
				c.SetSeparator(true)
				c.SetSpacing("padding")
				return c
			},
			`{"type":"Container","id":"c1","items":[{"type":"TextBlock","text":"a"}],"style":"emphasis","separator":true,"spacing":"padding"}`,
		},
		{
			"Column/no width",
			func() any { return NewColumn() },
			`{"type":"Column","items":[]}`,
		},
		{
			"Column/string width",
			func() any {
				c := NewColumn()
				c.SetWidth("stretch")
				return c
			},
			`{"type":"Column","items":[],"width":"stretch"}`,
		},
		{
			"Column/int width",
			func() any {
				c := NewColumn()
				c.SetWidth("2")
				return c
			},
			`{"type":"Column","items":[],"width":2}`,
		},
		{
			"Column/full",
			func() any {
				c := NewColumn()
				c.SetId("col1")
				c.Append(NewTextBlock("a"))
				c.SetSeparator(true)
				c.SetSpacing("default")
				c.SetStyle("default")
				c.SetWidth("auto")
				return c
			},
			`{"type":"Column","id":"col1","items":[{"type":"TextBlock","text":"a"}],"separator":true,"spacing":"default","style":"default","width":"auto"}`,
		},
		{
			"ColumnSet/empty",
			func() any { return NewColumnSet() },
			`{"type":"ColumnSet"}`,
		},
		{
			"ColumnSet/full",
			func() any {
				cs := NewColumnSet()
				cs.SetId("cs1")
				col := NewColumn()
				col.SetWidth("1")
				cs.Append(col)
				cs.SetHorizontalAlignment("left")
				cs.SetSpacing("small")
				cs.SetSeparator(true)
				return cs
			},
			`{"type":"ColumnSet","id":"cs1","columns":[{"type":"Column","items":[],"width":1}],"horizontalAlignment":"left","spacing":"small","separator":true}`,
		},
	}
	for i, c := range cases {
		data, err := json.Marshal(c.elem())
		if err != nil {
			t.Errorf("[Case%d %s] Unexpected error: %v", i+1, c.name, err)
			continue
		}
		if string(data) != c.expected {
			t.Errorf("[Case%d %s]\nExpected: %s\nResult:   %s", i+1, c.name, c.expected, data)
		}
	}
}

func TestAdaptiveCardMarshal(t *testing.T) {
	cases := []struct {
		name     string
		card     func() *AdaptiveCard
		expected string
	}{
		{
			"empty",
			func() *AdaptiveCard { return New() },
			`{"contentType":"application/vnd.microsoft.card.adaptive","content":{"version":"1.0","$schema":"http://adaptivecards.io/schemas/adaptive-card.json","type":"AdaptiveCard","body":[]}}`,
		},
		{
			"body and lang",
			func() *AdaptiveCard {
				ac := New()
				ac.SetLang("ja")
				ac.Append(NewTextBlock("hello"))
				return ac
			},
			`{"contentType":"application/vnd.microsoft.card.adaptive","content":{"version":"1.0","$schema":"http://adaptivecards.io/schemas/adaptive-card.json","type":"AdaptiveCard","body":[{"type":"TextBlock","text":"hello"}],"lang":"ja"}}`,
		},
		{
			"raised version",
			func() *AdaptiveCard {
				ac := New()
				ac.Append(&TextBlock{Type: "TextBlock", Text: "a", version: 1.5})
				return ac
			},
			`{"contentType":"application/vnd.microsoft.card.adaptive","content":{"version":"1.5","$schema":"http://adaptivecards.io/schemas/adaptive-card.json","type":"AdaptiveCard","body":[{"type":"TextBlock","text":"a"}]}}`,
		},
	}
	for i, c := range cases {
		ac := c.card()
		data, err := ac.Marshal()
		if err != nil {
			t.Errorf("[Case%d %s] Unexpected error: %v", i+1, c.name, err)
			continue
		}
		if string(data) != c.expected {
			t.Errorf("[Case%d %s]\nExpected: %s\nResult:   %s", i+1, c.name, c.expected, data)
		}
		// json.Marshal must give the same result as Marshal.
		data2, _ := json.Marshal(ac)
		if string(data2) != string(data) {
			t.Errorf("[Case%d %s] json.Marshal differs from Marshal\nMarshal:      %s\njson.Marshal: %s", i+1, c.name, data, data2)
		}
	}
}

// Marshaling must not modify the card.
func TestAdaptiveCardMarshalIdempotent(t *testing.T) {
	ac := New()
	ac.Append(&TextBlock{Type: "TextBlock", Text: "a", version: 1.2})
	first, _ := ac.Marshal()
	second, _ := ac.Marshal()
	if string(first) != string(second) {
		t.Errorf("Expected same output\nFirst:  %s\nSecond: %s", first, second)
	}
	if ac.Content.Ver != "1.0" {
		t.Errorf("Content.Ver should not be modified by Marshal, Result: %v", ac.Content.Ver)
	}
}
