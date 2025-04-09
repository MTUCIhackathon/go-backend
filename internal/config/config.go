package config

import (
	"os"
	"path"
)

type Config struct {
	JWT *Token
}

func New() (*Config, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return &Config{
		JWT: &Token{
			AccessTokenLifeTime:  2,
			RefreshTokenLifeTime: 24,
			PublicKeyPath:        path.Join(wd, "certs", "public_key.pem"),
			PrivateKeyPath:       path.Join(wd, "certs", "private_key.pem"),
		},
	}, nil
}
