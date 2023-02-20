package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

// docs: https://platform.openai.com/docs/api-reference/files/list
func (f *Files) ListFiles() (*FilesResponse, error) {
	url := fmt.Sprintf("%s/v1/files", f.domain)

	source, err := event_source.NewEventSource[any](url, "GET", http.Header{
		"Authorization": []string{"Bearer " + f.apiKey},
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

	var data FilesResponse

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}
