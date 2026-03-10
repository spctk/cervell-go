package cervell

import (
	"context"
)

type dimensionOutput struct {
	Dimension int `json:"dimension"`
}

// GetWordEmbeddingDim returns the vector dimension for the word embedding.
func (cl *Client) GetWordEmbeddingDim(ctx context.Context) (int, error) {
	r, err := getCall[dimensionOutput](ctx, cl, "/embedding/word/dimension")
	if err != nil {
		return 0, err
	}
	return r.Dimension, nil
}

// GetSentenceEmbeddingDim returns the vector dimension for the sentence embedding.
func (cl *Client) GetSentenceEmbeddingDim(ctx context.Context) (int, error) {
	r, err := getCall[dimensionOutput](ctx, cl, "/embedding/sentence/dimension")
	if err != nil {
		return 0, err
	}
	return r.Dimension, nil
}

// GetWordVector returns the vector of the provided word.
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

// GetSentenceVector returns the vector of the provided sentence.
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
