package hooka

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type (
	service uint8
	Hooka   interface {
		Send([]byte) error
	}
)

// httpClient is shared so that TCP/TLS connections are reused between requests.
var httpClient = &http.Client{Timeout: 30 * time.Second}

// parseUri parses uri and checks that its host is one of domains or a subdomain of them.
func parseUri(uri string, domains ...string) (*url.URL, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "https" {
		return nil, fmt.Errorf(`the passed uri must use "https" scheme`)
	}
	host := strings.ToLower(u.Hostname())
	for _, domain := range domains {
		if host == domain || strings.HasSuffix(host, "."+domain) {
			return u, nil
		}
	}
	return nil, fmt.Errorf(`the passed uri's host is not in %q`, domains)
}

func send(ctx context.Context, data []byte, uri *url.URL) error {
	// Request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri.String(), bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	// Send
	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(res.Body, 1024))
		return fmt.Errorf("unexpected response: %s: %s", res.Status, strings.TrimSpace(string(body)))
	}
	// Drain the body so the connection can be reused.
	_, _ = io.Copy(io.Discard, res.Body)
	return nil
}
