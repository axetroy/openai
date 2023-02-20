package completions

type Completions struct {
	apiKey string `json:"-"`
	domain string `json:"-"`
}

func New(apiKey, domain string) *Completions {
	return &Completions{apiKey: apiKey, domain: domain}
}

type CompletionParams struct {
	Prompt           *string                `json:"prompt,omitempty"`
	Model            string                 `json:"model"`
	Temperature      *float64               `json:"temperature,omitempty"`
	MaxTokens        *int                   `json:"max_tokens,omitempty"`
	TopP             *int                   `json:"top_p,omitempty"`
	N                *int                   `json:"n,omitempty"`
	Logprobs         *int                   `json:"logprobs,omitempty"`
	Echo             *bool                  `json:"echo,omitempty"`
	Stop             *[]string              `json:"stop,omitempty"`
	PresencePenalty  *int                   `json:"presence_penalty,omitempty"`
	FrequencyPenalty *int                   `json:"frequency_penalty,omitempty"`
	BestOf           *int                   `json:"best_of,omitempty"`
	LogitBias        map[string]interface{} `json:"logit_bias,omitempty"`
	User             *string                `json:"user,omitempty"`
}

type CompletionApiParams struct {
	CompletionParams
	Stream bool `json:"stream"`
}

type CompletionChoice struct {
	Text         string      `json:"text"`
	Index        int         `json:"index"`
	Logprobs     interface{} `json:"logprobs"`
	FinishReason interface{} `json:"finish_reason"`
}

type CompletionResponse struct {
	ID      string             `json:"id"`
	Object  string             `json:"object"`
	Created int                `json:"created"`
	Choices []CompletionChoice `json:"choices"`
	Model   string             `json:"model"`
}
