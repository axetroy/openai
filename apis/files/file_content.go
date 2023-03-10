package files

import (
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

// docs: https://platform.openai.com/docs/api-reference/files/retrieve-content
func (this *Files) RetrieveContent(fileId string) ([]byte, error) {
	url := fmt.Sprintf("%s/v1/files/%s/content", this.domain, fileId)

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

	return b, nil
}
