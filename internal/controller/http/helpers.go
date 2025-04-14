package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/MTUCIhackathon/go-backend/internal/controller"
)

func handleErr(err error) error {
	switch {
	case err == nil:
		return echo.NewHTTPError(http.StatusOK, "OK")
	case errors.Is(err, controller.ErrUnknown):
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	case errors.Is(err, controller.ErrBadRequest):
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	case errors.Is(err, controller.ErrNotFound):
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	case errors.Is(err, controller.ErrAlreadyExist):
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	case errors.Is(err, controller.ErrUnauthorized):
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	case errors.Is(err, controller.ErrForbidden):
		return echo.NewHTTPError(http.StatusForbidden, err.Error())
	case errors.Is(err, controller.ErrInternal):
		fallthrough
	default:
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
}
