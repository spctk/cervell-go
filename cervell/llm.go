package cervell

import (
	"context"

	"github.com/google/jsonschema-go/jsonschema"
)

type (
	llmPrompt struct {
		Prompt string     `json:"prompt"`
		Tools  []*LLMTool `json:"tools"`
	}

	// LLMTool ...
	LLMTool struct {
		Name        string             `json:"name"`
		Description string             `json:"description"`
		URL         string             `json:"url"`
		Schema      *jsonschema.Schema `json:"schema"`
	}

	// LLMResponse ...
	LLMResponse struct {
		Response string `json:"response"`
	}
)

// LLMPrompt ...
func (cl *Client) LLMPrompt(ctx context.Context, prompt string, tools []*LLMTool) (*LLMResponse, error) {
	r, err := postCall[LLMResponse](ctx, cl, "/llm/prompt", &llmPrompt{
		Prompt: prompt,
		Tools:  tools,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}
