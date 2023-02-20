package edits

type Edits struct {
	apiKey string `json:"-"`
	domain string `json:"-"`
}

type EditResponse struct {
	Object  string       `json:"object"`
	Created int          `json:"created"`
	Choices []EditChoice `json:"choices"`
	Usage   EditUsage    `json:"usage"`
}

type EditUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type EditChoice struct {
	Text  string `json:"text"`
	Index int    `json:"index"`
}

func New(apiKey, domain string) *Edits {
	return &Edits{apiKey: apiKey, domain: domain}
}
