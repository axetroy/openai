package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

// docs: https://platform.openai.com/docs/api-reference/files/delete
func (f *Files) DeleteFile(fileId string) (*FileDeleteResponse, error) {
	url := fmt.Sprintf("%s/v1/files/%s", f.domain, fileId)

	source, err := event_source.NewEventSource[any](url, "DELETE", http.Header{
		"Authorization": []string{"Bearer " + f.apiKey},
	}, nil)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer source.Close()

	b, err := io.ReadAll(source.Response().Body)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var data FileDeleteResponse

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}
