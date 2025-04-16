package config

import (
	"time"
)

type AWS struct {
	Host         string        `config:"host" toml:"host" yaml:"host" json:"host"`
	Region       string        `config:"region" toml:"region" yaml:"region" json:"region"`
	AccessKey    string        `config:"access_key" toml:"access_key" yaml:"access_key" json:"access_key"`
	SecretKey    string        `config:"secret_key" toml:"secret_key" yaml:"secret_key" json:"secret_key"`
	Bucket       string        `config:"bucket" toml:"bucket" yaml:"bucket" json:"bucket"`
	LinkLifeTime time.Duration `config:"link_life_time" toml:"link_life_time" yaml:"link_life_time" json:"link_life_time"`
}

func (c *AWS) copy() *AWS {
	if c == nil {
		return nil
	}

	return &AWS{
		Host:         c.Host,
		Region:       c.Region,
		AccessKey:    c.AccessKey,
		SecretKey:    c.SecretKey,
		Bucket:       c.Bucket,
		LinkLifeTime: c.LinkLifeTime,
	}
}
