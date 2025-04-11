package inmemory

import (
	"sync"

	"github.com/google/uuid"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
)

const firstTestLength = 174

type Cache struct {
	mu     sync.RWMutex
	config *config.Cache
	log    *zap.Logger
	data   map[uuid.UUID]dto.Test
	opts   []Option
}

func New(cfg *config.Cache, log *zap.Logger, opts ...Option) (*Cache, error) {
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
		config: cfg,
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
