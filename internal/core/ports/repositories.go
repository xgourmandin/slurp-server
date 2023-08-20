package ports

type ApiConfigurationRepository interface {
	AddApiConfiguration(configuration ApiConfiguration) error
	UpdateApiConfiguration(name string, configuration ApiConfiguration) error
	DeleteApiConfiguration(name string) error
	ListApiConfigurations() (*[]ApiConfiguration, error)
	GetApiConfiguration(name string) (*ApiConfiguration, error)
}

type ApiHistoryRepository interface {
	AddApiHistory(history ApiHistory) error
	DeleteApiHistory(id string) error
	ListApiHistories() (*[]ApiHistory, error)
}
