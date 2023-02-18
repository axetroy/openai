package openai

type Client struct {
	apiKey string `json:"-"`
}

const (
	API_DOMAIN = "https://api.openai.com"
)

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}
