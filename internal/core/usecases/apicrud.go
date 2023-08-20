package usecases

import "slurp-server/internal/core/ports"

type ApiCrud struct {
	Repo ports.ApiConfigurationRepository
}

func (a ApiCrud) CreateApi(config ports.ApiConfiguration) error {
	config.Active = true
	if config.AuthConfig == nil {
		config.AuthConfig = &ports.AuthenticationConfig{AuthType: "NONE"}
	}
	if config.PaginationConfig == nil {
		config.PaginationConfig = &ports.PaginationConfiguration{PaginationType: "NONE"}
	}
	if config.OutputConfig == nil {
		config.OutputConfig = &ports.OutputConfig{OutputType: "LOG"}
	}
	return a.Repo.AddApiConfiguration(config)
}

func (a ApiCrud) UpdateApi(config ports.ApiConfiguration, name string) error {
	return a.Repo.UpdateApiConfiguration(name, config)
}

func (a ApiCrud) DeleteApi(name string) error {
	return a.Repo.DeleteApiConfiguration(name)
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
			Paginated:     apiconf.PaginationConfig != nil && apiconf.PaginationConfig.PaginationType != "NONE",
			Authenticated: apiconf.AuthConfig != nil && apiconf.AuthConfig.AuthType != "NONE",
			OutputType:    outputType,
			Active:        apiconf.Active,
		})
	}
	return &resumes, nil
}

func (a ApiCrud) PauseApi(name string) error {
	return a.toggleSlurp(name, false)
}

func (a ApiCrud) UnPauseApi(name string) error {
	return a.toggleSlurp(name, true)
}

func (a ApiCrud) toggleSlurp(name string, activate bool) error {
	api, err := a.Repo.GetApiConfiguration(name)
	if err != nil {
		return err
	}
	api.Active = activate
	err = a.Repo.UpdateApiConfiguration("", *api)
	if err != nil {
		return err
	}
	return nil
}
