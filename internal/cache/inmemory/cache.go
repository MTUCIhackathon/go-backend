package inmemory

import (
	"sync"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

type Cache struct {
	data   sync.Map
	config *config.Cache
}

func New(config *config.Cache) *Cache {
	return &Cache{
		data:   sync.Map{},
		config: config,
	}
}
