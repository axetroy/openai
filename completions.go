package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// docs: https://platform.openai.com/docs/api-reference/completions

type CompletionParams struct {
	Prompt           *string                `json:"prompt,omitempty"`
	Model            string                 `json:"model"`
	Temperature      *float64               `json:"temperature,omitempty"`
	MaxTokens        *int                   `json:"max_tokens,omitempty"`
	TopP             *int                   `json:"top_p,omitempty"`
	N                *int                   `json:"n,omitempty"`
	Logprobs         *int                   `json:"logprobs,omitempty"`
	Echo             *bool                  `json:"echo,omitempty"`
	Stop             *[]string              `json:"stop,omitempty"`
	PresencePenalty  *int                   `json:"presence_penalty,omitempty"`
	FrequencyPenalty *int                   `json:"frequency_penalty,omitempty"`
	BestOf           *int                   `json:"best_of,omitempty"`
	LogitBias        map[string]interface{} `json:"logit_bias,omitempty"`
	User             *string                `json:"user,omitempty"`
}

type CompletionApiParams struct {
	CompletionParams
	Stream bool `json:"stream"`
}

type CompletionChoice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason interface{} `json:"finish_reason"`
}

type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int                `json:"created"`
	Choices []CompletionChoice `json:"choices"`
	Model   string             `json:"model"`
}

func (c *Client) CompletionsStream(params CompletionParams, writer io.Writer) error {
	url := fmt.Sprintf("%s/v1/completions", API_DOMAIN)

	apiParams := CompletionApiParams{
		CompletionParams: params,
		Stream:           true,
	}

	body, err := json.Marshal(apiParams)

	if err != nil {
		return errors.WithStack(err)
	}

	source, err := NewEventSource[CompletionResponse](url, "POST", http.Header{
		"Authorization": []string{"Bearer " + c.apiKey},
		"Content-Type":  []string{"application/json"},
	}, bytes.NewBuffer(body))

	if err != nil {
		return errors.WithStack(err)
	}

	defer source.Close()

	for {
		chunk, done, err := source.Next()

		if done {
			break
		}

		if err != nil {
			return errors.WithStack(err)
		}

		if chunk == nil {
			continue
		}

		writer.Write([]byte(chunk.Choices[0].Text))
	}

	return nil
}

func (c *Client) Completions(config CompletionParams) ([]byte, error) {
	var writer bytes.Buffer

	if err := c.CompletionsStream(config, &writer); err != nil {
		return nil, errors.WithStack(err)
	}

	return writer.Bytes(), nil
}
