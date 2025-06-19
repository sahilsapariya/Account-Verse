package inmemory

import (
	"sync"

	"server/memorystore/providers/inmemory/stores"
)

type provider struct {
	mutex           sync.Mutex
	stateStore      *stores.StateStore
	envStore        *stores.EnvStore
}

// NewInMemoryStore returns a new in-memory store.
func NewInMemoryProvider() (*provider, error) {
	return &provider{
		mutex:           sync.Mutex{},
		envStore:        stores.NewEnvStore(),
		stateStore:      stores.NewStateStore(),
	}, nil
}