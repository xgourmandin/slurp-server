package ports

type ApiResume struct {
	Name          string `json:"name"`
	Url           string `json:"url"`
	Method        string `json:"method"`
	DataType      string `json:"data_type"`
	Paginated     bool   `json:"paginated"`
	Authenticated bool   `json:"authenticated"`
	OutputType    string `json:"output_type"`
}
