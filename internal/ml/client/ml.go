package client

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/ml"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

type PythonClient struct {
	cfg *config.Config
	log *zap.Logger
	cli *resty.Client
}

func New(cfg *config.Config, log *zap.Logger) (*PythonClient, error) {
	if cfg == nil {
		return nil, ml.ErrNilConfig
	}

	if log == nil {
		log = zap.NewNop()
		log.Named("python-client")
		log.Warn("provided nil logger: creating with global logger")
	}

	c := &PythonClient{
		cfg: cfg,
		log: log,
		cli: resty.New(),
	}

	return c, nil
}
