package edits

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

type CreateEditParams struct {
	Model       string   `json:"model"`
	Input       *string  `json:"input,omitempty"`
	Instruction string   `json:"instruction"`
	N           *int     `json:"n,omitempty"`
	Temperature *float64 `json:"temperature,omitempty"`
	TopP        *float64 `json:"top_p,omitempty"`
}

// docs: https://platform.openai.com/docs/api-reference/edits/create
func (this *Edits) CreateEdit(params CreateEditParams) (*EditResponse, error) {
	url := fmt.Sprintf("%s/v1/edits", this.domain)

	jsonBytes, err := json.Marshal(params)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	source, err := event_source.NewEventSource[any](url, "POST", http.Header{
		"Authorization": []string{"Bearer " + this.apiKey},
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
