package production

import (
	"github.com/MTUCIhackathon/go-backend/internal/pkg/encrytpor"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"time"

	"github.com/MTUCIhackathon/go-backend/internal/controller"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/MTUCIhackathon/go-backend/internal/service"
)

func (s *Service) CreateResolved(req dto.CreateResolved) (*dto.Resolved, error) {
	return nil, service.NewError(
		controller.ErrInternal,
		errors.Wrap(nil, "some err"),
	)
}

// TODO: add user data validation
func (s *Service) CreateConsumer(e echo.Context, req dto.CreateConsumer) (*dto.Token, error) {
	var (
		data *dto.Consumer
		err  error
		res  *dto.Token
	)

	password, err := encrytpor.Interface().EncryptPassword()

	data = &dto.Consumer{
		ID:        uuid.New(),
		Login:     req.Login,
		Email:     req.Email,
		Password:  password,
		CreatedAt: time.Now(),
	}

	err = s.repo.Consumers().CreateConsumer(e.Request().Context(), data)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}

	err = s.repo.Consumers().CreateConsumer(e.Request().Context(), data)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}

	access, err := s.provider.CreateAccessTokenForUser(data.ID)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}

	refresh, err := s.provider.CreateRefreshTokenForUser(data.ID)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}

	res = &dto.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}
	return res, nil
}
