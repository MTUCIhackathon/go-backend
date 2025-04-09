package config

import "fmt"

type SMTP struct {
	Server       string
	Port         int
	LoginPath    string
	PasswordPath string
}

func (smtp SMTP) GetSMTPServerAddress() string {
	return fmt.Sprintf("%s:%d", smtp.Server, smtp.Port)
}
