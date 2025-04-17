package inmemory

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

func TestCache_New(t *testing.T) {
	cfg := &config.Config{
		Cache: &config.Cache{
			CachePath: os.Getenv("CACHE_PATH"),
		},
	}

	if cfg.Cache.CachePath == "" {
		t.Skip("cache path env not set")
	}

	log, _ := zap.NewProduction()

	c, err := New(cfg, log, WithLoader())
	require.NoError(t, err)
	require.NotNil(t, c)

	keys := c.GetKeys()
	t1, err := c.Get(keys[0])
	assert.NoError(t, err)
	t2, err := c.Get(keys[1])
	assert.NoError(t, err)
	require.Len(t, keys, 2)

	t.Log(t1.ID, t1.Name, t1.Description, t1.Questions)
	t.Log(t2.ID, t2.Name, t2.Description, t2.Questions)

	t.Log(c.GetAll())
}
