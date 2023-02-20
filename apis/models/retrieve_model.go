package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

type RetrieveModelResponse struct {
	ID         string          `json:"id"`
	Object     string          `json:"object"`
	OwnerBy    string          `json:"owner_by"`
	Permission ModelPermission `json:"permission"`
}

// docs: https://platform.openai.com/docs/api-reference/models/list
func (this *Models) RetrieveModel(model string) (*RetrieveModelResponse, error) {
	url := fmt.Sprintf("%s/v1/model/%s", this.domain, model)

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

	var data RetrieveModelResponse

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}
