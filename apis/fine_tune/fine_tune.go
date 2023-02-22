package fine_tune

// TODO: to be done

type FineTune struct {
	apiKey string `json:"-"`
	domain string `json:"-"`
}

type CreateFineTuneResponse struct {
	Object string           `json:"object"`
	Data   []FineTuneObject `json:"data"`
}

type FileDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type FineTuneObject struct {
	ID             string  `json:"id"`
	Object         string  `json:"object"`
	Model          string  `json:"model"`
	CreatedAt      int32   `json:"created_at"`
	FineTunedModel *string `json:"fine_tuned_model"`
	FileName       string  `json:"filename"`
	Purpose        string  `json:"purpose"`
}

func New(apiKey, domain string) *FineTune {
	return &FineTune{apiKey: apiKey, domain: domain}
}
