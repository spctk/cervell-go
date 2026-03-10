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
	if cl.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+cl.APIKey)
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
	return bodyCall[O](ctx, cl, http.MethodPost, path, in)
}

func patchCall[O, I any](ctx context.Context, cl *Client, path string, in *I) (*O, error) {
	return bodyCall[O](ctx, cl, http.MethodPatch, path, in)
}

func bodyCall[O, I any](ctx context.Context, cl *Client, method, path string, in *I) (*O, error) {
	b, err := json.Marshal(in)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, serviceURL+path, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	if cl.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+cl.APIKey)
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

func postCallNoResult[I any](ctx context.Context, cl *Client, path string, in *I) error {
	return bodyCallNoResult(ctx, cl, http.MethodPost, path, in)
}

func patchCallNoResult[I any](ctx context.Context, cl *Client, path string, in *I) error {
	return bodyCallNoResult(ctx, cl, http.MethodPatch, path, in)
}

func bodyCallNoResult[I any](ctx context.Context, cl *Client, method, path string, in *I) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, method, serviceURL+path, bytes.NewReader(b))
	if err != nil {
		return err
	}
	if cl.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+cl.APIKey)
	}
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code '%s'", resp.Status)
	}
	return nil
}
