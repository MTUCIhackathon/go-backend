package http

import (
	"context"
	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/controller/smtp"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Controller struct {
	ctx  context.Context
	cfg  *config.Config
	log  *zap.Logger
	prov *token.Provider
	smtp *smtp.SMTPClient
	srv  *echo.Echo
}

func New(ctx context.Context, cfg *config.Config, log *zap.Logger, prov *token.Provider, smtp *smtp.SMTPClient) (*Controller, error) {
	ctrl := &Controller{
		ctx:  ctx,
		cfg:  cfg,
		log:  log,
		prov: prov,
		smtp: smtp,
		srv:  echo.New(),
	}
	ctrl.configureMiddleware()
	ctrl.configureRoutes()
	return ctrl, nil
}

func (ctrl *Controller) configureMiddleware() {
}

func (ctrl *Controller) configureRoutes() {}
