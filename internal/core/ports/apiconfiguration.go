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
	PaginationType string `json:"type" validate:"required,oneof=PAGE_LIMIT OFFSET_LIMIT HATEOAS"`
	PageParam      string `json:"page_param"`
	LimitParam     string `json:"limit_param"`
	PageSize       int    `json:"page_size"`
	NextLinkPath   string `json:"next_link_path"`
}

type AuthenticationConfig struct {
	AuthType   string `json:"type" validate:"required,oneof=API_KEY"`
	InHeader   bool   `json:"in_header"`
	TokenEnv   string `json:"token_env"`
	TokenParam string `json:"token_param"`
}

type OutputConfig struct {
	OutputType string `json:"type" validate:"required,oneof=FILE BUCKET BIGQUERY"`
	FileName   string `json:"filename"`
	BucketName string `json:"bucket"`
	Project    string `json:"project"`
	Dataset    string `json:"dataset"`
	Table      string `json:"table"`
	Autodetect bool   `json:"autodetect"`
}
