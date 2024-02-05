package factory

import (
	"fmt"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]Store)
)

func Register(name string, store Store) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if store == nil {
		panic("store: Register provider is nil")
	}
	if _, exists := providers[name]; exists {
		panic("store: Register called twice for provider " + name)
	}
	providers[name] = store
}

func New(name string) (Store, error) {
	providersMu.RLock()
	store, ok := providers[name]
	if !ok {
		return nil, fmt.Errorf("store: unknown provider %s", name)
	}
	providersMu.RUnlock()
	return store, nil
}
