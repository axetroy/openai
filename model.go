package openai

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// docs: https://platform.openai.com/docs/api-reference/models

type ModelResponse struct {
	Data   []Model `json:"data"`
	Object string  `json:"object"`
}

type Model struct {
	ID         string          `json:"id"`
	Object     string          `json:"object"`
	OwnerBy    string          `json:"owner_by"`
	Permission ModelPermission `json:"permission"`
	Created    int             `json:"created"`
}

type ModelPermission struct {
	ID                 string  `json:"id"`
	Object             string  `json:"object"`
	Created            int     `json:"created"`
	AllowCreateEngine  bool    `json:"allow_create_engine"`
	AllowSampling      bool    `json:"allow_sampling"`
	AllowLogprobs      bool    `json:"allow_logprobs"`
	AllowSearchIndices bool    `json:"allow_search_indices"`
	AllowView          bool    `json:"allow_view"`
	AllowFineTuning    bool    `json:"allow_fine_tuning"`
	Organization       string  `json:"organization"`
	Group              *string `json:"group"`
	IsBlocking         bool    `json:"is_blocking"`
}

func (c *Client) Models() (*ModelResponse, error) {
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

	var data ModelResponse

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}

func (c *Client) Model(model string) (*Model, error) {
	url := fmt.Sprintf("%s/v1/model/%s", API_DOMAIN, model)

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

	var data Model

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}
