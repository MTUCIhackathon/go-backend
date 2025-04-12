package config

import "fmt"

type SMTP struct {
	Host     string `config:"host" toml:"host" yaml:"host" json:"host"`
	Port     int    `config:"port" toml:"port" yaml:"port" json:"port"`
	Login    string `config:"login" toml:"login" yaml:"login" json:"login"`
	Password string `config:"password" toml:"password" yaml:"password" json:"password"`
}

// Не ставь сюда указатель!!!
func (smtp SMTP) GetSMTPServerAddress() string {
	/*if smtp ==  {
		return ""
	}*/
	return fmt.Sprintf("%s:%d", smtp.Host, smtp.Port)
}

func (smtp *SMTP) copy() *SMTP {
	if smtp == nil {
		return nil
	}

	return &SMTP{
		Host:     smtp.Host,
		Port:     smtp.Port,
		Login:    smtp.Login,
		Password: smtp.Password,
	}
}
