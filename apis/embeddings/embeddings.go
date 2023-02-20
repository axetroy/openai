package embeddings

type Embeddings struct {
	apiKey string `json:"-"`
	domain string `json:"-"`
}

type EmbeddingResponse struct {
	Object string         `json:"object"`
	Data   []Embedding    `json:"data"`
	Model  string         `json:"model"`
	Usage  EmbeddingUsage `json:"usage"`
}

type EmbeddingUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

type EmbeddingParams struct {
	Model string `json:"model"`
	Input string `json:"input"`
	User  string `json:"user"`
}

type Embedding struct {
	Object    string `json:"object"`
	Embedding []int  `json:"embedding"`
	Index     int    `json:"index"`
}

func New(apiKey, domain string) *Embeddings {
	return &Embeddings{apiKey: apiKey, domain: domain}
}
