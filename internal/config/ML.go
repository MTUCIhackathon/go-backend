package config

import (
	"fmt"
)

type ML struct {
	Host string `config:"host" toml:"host" yaml:"host" json:"host" `
	Port int    `config:"port" toml:"port" yaml:"port" json:"port" `
}

func (c *ML) Bind() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *ML) copy() *ML {
	if c == nil {
		return nil
	}

	return &ML{
		Host: c.Host,
		Port: c.Port,
	}
}
