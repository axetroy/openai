package openai

import (
	api_completions "github.com/axetroy/openai/apis/completions"
	api_edits "github.com/axetroy/openai/apis/edits"
	api_embeddings "github.com/axetroy/openai/apis/embeddings"
	api_files "github.com/axetroy/openai/apis/files"
	api_models "github.com/axetroy/openai/apis/models"
)

type Client struct {
	apiKey      string `json:"-"`
	Files       *api_files.Files
	Models      *api_models.Models
	Embeddings  *api_embeddings.Embeddings
	Completions *api_completions.Completions
	Edits       *api_edits.Edits
}

const (
	API_DOMAIN = "https://api.openai.com"
)

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:      apiKey,
		Files:       api_files.New(apiKey, API_DOMAIN),
		Models:      api_models.New(apiKey, API_DOMAIN),
		Embeddings:  api_embeddings.New(apiKey, API_DOMAIN),
		Completions: api_completions.New(apiKey, API_DOMAIN),
		Edits:       api_edits.New(apiKey, API_DOMAIN),
	}
}
