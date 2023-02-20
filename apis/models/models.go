package models

type Models struct {
	apiKey string `json:"-"`
	domain string `json:"-"`
}

type ModelResponse struct {
	Data   []Model `json:"data"`
	Object string  `json:"object"`
}

type Model struct {
	ID         string          `json:"id"`
	Object     string          `json:"object"`
	OwnerBy    string          `json:"owner_by"`
	Permission ModelPermission `json:"permission"`
	Created    int             `json:"created"`
}

type ModelPermission struct {
	ID                 string  `json:"id"`
	Object             string  `json:"object"`
	Created            int     `json:"created"`
	AllowCreateEngine  bool    `json:"allow_create_engine"`
	AllowSampling      bool    `json:"allow_sampling"`
	AllowLogprobs      bool    `json:"allow_logprobs"`
	AllowSearchIndices bool    `json:"allow_search_indices"`
	AllowView          bool    `json:"allow_view"`
	AllowFineTuning    bool    `json:"allow_fine_tuning"`
	Organization       string  `json:"organization"`
	Group              *string `json:"group"`
	IsBlocking         bool    `json:"is_blocking"`
}

func New(apiKey, domain string) *Models {
	return &Models{apiKey: apiKey, domain: domain}
}
