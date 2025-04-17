package http

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
		log:    log.Named("http-controller"),
		srv:    service,
		server: echo.New(),
	}

	ctrl.configureMiddleware()
	ctrl.configureRoutes()

	return ctrl, nil
}

func (ctrl *Controller) configureMiddleware() {
	ctrl.server.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.Logger(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			// TODO
			AllowOrigins:     []string{ctrl.cfg.Controller.ServerAddress, ctrl.cfg.Controller.TestAddress, "5.228.114.32", "http://192.168.0.93:3000", "http://localhost:3000"},
			AllowMethods:     []string{http.MethodOptions, http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
			AllowCredentials: true,
		}),
	)
}

func (ctrl *Controller) configureRoutes() {
	api := ctrl.server.Group("/api")
	api.GET("/ping", ctrl.Ping)

	consumer := api.Group("/consumers")
	{
		consumer.POST("/registration", ctrl.CreateConsumer)
		consumer.PUT("/update_password", ctrl.UpdateConsumerPassword)
		consumer.DELETE("/delete", ctrl.DeleteConsumer)
		consumer.GET("/get_me", ctrl.GetMe)
		consumer.POST("/login", ctrl.Login)
		consumer.GET("/refresh-token", ctrl.RefreshToken)
	}

	test := api.Group("/tests")
	{
		test.GET("/all", ctrl.GetAllTest)
		test.GET("/:test_id", ctrl.GetTestByID)
	}

	result := api.Group("/results")
	{
		result.GET("/all", nil)
		result.GET("/:result_id", ctrl.GetResultByResolvedID)
		result.GET("/my", ctrl.GetMyResult)
		result.POST("/create", ctrl.CreateResult)
		result.POST("/send_on_email", ctrl.SendResultOnEmail)
	}

	resolved := api.Group("/resolved")
	{
		resolved.POST("/create", ctrl.CreateResolved)
		resolved.GET("/:resolved_id", ctrl.GetResolvedByID)
	}

	ai := api.Group("/ai")
	{
		ai.POST("/send_questions", ctrl.CreateDataForMl)
		ai.GET("/summary", ctrl.GetSummary)
		ai.POST("/get_result", ctrl.GetDataFromMl)
	}
}

func (ctrl *Controller) Start(ctx context.Context) error {
	ch := make(chan error, 1)

	go func() {
		ch <- ctrl.server.Start(ctrl.cfg.Controller.Bind())
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(time.Millisecond * 300):
		return nil
	}
}

func (ctrl *Controller) Stop(ctx context.Context) error {
	return ctrl.server.Shutdown(ctx)
}
