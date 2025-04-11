package inmemory

import (
	"sync"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

type Cache struct {
	mu     sync.RWMutex
	data   map[string]string
	config *config.Cache
}

/*func New(config *config.Cache, opts ...Options) *Cache {
	readed, err := os.ReadFile("")

	m := make(map[string]string)

	return &Cache{
		mu:     sync.RWMutex{},
		data:   m,
		config: config,
	}
}*/
