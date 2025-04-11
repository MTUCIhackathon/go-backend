package http

import (
	"errors"

	"github.com/MTUCIhackathon/go-backend/internal/controller"
)

func handleErr(err error) error {
	if errors.Is(err, controller.ErrorBadRequest) {

	}
}
