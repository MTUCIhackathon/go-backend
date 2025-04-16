package config

import (
	"context"
	"os"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var defaultConfig = &Config{
	JWT: &Token{
		AccessTokenLifeTime:  60 * 2,
		RefreshTokenLifeTime: 60 * 24,
		PublicKeyPath:        "certs/public_key.pem",
		PrivateKeyPath:       "certs/private_key.pem",
		SigningAlgorithm:     SigningAlgorithmRS256,
	},
	SMTP: &SMTP{
		Host:     "smtp.mail.ru",
		Port:     587,
		Login:    "login",
		Password: "password",
	},
	Cache: &Cache{
		CachePath: "test.yaml",
	},
	Postgres: &Postgres{
		Host:             "localhost",
		Port:             5432,
		User:             "postgres",
		Password:         "postgres",
		Database:         "system",
		VersionTableName: "versions",
	},
	AWS: &AWS{
		Host:         "s3",
		Region:       "ru",
		AccessKey:    "access_key",
		SecretKey:    "secret_key",
		Bucket:       "bucket",
		LinkLifeTime: 60 * 3,
	},
	Controller: &Controller{
		Host:           "localhost",
		Port:           8081,
		TimeoutSeconds: 0,
	},
	ML: &ML{
		Host: "localhost",
		Port: 8000,
	},
}

var configPath = os.Getenv("CONFIG_FILE_PATH")

type Config struct {
	JWT        *Token      `config:"jwt" toml:"jwt" yaml:"jwt" json:"jwt"`
	SMTP       *SMTP       `config:"smtp" toml:"smtp" yaml:"smtp" json:"smtp"`
	Cache      *Cache      `config:"cache" toml:"cache" yaml:"cache" json:"cache"`
	Postgres   *Postgres   `config:"postgres" toml:"postgres" yaml:"postgres" json:"postgres"`
	AWS        *AWS        `config:"aws" toml:"aws" yaml:"aws" json:"aws"`
	Controller *Controller `config:"controller" toml:"controller" yaml:"controller" json:"controller"`
	ML         *ML         `config:"ml" toml:"ml" yaml:"ml" json:"ml"`
}

func New() (*Config, error) {
	cfg := defaultConfig.copy()

	l := confita.NewLoader(
		file.NewBackend(configPath),
	)

	err := l.Load(context.Background(), cfg)
	if err != nil {
		return nil, errors.Wrap(err, "error while loading config")
	}

	zap.NewNop().Named("config").Info("loaded config", zap.Any("config", cfg))

	return cfg, nil
}

func (c *Config) copy() *Config {
	return &Config{
		JWT:        c.JWT.copy(),
		SMTP:       c.SMTP.copy(),
		Cache:      c.Cache.copy(),
		Postgres:   c.Postgres.copy(),
		AWS:        c.AWS.copy(),
		Controller: c.Controller.copy(),
		ML:         c.ML.copy(),
	}
}
