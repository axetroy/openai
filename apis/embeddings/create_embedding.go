package embeddings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

type CreateEmbeddingParams struct {
	Model string  `json:"model"`
	Input string  `json:"input"`
	User  *string `json:"user,omitempty"`
}

// docs: https://platform.openai.com/docs/api-reference/embeddings/create
func (this *Embeddings) CreateEmbedding(params CreateEmbeddingParams) (*EmbeddingResponse, error) {
	url := fmt.Sprintf("%s/v1/embeddings", this.domain)

	body, err := json.Marshal(params)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	source, err := event_source.NewEventSource[any](url, "POST", http.Header{
		"Authorization": []string{"Bearer " + this.apiKey},
		"Content-Type":  []string{"application/json"},
	}, bytes.NewBuffer(body))

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
