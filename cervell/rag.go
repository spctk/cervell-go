package cervell

import (
	"context"

	"github.com/google/uuid"
)

// Document ...
type Document struct {
	ID     uuid.UUID `json:"id"`
	Text   string    `json:"text"`
	Group  string    `json:"group"`
	Vector []float64 `json:"vector,omitempty"`
}

// InsertDocument ...
func (cl *Client) InsertDocument(ctx context.Context, doc *Document) error {
	return postCallNoResult(ctx, cl, "/documents", doc)
}

// UpdateDocument ...
func (cl *Client) UpdateDocument(ctx context.Context, doc *Document) error {
	return patchCallNoResult(ctx, cl, "/documents", doc)
}

// FindDocuments ...
func (cl *Client) FindDocuments(ctx context.Context, group, text string, limit int) ([]*Document, error) {
	r, err := postCall[struct {
		Documents []*Document `json:"documents"`
	}](ctx, cl, "/documents/similar", &struct {
		Text  string `json:"text"`
		Group string `json:"group"`
		Limit int    `json:"limit"`
	}{
		Text:  text,
		Group: group,
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}
	return r.Documents, nil
}
