package adaptive_card

type (
	FactSet struct {
		version   float64
		Type      string  `json:"type"`
		Id        string  `json:"id,omitempty"`
		Facts     []*Fact `json:"facts"`
		Separator bool    `json:"separator,omitempty"`
		Spacing   string  `json:"spacing,omitempty"`
	}
)

func NewFactSet() *FactSet {
	fs := &FactSet{
		version:   1.0,
		Type:      "FactSet",
		Id:        "",
		Facts:     []*Fact{},
		Separator: false,
		Spacing:   "",
	}
	return fs
}

func (fs *FactSet) GetVersion() float64 {
	return fs.version
}

func (fs *FactSet) GetType() string {
	return fs.Type
}

func (fs *FactSet) SetId(id string) {
	fs.Id = id
}

func (fs *FactSet) Append(fact *Fact) {
	fs.Facts = append(fs.Facts, fact)
}

func (fs *FactSet) SetSeparator(separator bool) {
	fs.Separator = separator
}

func (fs *FactSet) SetSpacing(spacing string) {
	fs.Spacing = normalize(spacing, spacingValues)
}
