package repositories

import (
	"fmt"
	"slurp-server/internal/core/ports"
)

type InMemoryHistoryRepository struct {
	configs []ports.ApiHistory
}

func NewInMemoryHistoryRepository() *InMemoryHistoryRepository {
	return &InMemoryHistoryRepository{configs: []ports.ApiHistory{}}
}

func (r *InMemoryHistoryRepository) AddApiHistory(configuration ports.ApiHistory) error {
	r.configs = append(r.configs, configuration)
	return nil
}

func (r *InMemoryHistoryRepository) DeleteApiHistory(name string) error {
	index := r.indexOf(name)
	if index == -1 {
		return fmt.Errorf("not found")
	}
	r.configs = r.removeIndex(index)
	return nil
}

func (r *InMemoryHistoryRepository) ListApiHistories() (*[]ports.ApiHistory, error) {
	return &r.configs, nil
}

func (r *InMemoryHistoryRepository) indexOf(name string) int {
	for i, c := range r.configs {
		if c.Name == name {
			return i
		}
	}
	return -1
}

func (r *InMemoryHistoryRepository) removeIndex(index int) []ports.ApiHistory {
	return append(r.configs[:index], r.configs[index+1:]...)
}
