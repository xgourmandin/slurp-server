package usecases

import (
	"github.com/xgourmandin/slurp"
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

func createSlurpConfiguration(configuration ports.ApiConfiguration) slurp.ApiConfiguration {
	authConfig := slurp.AuthenticationConfig{
		AuthType:   configuration.AuthConfig.AuthType,
		InHeader:   configuration.AuthConfig.InHeader,
		TokenEnv:   configuration.AuthConfig.TokenEnv,
		TokenParam: configuration.AuthConfig.TokenParam,
	}
	paginationConfig := slurp.PaginationConfiguration{
		PaginationType: configuration.PaginationConfig.PaginationType,
		PageParam:      configuration.PaginationConfig.PageParam,
		LimitParam:     configuration.PaginationConfig.LimitParam,
		PageSize:       configuration.PaginationConfig.PageSize,
		NextLinkPath:   configuration.PaginationConfig.NextLinkPath,
	}
	dataConfig := slurp.DataConfiguration{
		DataType: configuration.DataConfig.DataType,
		DataRoot: configuration.DataConfig.DataRoot,
	}
	outputConfig := slurp.OutputConfig{
		OutputType: configuration.OutputConfig.OutputType,
		FileName:   configuration.OutputConfig.FileName,
		BucketName: configuration.OutputConfig.BucketName,
		Project:    configuration.OutputConfig.Project,
		Dataset:    configuration.OutputConfig.Dataset,
		Table:      configuration.OutputConfig.Table,
		Autodetect: configuration.OutputConfig.Autodetect,
	}
	return slurp.ApiConfiguration{
		Url:                   configuration.Url,
		Method:                configuration.Method,
		AuthConfig:            authConfig,
		PaginationConfig:      paginationConfig,
		DataConfig:            dataConfig,
		AdditionalHeaders:     configuration.AdditionalHeaders,
		AdditionalQueryParams: configuration.AdditionalQueryParams,
		OutputConfig:          outputConfig,
	}
}
