package completions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/axetroy/openai/pkg/event_source"
	"github.com/pkg/errors"
)

// docs: https://platform.openai.com/docs/api-reference/completions
func (c *Completions) CreateCompletionsStream(params CompletionParams, writer io.Writer) error {
	url := fmt.Sprintf("%s/v1/completions", c.domain)

	apiParams := CompletionApiParams{
		CompletionParams: params,
		Stream:           true,
	}

	body, err := json.Marshal(apiParams)

	if err != nil {
		return errors.WithStack(err)
	}

	source, err := event_source.NewEventSource[CompletionResponse](url, "POST", http.Header{
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

		if _, err = writer.Write([]byte(chunk.Choices[0].Text)); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// docs: https://platform.openai.com/docs/api-reference/completions
func (c *Completions) CreateCompletions(config CompletionParams) ([]byte, error) {
	var writer bytes.Buffer

	if err := c.CreateCompletionsStream(config, &writer); err != nil {
		return nil, errors.WithStack(err)
	}

	return writer.Bytes(), nil
}
