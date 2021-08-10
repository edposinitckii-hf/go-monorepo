package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/edposinitckii-hf/monorepo/svc/echo/internal/transport"
)

type MessageEcho struct {
	Message   string    `json:"message,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Error     string    `json:"error,omitempty"`
}

// EchoClient represents client for echo Server
type EchoClient struct {
	host string
}

func NewEchoClient(host string) *EchoClient {
	return &EchoClient{host: host}
}

// Do makes call to a service
func (c EchoClient) Do(ctx context.Context, message string) (*MessageEcho, error) {
	var req = struct {
		Message string `json:"message"`
	}{Message: message}

	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(&req); err != nil {
		return nil, err
	}

	u := url.URL{Host: c.host, Path: transport.EchoURL}
	r, err := http.NewRequestWithContext(ctx, "POST", u.String(), body)
	if err != nil {
		return nil, err
	}

	rsp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}

	var echo MessageEcho
	defer rsp.Body.Close()

	if err = json.NewDecoder(rsp.Body).Decode(&echo); err != nil {
		return nil, err
	}

	return &echo, nil
}
