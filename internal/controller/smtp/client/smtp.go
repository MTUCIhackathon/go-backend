package client

import (
	"net/smtp"
	"os"

	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
	smtpclient "github.com/MTUCIhackathon/go-backend/internal/controller"
)

type SMTP struct {
	log  *zap.Logger
	cfg  *config.Config
	smtp *smtp.Client
}

//TODO change SMTP struct

func New(cfg *config.Config, log *zap.Logger) (*SMTP, error) {
	passwordRaw, err := os.ReadFile(cfg.SMTP.Password)
	if err != nil {
		return nil, smtpclient.ErrorReadPassword
	}

	_ = string(passwordRaw)

	log.Debug("get password")

	loginRaw, err := os.ReadFile(cfg.SMTP.Login)
	if err != nil {
		return nil, smtpclient.ErrorReadLogin
	}

	_ = string(loginRaw)

	log.Debug("get login")
	s, _ := smtp.NewClient(nil, "")
	client := &SMTP{
		log:  log,
		cfg:  cfg,
		smtp: s,
	}
	return client, nil
}
