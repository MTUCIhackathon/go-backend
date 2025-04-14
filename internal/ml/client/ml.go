package ml

import (
	"go.uber.org/zap"
	"gopkg.in/resty.v1"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

type PythonClient struct {
	cfg *config.Config
	log *zap.Logger
	cli *resty.Client
}

func New(cfg *config.Config, log *zap.Logger) (*PythonClient, error) {
	if cfg == nil {
		return nil, errNilConfig
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

	c.cli.HostURL = cfg.ML.Bind()

	return c, nil
}
