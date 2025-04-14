package config

import (
	"fmt"
	"time"
)

type ML struct {
	Host string `config:"host" toml:"host" yaml:"host" json:"host" `
	Port int    `config:"port" toml:"port" yaml:"port" json:"port" `
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
