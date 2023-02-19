package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type EditResponse struct {
	Object  string       `json:"object"`
	Created int          `json:"created"`
	Choices []EditChoice `json:"choices"`
	Usage   EditUsage    `json:"usage"`
}

type EditUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type EditChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

type EditParams struct {
	Model       string   `json:"model"`
	Input       *string  `json:"input,omitempty"`
	Instruction string   `json:"instruction"`
	N           *int     `json:"n,omitempty"`
	Temperature *float64 `json:"temperature,omitempty"`
	TopP        *float64 `json:"top_p,omitempty"`
}

// docs: https://platform.openai.com/docs/api-reference/edits
func (c *Client) Edits(params EditParams) (*EditResponse, error) {
	url := fmt.Sprintf("%s/v1/edits", API_DOMAIN)

	jsonBytes, err := json.Marshal(params)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	source, err := NewEventSource[any](url, "POST", http.Header{
		"Authorization": []string{"Bearer " + c.apiKey},
		"Content-Type":  []string{"application/json"},
	}, bytes.NewBuffer(jsonBytes))

	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer source.Close()

	b, err := io.ReadAll(source.Response().Body)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var data EditResponse

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}
