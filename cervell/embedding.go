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

// GetWordVector ...
func (cl *Client) GetWordVector(ctx context.Context, word string) ([]float64, error) {
	r, err := postCall[[]float64](ctx, cl, "/embedding/word", &struct {
		Word string `json:"word"`
	}{
		Word: word,
	})
	if err != nil {
		return nil, err
	}
	return *r, nil
}

// GetSentenceVector ...
func (cl *Client) GetSentenceVector(ctx context.Context, sentence string) ([]float64, error) {
	r, err := postCall[[]float64](ctx, cl, "/embedding/sentence", &struct {
		Sentence string `json:"sentence"`
	}{
		Sentence: sentence,
	})
	if err != nil {
		return nil, err
	}
	return *r, nil
}
