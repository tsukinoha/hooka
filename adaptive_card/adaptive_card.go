package adaptive_card

import (
	"encoding/json"
	"strconv"
)

type (
	AdaptiveCard struct {
		ContentType string              `json:"contentType"`
		Content     AdaptiveCardContent `json:"content"`
	}
	AdaptiveCardContent struct {
		Ver             string    `json:"version"`
		Schema          string    `json:"$schema"`
		Type            string    `json:"type"`
		Version         float64   `json:"-"`
		Bodies          []Element `json:"body"`
		FallbackText    string    `json:"fallbackText,omitempty"`
		BackgroundImage string    `json:"backgroundImage,omitempty"`
		Speak           string    `json:"speak,omitempty"`
		Lang            string    `json:"lang,omitempty"`
	}
	Element interface {
		GetVersion() float64
		GetType() string
	}
)

func New() *AdaptiveCard {
	ac := &AdaptiveCard{
		ContentType: "application/vnd.microsoft.card.adaptive",
		Content: AdaptiveCardContent{
			Ver:             "1.0",
			Schema:          "http://adaptivecards.io/schemas/adaptive-card.json",
			Type:            "AdaptiveCard",
			Version:         1.0,
			Bodies:          []Element{},
			FallbackText:    "",
			BackgroundImage: "",
			Speak:           "",
			Lang:            "",
		},
	}
	return ac
}

func (ac *AdaptiveCard) GetVersion() float64 {
	return ac.Content.Version
}

func (ac *AdaptiveCard) GetType() string {
	return ac.Content.Type
}

func (ac *AdaptiveCard) Append(elem Element) {
	ac.Content.Bodies = append(ac.Content.Bodies, elem)
	if elem.GetVersion() > ac.GetVersion() {
		ac.Content.Version = elem.GetVersion()
	}
}

func (ac *AdaptiveCard) SetLang(lang string) {
	ac.Content.Lang = lang
}

func (ac *AdaptiveCard) Marshal() ([]byte, error) {
	return json.Marshal(ac)
}

// MarshalJSON writes "version" from Version so that it is correct however the card is marshaled.
func (acc AdaptiveCardContent) MarshalJSON() ([]byte, error) {
	type content AdaptiveCardContent
	acc.Ver = strconv.FormatFloat(acc.Version, 'f', 1, 64)
	return json.Marshal(content(acc))
}
