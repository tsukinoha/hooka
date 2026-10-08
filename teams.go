package hooka

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/tsukinoha/hooka/adaptive_card"
)

type (
	teamsMessage struct {
		Type        string                        `json:"type"`
		Attachments []*adaptive_card.AdaptiveCard `json:"attachments"`
	}
	Teams struct {
		uri     *url.URL
		message *teamsMessage
	}
)

func NewTeams(uri string) (*Teams, error) {
	// Workflows URLs are issued under logic.azure.com (legacy) or powerplatform.com.
	u, err := parseUri(uri, "logic.azure.com", "powerplatform.com")
	if err != nil {
		return nil, err
	}
	t := &Teams{
		uri: u,
		message: &teamsMessage{
			Type:        "message",
			Attachments: []*adaptive_card.AdaptiveCard{},
		},
	}
	return t, nil
}

func (t *Teams) Send(data []byte) error {
	return t.SendContext(context.Background(), data)
}

// SendContext is like Send but the request is canceled when ctx is done.
func (t *Teams) SendContext(ctx context.Context, data []byte) error {
	return send(ctx, data, t.uri)
}

func (t *Teams) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.message)
}

func (t *Teams) Attach(adaptiveCard *adaptive_card.AdaptiveCard) {
	t.message.Attachments = append(t.message.Attachments, adaptiveCard)
}
