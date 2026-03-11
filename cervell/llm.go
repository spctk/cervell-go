package cervell

import (
	"context"

	"github.com/phomola/jsonschema-go/jsonschema"
)

type (
	llmPrompt struct {
		Prompt string             `json:"prompt"`
		Schema *jsonschema.Schema `json:"schema,omitempty"`
		Tools  []*LLMTool         `json:"tools"`
	}

	// LLMTool is a tool callable by an LLM.
	LLMTool struct {
		Name        string             `json:"name"`
		Description string             `json:"description"`
		URL         string             `json:"url"`
		Schema      *jsonschema.Schema `json:"schema"`
	}

	// LLMResponse is a response to an LLM prompt.
	LLMResponse struct {
		Text      string         `json:"text"`
		Structure map[string]any `json:"structure"`
	}
)

// LLMPrompt sends a prompt to an LLM.
func (cl *Client) LLMPrompt(ctx context.Context, prompt string, schema *jsonschema.Schema, tools []*LLMTool) (*LLMResponse, error) {
	r, err := postCall[LLMResponse](ctx, cl, "/llm/prompt", &llmPrompt{
		Prompt: prompt,
		Schema: schema,
		Tools:  tools,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}
