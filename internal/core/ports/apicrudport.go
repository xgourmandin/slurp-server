package ports

type ApiCrud interface {
	CreateApi(config ApiConfiguration) error
	UpdateApi(config ApiConfiguration) error
	GetApi(name string) (*ApiConfiguration, error)
	ListApi() (*[]ApiResume, error)
}
