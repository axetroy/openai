package openai

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"

// 	"github.com/pkg/errors"
// )

// type FineTunesResponse struct {
// 	Data   []File `json:"data"`
// 	Object string `json:"object"`
// }

// type FineTunesParas struct {
// 	TrainingFile string `json:"training_file"`
// }

// type FileDeleteResponse struct {
// 	ID      string `json:"id"`
// 	Object  string `json:"object"`
// 	Deleted bool   `json:"deleted"`
// }

// type File struct {
// 	ID        string `json:"id"`
// 	Object    string `json:"object"`
// 	Bytes     uint64 `json:"bytes"`
// 	CreatedAt int32  `json:"created_at"`
// 	FileName  string `json:"filename"`
// 	Purpose   string `json:"purpose"`
// }

// // docs: https://platform.openai.com/docs/api-reference/fine-tunes/create
// func (c *Client) FineTunesCreate(params FineTunesParas) (*FineTunesResponse, error) {
// 	url := fmt.Sprintf("%s/v1/files", API_DOMAIN)

// 	body, err := json.Marshal(params)

// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	source, err := NewEventSource[any](url, "POST", http.Header{
// 		"Authorization": []string{"Bearer " + c.apiKey},
// 		"Content-Type":  []string{"application/json"},
// 	}, bytes.NewBuffer(body))

// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	defer source.Close()

// 	b, err := io.ReadAll(source.Response().Body)

// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	var data FineTunesResponse

// 	if err := json.Unmarshal(b, &data); err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	return &data, nil
// }

// // docs: https://platform.openai.com/docs/api-reference/files/upload
// func (c *Client) FileUpload(file io.Reader) (*File, error) {
// 	url := fmt.Sprintf("%s/v1/files", API_DOMAIN)

// 	source, err := NewEventSource[any](url, "POST", http.Header{
// 		"Authorization": []string{"Bearer " + c.apiKey},
// 		"Content-Type":  []string{"application/json"},
// 	}, file)

// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	defer source.Close()

// 	b, err := io.ReadAll(source.Response().Body)

// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	var data File

// 	if err := json.Unmarshal(b, &data); err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	return &data, nil
// }

// // docs: https://platform.openai.com/docs/api-reference/files/delete
// func (c *Client) FileDelete(fileId string) (*FileDeleteResponse, error) {
// 	url := fmt.Sprintf("%s/v1/files/%s", API_DOMAIN, fileId)

// 	source, err := NewEventSource[any](url, "DELETE", http.Header{
// 		"Authorization": []string{"Bearer " + c.apiKey},
// 		"Content-Type":  []string{"application/json"},
// 	}, nil)

// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	defer source.Close()

// 	b, err := io.ReadAll(source.Response().Body)

// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	var data FileDeleteResponse

// 	if err := json.Unmarshal(b, &data); err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	return &data, nil
// }

// // docs: https://platform.openai.com/docs/api-reference/files/retrieve-content
// func (c *Client) FileContent(fileId string) ([]byte, error) {
// 	url := fmt.Sprintf("%s/v1/files/%s/content", API_DOMAIN, fileId)

// 	source, err := NewEventSource[any](url, "GET", http.Header{
// 		"Authorization": []string{"Bearer " + c.apiKey},
// 	}, nil)

// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	defer source.Close()

// 	b, err := io.ReadAll(source.Response().Body)

// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	return b, nil
// }
