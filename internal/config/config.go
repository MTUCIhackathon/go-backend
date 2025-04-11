package config

import (
	"context"
	"os"
	"path"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/pkg/errors"
)

var currentDir, _ = os.Executable()

var defaultConfig = &Config{
	JWT: &Token{
		AccessTokenLifeTime:  60 * 2,
		RefreshTokenLifeTime: 60 * 24,
		PublicKeyPath:        path.Join(currentDir, "certs", "public_key.pem"),
		PrivateKeyPath:       path.Join(currentDir, "certs", "private_key.pem"),
		SigningAlgorithm:     SigningAlgorithmRS256,
	},
	SMTP: &SMTP{
		Host:     "smtp.mail.ru",
		Port:     587,
		Login:    "login",
		Password: "password",
	},
}

type Config struct {
	JWT   *Token `config:"jwt" toml:"jwt" yaml:"jwt" json:"jwt"`
	SMTP  *SMTP  `config:"smtp" toml:"smtp" yaml:"smtp" json:"smtp"`
	Cache *Cache `config:"cache" toml:"cache" yaml:"cache" json:"cache"`
}

func New() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	cfg := defaultConfig.copy()

	l := confita.NewLoader(
		file.NewBackend(path.Join(wd, "config.toml")),
		file.NewBackend(path.Join(wd, "config.yaml")),
		file.NewBackend(path.Join(wd, "config.json")),
	)

	err = l.Load(context.Background(), cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error while loading config")
	}

	return cfg, nil
}

func (c *Config) copy() *Config {
	return &Config{
		JWT:   c.JWT.copy(),
		SMTP:  c.SMTP.copy(),
		Cache: c.Cache.copy(),
	}
}
