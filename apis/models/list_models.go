package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

// docs: https://platform.openai.com/docs/api-reference/files/upload
func (this *Models) ListModels() (*ModelResponse, error) {
	url := fmt.Sprintf("%s/v1/models", this.domain)

	source, err := event_source.NewEventSource[any](url, "GET", http.Header{
		"Authorization": []string{"Bearer " + this.apiKey},
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
