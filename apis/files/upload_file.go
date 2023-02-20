package files

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

// docs: https://platform.openai.com/docs/api-reference/files/upload
func (f *Files) UploadFile(file io.Reader) (*File, error) {
	url := fmt.Sprintf("%s/v1/files", f.domain)

	source, err := event_source.NewEventSource[any](url, "POST", http.Header{
		"Authorization": []string{"Bearer " + f.apiKey},
	}, file)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer source.Close()

	b, err := io.ReadAll(source.Response().Body)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var data File

	if err := json.Unmarshal(b, &data); err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}
