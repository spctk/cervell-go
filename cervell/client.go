package cervell

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client ...
type Client struct {
	APIKey     string
	httpClient http.Client
}

func getCall[O any](ctx context.Context, cl *Client, path string) (*O, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serviceURL+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code '%s'", resp.Status)
	}
	var out O
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func postCall[O, I any](ctx context.Context, cl *Client, path string, in *I) (*O, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, serviceURL+path, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code '%s'", resp.Status)
	}
	var out O
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}
