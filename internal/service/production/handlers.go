package production

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"time"

	"github.com/MTUCIhackathon/go-backend/internal/controller"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/MTUCIhackathon/go-backend/internal/service"
)

/*func (s *Service) CreateResolved(req dto.CreateResolved) (*dto.Resolved, error) {
	return nil, service.NewError(
		controller.ErrInternal,
		errors.Wrap(nil, "some err"),
	)
}*/

func (s *Service) CreateConsumer(e echo.Context, req dto.CreateConsumer) (*dto.Token, error) {
	var (
		data dto.Consumer
		err  error
		res  *dto.Token
	)
	err = s.valid.ValidatePassword(req.Password)
	if err != nil {
		return nil, service.NewError(controller.ErrBadRequest, err)
	}
	password, err := s.encrypt.EncryptPassword(req.Password)
	data = dto.Consumer{
		ID:        uuid.New(),
		Login:     req.Login,
		Password:  password,
		CreatedAt: time.Now(),
	}

	check, err := s.repo.Consumers().GetLoginAvailable(e.Request().Context(), data.Login)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}
	if !check {
		return nil, service.NewError(controller.ErrBadRequest, controller.ErrUnauthorized)
	}

	err = s.repo.Consumers().Create(e.Request().Context(), data)
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
