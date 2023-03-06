package usecases

import (
	"github.com/xgourmandin/slurp"
	"github.com/xgourmandin/slurp/configuration"
	"slurp-server/internal/core/ports"
)

type SlurpUseCase struct {
	ApiCrud ApiCrud
}

func (c SlurpUseCase) CreateSlurp(name string) error {
	api, err := c.ApiCrud.GetApi(name)
	if err != nil {
		return nil
	}
	apiConfiguration := createSlurpConfiguration(*api)
	context, err := slurp.NewContextFactory().CreateContextFromConfig(&apiConfiguration)
	if err != nil {
		return err
	}
	engine := slurp.NewSlurpEngine()
	engine.SlurpAPI(*context)
	return nil
}

func createSlurpConfiguration(config ports.ApiConfiguration) configuration.ApiConfiguration {
	authConfig := configuration.AuthenticationConfig{
		AuthType:   config.AuthConfig.AuthType,
		InHeader:   config.AuthConfig.InHeader,
		TokenEnv:   config.AuthConfig.TokenEnv,
		TokenParam: config.AuthConfig.TokenParam,
	}
	paginationConfig := configuration.PaginationConfiguration{
		PaginationType: config.PaginationConfig.PaginationType,
		PageParam:      config.PaginationConfig.PageParam,
		LimitParam:     config.PaginationConfig.LimitParam,
		PageSize:       config.PaginationConfig.PageSize,
		NextLinkPath:   config.PaginationConfig.NextLinkPath,
	}
	dataConfig := configuration.DataConfiguration{
		DataType: config.DataConfig.DataType,
		DataRoot: config.DataConfig.DataRoot,
	}
	outputConfig := configuration.OutputConfig{
		OutputType: config.OutputConfig.OutputType,
		FileName:   config.OutputConfig.FileName,
		BucketName: config.OutputConfig.BucketName,
		Project:    config.OutputConfig.Project,
		Dataset:    config.OutputConfig.Dataset,
		Table:      config.OutputConfig.Table,
		Autodetect: config.OutputConfig.Autodetect,
	}
	return configuration.ApiConfiguration{
		Url:                   config.Url,
		Method:                config.Method,
		AuthConfig:            authConfig,
		PaginationConfig:      paginationConfig,
		DataConfig:            dataConfig,
		AdditionalHeaders:     config.AdditionalHeaders,
		AdditionalQueryParams: config.AdditionalQueryParams,
		OutputConfig:          outputConfig,
	}
}
