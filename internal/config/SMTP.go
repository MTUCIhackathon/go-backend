package config

import "fmt"

type SMTP struct {
	Host         string `config:"smtp_host"`
	Port         int    `config:"smtp_port"`
	LoginPath    string `config:"smtp_login_path"`
	PasswordPath string `config:"smtp_password_path"`
}

func (smtp SMTP) GetSMTPServerAddress() string {
	return fmt.Sprintf("%s:%d", smtp.Host, smtp.Port)
}
