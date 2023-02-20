package files

type Files struct {
	apiKey string `json:"-"`
	domain string `json:"-"`
}

type FilesResponse struct {
	Data   []File `json:"data"`
	Object string `json:"object"`
}

type FileDeleteResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Deleted bool   `json:"deleted"`
}

type File struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     uint64 `json:"bytes"`
	CreatedAt int32  `json:"created_at"`
	FileName  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

func New(apiKey, domain string) *Files {
	return &Files{apiKey: apiKey, domain: domain}
}
