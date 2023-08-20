package ports

type ApiCrud interface {
	CreateApi(config ApiConfiguration) error
	UpdateApi(config ApiConfiguration, name string) error
	DeleteApi(name string) error
	GetApi(name string) (*ApiConfiguration, error)
	ListApi() (*[]ApiResume, error)
	PauseApi(name string) error
	UnPauseApi(name string) error
}
