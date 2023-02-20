package event_source

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/axetroy/openai/pkg/iterable"
	"github.com/pkg/errors"
)

type EventSource[T any] struct {
	*iterable.Iterable[*T]
	resp *http.Response
}

func NewEventSource[T any](url string, method string, headers http.Header, body io.Reader) (*EventSource[T], error) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for key, values := range headers {
		for _, value := range values {
			req.Header.Set(key, value)
		}
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	reader := bufio.NewReader(resp.Body)

	iterator := iterable.NewIterable(func() (*T, bool, error) {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			return nil, false, errors.WithStack(err)
		}

		lineText := strings.TrimSpace(string(line))

		if !strings.HasPrefix(lineText, "data: ") {
			return nil, false, nil
		}

		dataRaw := strings.TrimPrefix(lineText, "data: ")

		if dataRaw == "[DONE]" {
			return nil, true, nil
		}

		var chunk T

		if err := json.Unmarshal([]byte(dataRaw), &chunk); err != nil {
			return nil, false, errors.WithStack(err)
		}

		return &chunk, false, nil
	})

	return &EventSource[T]{
		Iterable: iterator,
		resp:     resp,
	}, nil
}

func (e *EventSource[T]) Response() *http.Response {
	return e.resp
}

func (e *EventSource[T]) Close() {
	e.resp.Body.Close()
}
