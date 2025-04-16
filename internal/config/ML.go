package config

import (
	"fmt"
)

type ML struct {
	Host  string `config:"host" toml:"host" yaml:"host" json:"host"`
	Port  int    `config:"port" toml:"port" yaml:"port" json:"port"`
	Route string `config:"route" toml:"route" yaml:"route" json:"route"`
}

func (c *ML) Bind() string {
	return fmt.Sprintf("%s:%d/%s/", c.Host, c.Port, c.Route)
}

func (c *ML) copy() *ML {
	if c == nil {
		return nil
	}

	return &ML{
		Host:  c.Host,
		Port:  c.Port,
		Route: c.Route,
	}
}
