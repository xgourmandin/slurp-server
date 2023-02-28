package usecases

import "slurp-server/internal/core/ports"

type ApiCrud struct {
	Repo ports.ApiConfigurationRepository
}

func (a ApiCrud) CreateApi(config ports.ApiConfiguration) error {
	return a.Repo.AddApiConfiguration(config)
}

func (a ApiCrud) UpdateApi(config ports.ApiConfiguration) error {
	return a.Repo.UpdateApiConfiguration(config)
}

func (a ApiCrud) GetApi(name string) (*ports.ApiConfiguration, error) {
	return a.Repo.GetApiConfiguration(name)
}

func (a ApiCrud) ListApi() (*[]ports.ApiResume, error) {
	configurations, err := a.Repo.ListApiConfigurations()
	if err != nil {
		return nil, err
	}
	var resumes []ports.ApiResume
	for _, apiconf := range *configurations {
		outputType := "LOG"
		if apiconf.OutputConfig != nil {
			outputType = apiconf.OutputConfig.OutputType
		}
		resumes = append(resumes, ports.ApiResume{
			Name:          apiconf.Name,
			Url:           apiconf.Url,
			Method:        apiconf.Method,
			DataType:      apiconf.DataConfig.DataType,
			Paginated:     apiconf.PaginationConfig != nil,
			Authenticated: apiconf.AuthConfig != nil,
			OutputType:    outputType,
		})
	}
	return &resumes, nil
}
