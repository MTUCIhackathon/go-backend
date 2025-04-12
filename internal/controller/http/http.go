package http

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/service"
)

type Controller struct {
	cfg    *config.Config
	log    *zap.Logger
	srv    service.Interface
	server *echo.Echo
}

func New(cfg *config.Config, log *zap.Logger, service service.Interface) (*Controller, error) {
	ctrl := &Controller{
		cfg:    cfg,
		log:    log,
		srv:    service,
		server: echo.New(),
	}

	ctrl.configureMiddleware()
	ctrl.configureRoutes()

	return ctrl, nil
}

func (ctrl *Controller) configureMiddleware() {
}

func (ctrl *Controller) configureRoutes() {
	api := ctrl.server.Group("/api")

	api.GET("/test/:name", ctrl.GetTestByName)

	consumer := api.Group("/consumer")
	{
		consumer.POST("/registration", ctrl.CreateConsumer)
		consumer.PUT("/update", ctrl.UpdateConsumerPassword)
		consumer.DELETE("/delete", ctrl.DeleteConsumer)
		consumer.GET("", ctrl.GetConsumer)
		consumer.POST("/login", ctrl.Login)
		consumer.GET("/refresh-token", ctrl.RefreshToken)
	}

}
