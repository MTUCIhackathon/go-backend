package config

import "fmt"

type SMTP struct {
	Host     string `config:"host" toml:"host" yaml:"host" json:"host"`
	Port     int    `config:"port" toml:"port" yaml:"port" json:"port"`
	Login    string `config:"login" toml:"login" yaml:"login" json:"login"`
	Password string `config:"password" toml:"password" yaml:"password" json:"password"`
}

func (smtp SMTP) GetSMTPServerAddress() string {
	return fmt.Sprintf("%s:%d", smtp.Host, smtp.Port)
}
