package openai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type EmbeddingResponse struct {
	Object string         `json:"object"`
	Data   []Embedding    `json:"data"`
	Model  string         `json:"model"`
	Usage  EmbeddingUsage `json:"usage"`
}

type EmbeddingUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type User interface {
	string | []string
}

type EmbeddingParams struct {
	Model string `json:"model"`
	Input string `json:"input"`
	User  string `json:"user"`
}

type Embedding struct {
	Object    string `json:"object"`
	Embedding []int  `json:"embedding"`
	Index     int    `json:"index"`
}

// docs: https://platform.openai.com/docs/api-reference/embeddings
func (c *Client) Embeddings() (*EmbeddingResponse, error) {
	url := fmt.Sprintf("%s/v1/models", API_DOMAIN)

	source, err := NewEventSource[any](url, "GET", http.Header{
		"Authorization": []string{"Bearer " + c.apiKey},
		"Content-Type":  []string{"application/json"},
	}, nil)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer source.Close()

	b, err := io.ReadAll(source.Response().Body)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var data EmbeddingResponse

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}
