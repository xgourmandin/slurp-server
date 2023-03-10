package ports

type HistoryCrud interface {
	ListHistory() (*[]ApiHistory, error)
}
