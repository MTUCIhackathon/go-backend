package config

import (
	"fmt"
)

const (
	defaultLogLevel = "none"
)

type Postgres struct {
	Host             string `config:"host" toml:"host" yaml:"host" json:"host"`
	Port             int    `config:"port" yaml:"port" toml:"port" json:"port"`
	User             string `config:"user" yaml:"user" toml:"user" json:"user"`
	Password         string `config:"password" yaml:"password" toml:"password" json:"password"`
	Database         string `config:"database" yaml:"database" toml:"database" json:"database"`
	VersionTableName string `config:"versions" yaml:"versions" toml:"versions"  json:"versions"`
	RunMigrations    bool   `config:"run_migrations" yaml:"run_migrations" json:"run_migrations"`
}

func (c *Postgres) GetURI() string {
	if c == nil {
		return ""
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.User, c.Password, c.Host, c.Port, c.Database,
	)
}

func (c *Postgres) copy() *Postgres {
	if c == nil {
		return nil
	}

	return &Postgres{
		Host:             c.Host,
		Port:             c.Port,
		User:             c.User,
		Password:         c.Password,
		Database:         c.Database,
		VersionTableName: c.VersionTableName,
		RunMigrations:    c.RunMigrations,
	}
}
