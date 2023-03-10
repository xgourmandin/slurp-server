package ports

import "time"

type ApiHistory struct {
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
	Url            string    `json:"url"`
	Method         string    `json:"method"`
	DataType       string    `json:"data"`
	WithAuth       bool      `json:"with_auth"`
	WithPagination bool      `json:"with_pagination"`
	OutputType     string    `json:"output"`
	DataCount      int       `json:"data_count"`
}
