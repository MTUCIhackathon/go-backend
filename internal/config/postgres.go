package config

import "fmt"

type Postgres struct {
	Host     string `config:"host" toml:"host" yaml:"host" json:"host"`
	Port     int    `config:"port" toml:"port" yaml:"port" json:"port"`
	User     string `config:"user" toml:"user" yaml:"user" json:"user"`
	Password string `config:"password" toml:"password" yaml:"password" json:"password"`
	Database string `config:"database" toml:"database" yaml:"database" json:"database"`
}

func (p Postgres) GetDNS() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.Database)
}

func (p *Postgres) copy() *Postgres {
	if p == nil {
		return nil
	}
	return &Postgres{
		Host:     p.Host,
		Port:     p.Port,
		User:     p.User,
		Password: p.Password,
		Database: p.Database,
	}
}
