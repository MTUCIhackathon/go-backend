package inmemory

import (
	"sync"

	"github.com/google/uuid"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/cache"
	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

var _ cache.Cache = (*Cache)(nil)

const firstTestLength = 174

type Cache struct {
	mu     sync.RWMutex
	config *config.Cache
	log    *zap.Logger
	data   map[uuid.UUID]dto.Test
	opts   []Option
}

func (c *Cache) Set(key uuid.UUID, test dto.Test) error {
	c.mu.Lock()
	c.data[key] = test
	c.mu.Unlock()
	return nil
}

// TODO think about pointer
func (c *Cache) Get(key uuid.UUID) (dto.Test, error) {
	c.mu.RLock()
	test, ok := c.data[key]
	if !ok {
		return dto.Test{}, ErrNotFound
	}
	c.mu.RUnlock()
	return test, nil
}

func (c *Cache) GetAll() ([]dto.Test, error) {
	c.mu.RLock()
	tests := make([]dto.Test, 0, len(c.data))
	for _, t := range c.data {
		tests = append(tests, t)
	}
	c.mu.RUnlock()
	return tests, nil
}

func (c *Cache) GetKeys() []uuid.UUID {
	c.mu.RLock()
	keys := make([]uuid.UUID, 0, len(c.data))
	for key := range c.data {
		keys = append(keys, key)
	}
	c.mu.RUnlock()
	return keys
}

func New(cfg *config.Config, log *zap.Logger, opts ...Option) (*Cache, error) {
	if cfg == nil {
		return nil, ErrNilConfig
	}

	if log == nil {
		log = zap.NewNop()
		log.Named("cache")
		log.Warn("provided logger is nil, initializing with global logger")
	}

	c := &Cache{
		mu:     sync.RWMutex{},
		config: cfg.Cache,
		log:    log,
		data:   make(map[uuid.UUID]dto.Test, 4),
		opts:   opts,
	}

	return c, c.onStart()
}

func (c *Cache) Close() error {
	return c.onStop()
}

func (c *Cache) onStart() error {
	var err error

	for _, opt := range c.opts {
		err = multierr.Append(err, opt.onStart(c))
	}

	return err
}

func (c *Cache) onStop() error {
	var err error

	for _, opt := range c.opts {
		err = multierr.Append(err, opt.onStop(c))
	}

	return err
}

func (c *Cache) validate() error {
	if c.config == nil {
		return ErrNilConfig
	}

	if c.log == nil {
		c.log = zap.NewNop().Named("cache")
	}

	if c.data == nil {
		c.data = nil
	}

	return nil
}
