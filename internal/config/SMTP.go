package config

import "fmt"

type SMTP struct {
	Host     string `config:"host" toml:"host" yaml:"host" json:"host"`
	Port     int    `config:"port" toml:"port" yaml:"port" json:"port"`
	Login    string `config:"login" toml:"login" yaml:"login" json:"login"`
	Password string `config:"password" toml:"password" yaml:"password" json:"password"`
}

func (c *SMTP) GetSMTPServerAddress() string {
	if c == nil {
		return ""
	}

	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *SMTP) copy() *SMTP {
	if c == nil {
		return nil
	}

	return &SMTP{
		Host:     c.Host,
		Port:     c.Port,
		Login:    c.Login,
		Password: c.Password,
	}
}
