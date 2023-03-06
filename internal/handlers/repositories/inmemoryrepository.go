package repositories

import (
	"fmt"
	"slurp-server/internal/core/ports"
)

type InMemoryRepository struct {
	configs []ports.ApiConfiguration
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{configs: []ports.ApiConfiguration{}}
}

func (r *InMemoryRepository) AddApiConfiguration(configuration ports.ApiConfiguration) error {
	r.configs = append(r.configs, configuration)
	return nil
}

func (r *InMemoryRepository) UpdateApiConfiguration(configuration ports.ApiConfiguration) error {
	index := r.indexOf(configuration.Name)
	if index == -1 {
		return fmt.Errorf("not found")
	}
	r.configs = r.removeIndex(index)
	r.AddApiConfiguration(configuration)
	return nil
}

func (r *InMemoryRepository) DeleteApiConfiguration(name string) error {
	index := r.indexOf(name)
	if index == -1 {
		return fmt.Errorf("not found")
	}
	r.configs = r.removeIndex(index)
	return nil
}

func (r *InMemoryRepository) ListApiConfigurations() (*[]ports.ApiConfiguration, error) {
	return &r.configs, nil
}

func (r *InMemoryRepository) GetApiConfiguration(name string) (*ports.ApiConfiguration, error) {
	index := r.indexOf(name)
	if index == -1 {
		return nil, fmt.Errorf("not found")
	}
	return &r.configs[index], nil
}

func (r *InMemoryRepository) indexOf(name string) int {
	for i, c := range r.configs {
		if c.Name == name {
			return i
		}
	}
	return -1
}

func (r *InMemoryRepository) removeIndex(index int) []ports.ApiConfiguration {
	return append(r.configs[:index], r.configs[index+1:]...)
}
