package main

import (
	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token/jwt"
	"github.com/MTUCIhackathon/go-backend/pkg/logger"
	"github.com/labstack/echo/v4"
)

func main() {
	log, err := logger.New("dev")

	if err != nil {
		panic(err)
	}
	log.Info("initialized logger")
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	_, err = jwt.NewProvider(cfg, log)
	if err != nil {
		panic(err)
	}

	//_, err = smtp.New(cfg, log)
	/*if err != nil {
		panic(err)
	}*/
	e := echo.New()
	if err := e.Start("localhost:8080"); err != nil {
		panic(err)
	}
}
