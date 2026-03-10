package cervell

import (
	"context"
)

type dimensionOutput struct {
	Dimension int `json:"dimension"`
}

// GetWordEmbeddingDim ...
func (cl *Client) GetWordEmbeddingDim(ctx context.Context) (int, error) {
	r, err := getCall[dimensionOutput](ctx, cl, "/embedding/word/dimension")
	if err != nil {
		return 0, err
	}
	return r.Dimension, nil
}

// GetSentenceEmbeddingDim ...
func (cl *Client) GetSentenceEmbeddingDim(ctx context.Context) (int, error) {
	r, err := getCall[dimensionOutput](ctx, cl, "/embedding/sentence/dimension")
	if err != nil {
		return 0, err
	}
	return r.Dimension, nil
}
