package ports

type ApiConfiguration struct {
	Name                  string                   `json:"name" validate:"required"`
	Url                   string                   `json:"url" validate:"url"`
	Method                string                   `json:"method" validate:"oneof=GET POST"`
	AuthConfig            *AuthenticationConfig    `json:"auth" validate:"omitempty,dive"`
	PaginationConfig      *PaginationConfiguration `json:"pagination" validate:"omitempty,dive"`
	DataConfig            DataConfiguration        `json:"data" validate:"dive"`
	AdditionalHeaders     map[string]string        `json:"additional_headers"`
	AdditionalQueryParams map[string]string        `json:"additional_queryparams"`
	OutputConfig          *OutputConfig            `json:"output" validate:"omitempty,dive"`
	Active                bool                     `json:"active"`
}

type DataConfiguration struct {
	DataType string `json:"type" validate:"required,oneof=JSON"`
	DataRoot string `json:"root" validate:"required"`
}

type PaginationConfiguration struct {
	PaginationType string `json:"type" validate:"required,oneof=NONE PAGE_LIMIT OFFSET_LIMIT HATEOAS"`
	PageParam      string `json:"page_param"`
	LimitParam     string `json:"limit_param"`
	PageSize       int    `json:"page_size" validate:"gt=0"`
	NextLinkPath   string `json:"next_link_path"`
	BatchSize      int    `json:"batch_size"`
}

type AuthenticationConfig struct {
	AuthType        string `json:"type" validate:"required,oneof=NONE API_KEY CLIENT_CREDS"`
	InHeader        bool   `json:"in_header"`
	TokenSecret     string `json:"token_secret"`
	TokenParam      string `json:"token_param"`
	AccessTokenUrl  string `json:"access_token_url"`
	PayloadTemplate string `json:"payload_template"`
	ClientId        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	AccessTokenPath string `json:"access_token_path"`
}

type OutputConfig struct {
	OutputType string `json:"type" validate:"required,oneof=NONE FILE BUCKET BIGQUERY"`
	FileName   string `json:"filename"`
	BucketName string `json:"bucket"`
	Project    string `json:"project"`
	Dataset    string `json:"dataset"`
	Table      string `json:"table"`
	Autodetect bool   `json:"autodetect"`
}
