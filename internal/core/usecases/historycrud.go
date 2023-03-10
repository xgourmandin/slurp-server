package usecases

import "slurp-server/internal/core/ports"

type HistoryCrud struct {
	Repo ports.ApiHistoryRepository
}

func (a HistoryCrud) ListHistory() (*[]ports.ApiHistory, error) {
	histories, err := a.Repo.ListApiHistories()
	if err != nil {
		return nil, err
	}
	return histories, nil
}
