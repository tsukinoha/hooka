package hooka

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/tsukinoha/hooka/adaptive_card"
)

func TestParseUri(t *testing.T) {
	domains := []string{"logic.azure.com", "powerplatform.com"}
	cases := []struct {
		uri string
		ok  bool
	}{
		{"https://prod-00.japaneast.logic.azure.com/workflows/xxx", true},
		{"https://default.environment.api.powerplatform.com/powerautomate/xxx", true},
		{"http://prod-00.japaneast.logic.azure.com/workflows/xxx", false},
		{"https://logic.azure.com.example.com/workflows/xxx", false},
		{"https://evillogic.azure.com/workflows/xxx", false},
		{"https://example.com/?logic.azure.com", false},
		{"https://PROD-00.JAPANEAST.LOGIC.AZURE.COM/workflows/xxx", true},
		{"https://prod-00.japaneast.logic.azure.com:443/workflows/xxx", true},
		{"https://logic.azure.com@example.com/workflows/xxx", false},
		{"https://logic.azure.com/workflows/xxx", true},
		{"https://azure.com/workflows/xxx", false},
		{"logic.azure.com/workflows/xxx", false},
		{"", false},
		{"https://%zz", false},
	}
	for i, c := range cases {
		_, err := parseUri(c.uri, domains...)
		if (err == nil) != c.ok {
			t.Errorf("[Case%d] uri: %s, Expected ok: %v, Result: %v", i+1, c.uri, c.ok, err)
		}
	}
}

func TestSend(t *testing.T) {
	cases := []struct {
		status int
		ok     bool
	}{
		{http.StatusOK, true},
		{http.StatusAccepted, true},
		{http.StatusBadRequest, false},
		{http.StatusInternalServerError, false},
	}
	for i, c := range cases {
		var gotType, gotBody string
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gotType = r.Header.Get("Content-Type")
			b, _ := io.ReadAll(r.Body)
			gotBody = string(b)
			w.WriteHeader(c.status)
		}))
		u, _ := url.Parse(srv.URL)
		err := send(context.Background(), []byte(`{"a":1}`), u)
		srv.Close()
		if (err == nil) != c.ok {
			t.Errorf("[Case%d] Expected ok: %v, Result: %v", i+1, c.ok, err)
		}
		if gotType != "application/json" || gotBody != `{"a":1}` {
			t.Errorf("[Case%d] Unexpected request: %s %s", i+1, gotType, gotBody)
		}
	}
}

func TestSendErrorMessage(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid card\n"))
	}))
	defer srv.Close()
	u, _ := url.Parse(srv.URL)
	err := send(context.Background(), []byte(`{}`), u)
	if err == nil || !strings.Contains(err.Error(), "400 Bad Request") || !strings.Contains(err.Error(), "invalid card") {
		t.Errorf("Error should contain status and body, Result: %v", err)
	}
}

func TestSendConnectionError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	u, _ := url.Parse(srv.URL)
	srv.Close()
	if err := send(context.Background(), []byte(`{}`), u); err == nil {
		t.Errorf("Expected error for closed server")
	}
}

func TestSendContextCanceled(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer srv.Close()
	u, _ := url.Parse(srv.URL)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := send(ctx, []byte(`{}`), u); !errors.Is(err, context.Canceled) {
		t.Errorf("Expected: %v, Result: %v", context.Canceled, err)
	}
}

func TestTeamsMarshalJSON(t *testing.T) {
	tm, err := NewTeams("https://prod-00.japaneast.logic.azure.com/workflows/xxx")
	if err != nil {
		t.Fatal(err)
	}
	ac := adaptive_card.New()
	ac.Append(adaptive_card.NewTextBlock("hello"))
	tm.Attach(ac)
	data, err := json.Marshal(tm)
	if err != nil {
		t.Fatal(err)
	}
	var m struct {
		Type        string `json:"type"`
		Attachments []struct {
			ContentType string `json:"contentType"`
			Content     struct {
				Version string `json:"version"`
				Body    []any  `json:"body"`
			} `json:"content"`
		} `json:"attachments"`
	}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if m.Type != "message" || len(m.Attachments) != 1 || m.Attachments[0].Content.Version != "1.0" || len(m.Attachments[0].Content.Body) != 1 {
		t.Errorf("Unexpected JSON: %s", data)
	}
	if strings.Contains(string(data), "1.00") {
		t.Errorf("Version should be formatted as \"1.0\": %s", data)
	}
}

func TestNewTeams(t *testing.T) {
	cases := []struct {
		uri string
		ok  bool
	}{
		{"https://prod-00.japaneast.logic.azure.com/workflows/xxx", true},
		{"https://default.environment.api.powerplatform.com/powerautomate/xxx", true},
		{"https://example.com/workflows/xxx", false},
		{"https://management.azure.com/workflows/xxx", false},
		{"", false},
	}
	for i, c := range cases {
		tm, err := NewTeams(c.uri)
		if (err == nil) != c.ok {
			t.Errorf("[Case%d] uri: %s, Expected ok: %v, Result: %v", i+1, c.uri, c.ok, err)
		}
		if c.ok && (tm == nil || tm.uri.String() != c.uri || tm.message.Type != "message" || len(tm.message.Attachments) != 0) {
			t.Errorf("[Case%d] Unexpected Teams: %+v", i+1, tm)
		}
		if !c.ok && tm != nil {
			t.Errorf("[Case%d] Expected nil Teams, Result: %+v", i+1, tm)
		}
	}
}

func TestTeamsAttach(t *testing.T) {
	tm, _ := NewTeams("https://prod-00.japaneast.logic.azure.com/workflows/xxx")
	ac1, ac2 := adaptive_card.New(), adaptive_card.New()
	tm.Attach(ac1)
	tm.Attach(ac2)
	if len(tm.message.Attachments) != 2 || tm.message.Attachments[0] != ac1 || tm.message.Attachments[1] != ac2 {
		t.Errorf("Unexpected attachments: %v", tm.message.Attachments)
	}
}

func TestTeamsMarshalJSONEmpty(t *testing.T) {
	tm, _ := NewTeams("https://prod-00.japaneast.logic.azure.com/workflows/xxx")
	data, err := json.Marshal(tm)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `{"type":"message","attachments":[]}` {
		t.Errorf("Unexpected JSON: %s", data)
	}
}

func TestTeamsSend(t *testing.T) {
	var got []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/workflows/xxx" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		got, _ = io.ReadAll(r.Body)
		w.WriteHeader(http.StatusAccepted)
	}))
	defer srv.Close()
	// NewTeams rejects non-Teams hosts, so build Teams directly for the test server.
	u, _ := url.Parse(srv.URL + "/workflows/xxx")
	tm := &Teams{uri: u, message: &teamsMessage{Type: "message", Attachments: []*adaptive_card.AdaptiveCard{}}}
	tm.Attach(adaptive_card.New())
	data, _ := json.Marshal(tm)
	if err := tm.Send(data); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if string(got) != string(data) {
		t.Errorf("Expected: %s, Result: %s", data, got)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := tm.SendContext(ctx, data); !errors.Is(err, context.Canceled) {
		t.Errorf("Expected: %v, Result: %v", context.Canceled, err)
	}
}
