package ports

type ApiCrud interface {
	CreateApi(config ApiConfiguration)
	UpdateApi(config ApiConfiguration)
	GetApi(name string) ApiConfiguration
	ListApi() []ApiResume
}
