package tests

import (
	"slurp-server/internal/core/ports"
	"slurp-server/internal/core/usecases"
	"testing"
)

type MockRepository struct {
	PaginationType *string
	AuthType       *string
	OutType        *string
}

func (m MockRepository) AddApiConfiguration(configuration ports.ApiConfiguration) error {
	return nil
}

func (m MockRepository) UpdateApiConfiguration(name string, configuration ports.ApiConfiguration) error {
	return nil
}

func (m MockRepository) DeleteApiConfiguration(name string) error {
	return nil
}

func (m MockRepository) ListApiConfigurations() (*[]ports.ApiConfiguration, error) {
	apis := []ports.ApiConfiguration{}
	apis = append(apis, generateApiConfig(m.PaginationType))
	return &apis, nil
}

func (m MockRepository) GetApiConfiguration(name string) (*ports.ApiConfiguration, error) {
	apis := generateApiConfig(nil)
	return &apis, nil
}

func generateApiConfig(paginationType *string) ports.ApiConfiguration {
	var paginationConfig *ports.PaginationConfiguration
	if paginationType != nil {
		paginationConfig = &ports.PaginationConfiguration{
			PaginationType: *paginationType,
			PageParam:      "page",
			LimitParam:     "limit",
			PageSize:       25,
			NextLinkPath:   "$.next",
		}
	}
	return ports.ApiConfiguration{
		Name:             "name-1",
		Url:              "https://test.com",
		Method:           "GET",
		AuthConfig:       nil,
		PaginationConfig: paginationConfig,
		DataConfig: ports.DataConfiguration{
			DataType: "JSON",
			DataRoot: "$.results",
		},
		AdditionalHeaders:     nil,
		AdditionalQueryParams: nil,
		OutputConfig:          nil,
	}
}

func TestApiList(t *testing.T) {
	apiCrud := usecases.ApiCrud{Repo: MockRepository{
		PaginationType: nil,
		AuthType:       nil,
		OutType:        nil,
	}}
	apis, err := apiCrud.ListApi()
	if err != nil {
		t.Errorf("Unexpected error during the listing of APIs: %v", err)
	}
	resumes := *apis
	if len(resumes) != 1 {
		t.Errorf("Incorrect number of Api resumes. Expected 1, got %d", len(resumes))
	}
	resume := resumes[0]
	if resume.Authenticated {
		t.Errorf("This API is not authenticated")
	}
	if resume.Paginated {
		t.Errorf("This API is not paginated")
	}
	if resume.OutputType != "LOG" {
		t.Errorf("Default output type is not correctly set")
	}
	if resume.DataType != "JSON" {
		t.Errorf("The correct output data type is not set; Expected JSON, for %s", resume.DataType)
	}
}

func TestPaginatedApiList(t *testing.T) {
	paginationType := "PAGE_LIMIT"
	var apiCrud = usecases.ApiCrud{Repo: MockRepository{
		PaginationType: &paginationType,
		AuthType:       nil,
		OutType:        nil,
	}}
	apis, _ := apiCrud.ListApi()
	resumes := *apis
	resume := resumes[0]
	if !resume.Paginated {
		t.Errorf("This API is not paginated")
	}
}
