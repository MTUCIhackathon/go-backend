package inmemory

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

func TestCache_New(t *testing.T) {
	cfg := &config.Cache{
		CachePath: os.Getenv("CACHE_PATH"),
	}

	if cfg.CachePath == "" {
		t.Skip("cache path env not set")
	}

	log, _ := zap.NewProduction()

	c, err := New(cfg, log, WithLoader())
	require.NoError(t, err)
	require.NotNil(t, c)
}
