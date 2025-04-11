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
	LogLevel         string `config:"log-level" yaml:"log_level" toml:"log_level" json:"log_level"`
	VersionTableName string `config:"versions" yaml:"versions" toml:"versions"  json:"versions"`
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

func (c *Postgres) GetLogLevel() string {
	if c == nil {
		return defaultLogLevel
	}

	switch c.LogLevel {
	case "trace", "TRACE":
		return "trace"
	case "debug", "DEBUG":
		return "debug"
	case "info", "INFO":
		return "info"
	case "warn", "warning", "WARN", "WARNING":
		return "warn"
	case "error", "ERROR":
		return "error"
	case "none", "NONE":
		return "none"
	default:
		return defaultLogLevel
	}
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
		LogLevel:         c.LogLevel,
		VersionTableName: c.VersionTableName,
	}
}
