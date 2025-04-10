package smtp

import (
	"github.com/MTUCIhackathon/go-backend/internal/config"
	smtpclient "github.com/MTUCIhackathon/go-backend/internal/controller"
	"go.uber.org/zap"
	"net/smtp"
	"os"
)

type SMTP struct {
	log  *zap.Logger
	cfg  *config.Config
	smtp smtp.Auth
}

func New(cfg *config.Config, log *zap.Logger) (*SMTP, error) {
	passwordRaw, err := os.ReadFile(cfg.SMTP.PasswordPath)
	if err != nil {
		return nil, smtpclient.ErrorReadPassword
	}

	password := string(passwordRaw)

	log.Debug("get password")

	loginRaw, err := os.ReadFile(cfg.SMTP.LoginPath)
	if err != nil {
		return nil, smtpclient.ErrorReadLogin
	}

	login := string(loginRaw)

	log.Debug("get login")

	client := &SMTP{
		log:  log,
		cfg:  cfg,
		smtp: smtp.PlainAuth("", login, password, cfg.SMTP.GetSMTPServerAddress()),
	}
	return client, nil
}
