package production

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
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
		s.log.Debug("failed to validate password", zap.Error(err))
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successful validate password")

	password, err := s.encrypt.EncryptPassword(req.Password)
	if err != nil {
		s.log.Debug("failed to encrypt password", zap.Error(err))
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successful encrypt password", zap.String("password", password))

	data = dto.Consumer{
		ID:        uuid.New(),
		Login:     req.Login,
		Password:  password,
		CreatedAt: time.Now(),
	}

	check, err := s.repo.Consumers().GetLoginAvailable(e.Request().Context(), data.Login)
	if err != nil {
		s.log.Debug("failed to fetch available consumers", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	if !check {
		s.log.Debug("failed to fetch available consumers", zap.String("consumer", data.Login))
		return nil, service.NewError(controller.ErrBadRequest, controller.ErrUnauthorized)
	}

	s.log.Debug("successfully fetch available consumers")

	err = s.repo.Consumers().Create(e.Request().Context(), data)
	if err != nil {
		s.log.Debug("failed to create consumer", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully create consumer", zap.Any("consumer", data))

	access, err := s.provider.CreateAccessTokenForUser(data.ID)
	if err != nil {
		s.log.Debug("failed to create access token", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully create access token", zap.Any("access", access))
	refresh, err := s.provider.CreateRefreshTokenForUser(data.ID)
	if err != nil {
		s.log.Debug("failed to create refresh token", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}
	s.log.Debug("successfully create refresh token", zap.Any("refresh", refresh))
	res = &dto.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}
	s.log.Debug("successfully create tokens", zap.Any("tokens", res))
	return res, nil
}

func (s *Service) UpdateConsumerPassword(e echo.Context, req dto.UpdatePassword) error {
	var (
		err error
	)
	id, err := s.getConsumerIDFromRequest(e.Request())
	if err != nil {
		s.log.Debug("failed to fetch consumer ID", zap.Error(err))
		return service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successfully fetch consumer ID", zap.Any("id", id))

	err = s.valid.ValidatePassword(req.NewPassword)
	if err != nil {
		s.log.Debug("failed to validate password", zap.Error(err))
		return service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successfully validate password", zap.Any("password", req.NewPassword))

	oldPassword, err := s.repo.Consumers().GetPasswordByID(e.Request().Context(), id)
	if err != nil {
		s.log.Debug("failed to fetch old password", zap.Error(err))
		return service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully fetch old password", zap.Any("oldPassword", oldPassword))

	err = s.encrypt.CompareHashAndPassword(oldPassword, req.OldPassword)
	if err != nil {
		s.log.Debug("failed to compare old password", zap.Error(err))
		return service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successfully compare old password", zap.Any("oldPassword", oldPassword))

	newPassword, err := s.encrypt.EncryptPassword(req.NewPassword)
	if err != nil {
		s.log.Debug("failed to encrypt password", zap.Error(err))
		return service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successfully encrypt password", zap.Any("newPassword", newPassword))

	err = s.repo.Consumers().UpdatePasswordByID(e.Request().Context(), id, newPassword)
	if err != nil {
		s.log.Debug("failed to update password", zap.Error(err))
		return service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully update password", zap.Any("oldPassword", oldPassword))
	return nil
}

func (s *Service) DeleteConsumerByID(e echo.Context) error {
	id, err := s.getConsumerIDFromRequest(e.Request())
	if err != nil {
		s.log.Debug("failed to fetch consumer ID", zap.Error(err))
		return service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successfully fetch consumer ID", zap.Any("id", id))

	err = s.repo.Consumers().DeleteByID(e.Request().Context(), id)
	if err != nil {
		s.log.Debug("failed to delete consumer", zap.Error(err))
		return service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully delete consumer", zap.Any("consumer", id))
	return nil
}

func (s *Service) GetConsumerByID(c echo.Context) (*dto.Consumer, error) {
	id, err := s.getConsumerIDFromRequest(c.Request())
	if err != nil {
		s.log.Debug("failed to fetch consumer ID", zap.Error(err))
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successfully fetch consumer ID", zap.Any("id", id))

	data, err := s.repo.Consumers().GetByID(c.Request().Context(), id)
	if err != nil {
		s.log.Debug("failed to fetch consumer", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully fetch consumer", zap.Any("consumer", data))

	return data, nil

}

func (s *Service) Login(c echo.Context, req dto.Login) (*dto.Token, error) {
	data, err := s.repo.Consumers().GetByLogin(c.Request().Context(), req.Login)
	if err != nil {
		s.log.Debug("failed to fetch consumer", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully fetch consumer", zap.Any("consumer", data))

	err = s.encrypt.CompareHashAndPassword(data.Password, req.Password)
	if err != nil {
		s.log.Debug("failed to compare password", zap.Error(err))
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successfully compare password", zap.Any("password", data.Password))

	access, err := s.provider.CreateAccessTokenForUser(data.ID)
	if err != nil {
		s.log.Debug("failed to create access token", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully create access token", zap.Any("access", access))

	refresh, err := s.provider.CreateRefreshTokenForUser(data.ID)
	if err != nil {
		s.log.Debug("failed to create refresh token", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully create refresh token", zap.Any("refresh", refresh))

	res := &dto.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	s.log.Debug("successfully create tokens", zap.Any("tokens", res))

	return res, nil
}

func (s *Service) RefreshToken(c echo.Context) (*dto.Token, error) {
	id, err := s.validRefreshToken(c.Request())
	if err != nil {
		s.log.Debug("failed to validate refresh token", zap.Error(err))
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successfully validate refresh token", zap.Any("id", id))

	access, err := s.provider.CreateAccessTokenForUser(id)
	if err != nil {
		s.log.Debug("failed to create access token", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully create access token", zap.Any("access", access))

	refresh, err := s.provider.CreateRefreshTokenForUser(id)
	if err != nil {
		s.log.Debug("failed to create refresh token", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully create refresh token", zap.Any("refresh", refresh))

	res := &dto.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}

	s.log.Debug("successfully create tokens", zap.Any("tokens", res))
	return res, nil
}
