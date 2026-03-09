package cervell

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetWordEmbeddingDim ...
func (cl *Client) GetWordEmbeddingDim(ctx context.Context) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serviceURL+"/embedding/word/dimension", nil)
	if err != nil {
		return 0, err
	}
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code '%s'", resp.Status)
	}
	var out struct {
		Dimension int `json:"dimension"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return 0, err
	}
	return out.Dimension, nil
}

// GetSentenceEmbeddingDim ...
func (cl *Client) GetSentenceEmbeddingDim(ctx context.Context) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, serviceURL+"/embedding/sentence/dimension", nil)
	if err != nil {
		return 0, err
	}
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code '%s'", resp.Status)
	}
	var out struct {
		Dimension int `json:"dimension"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return 0, err
	}
	return out.Dimension, nil
}
