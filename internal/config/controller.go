package config

import (
	"fmt"
	"time"
)

type Controller struct {
	Host           string `config:"host" toml:"host" yaml:"host" json:"host" `
	Port           int    `config:"port" toml:"port" yaml:"port" json:"port" `
	TimeoutSeconds int    `config:"timeout_seconds" toml:"timeout_seconds" yaml:"timeout_seconds" json:"timeout_seconds" `
	LogLevel       string `config:"log_level" toml:"log_level" yaml:"log_level" json:"log_level" `
}

func (c *Controller) Bind() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Controller) Timeout() time.Duration {
	if c == nil {
		return 0
	}
	return time.Duration(c.TimeoutSeconds) * time.Second
}

func (c *Controller) copy() *Controller {
	if c == nil {
		return nil
	}

	return &Controller{
		Host:           c.Host,
		Port:           c.Port,
		TimeoutSeconds: c.TimeoutSeconds,
	}
}
