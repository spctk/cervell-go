package cervell

import (
	"context"
)

// GetWordEmbeddingDim ...
func (cl *Client) GetWordEmbeddingDim(ctx context.Context) (int, error) {
	x, err := getCall[struct {
		Dimension int `json:"dimension"`
	}](ctx, cl, "/embedding/word/dimension")
	if err != nil {
		return 0, err
	}
	return x.Dimension, nil
}

// GetSentenceEmbeddingDim ...
func (cl *Client) GetSentenceEmbeddingDim(ctx context.Context) (int, error) {
	x, err := getCall[struct {
		Dimension int `json:"dimension"`
	}](ctx, cl, "/embedding/sentence/dimension")
	if err != nil {
		return 0, err
	}
	return x.Dimension, nil
}
