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

func (s *Service) UpdateConsumerPassword(e echo.Context, req dto.UpdatePassword) error {
	var (
		err error
	)
	id, err := s.getConsumerIDFromRequest(e.Request())
	if err != nil {
		return service.NewError(controller.ErrBadRequest, err)
	}

	err = s.valid.ValidatePassword(req.NewPassword)
	if err != nil {
		return service.NewError(controller.ErrBadRequest, err)
	}

	oldPassword, err := s.repo.Consumers().GetPasswordByID(e.Request().Context(), id)
	if err != nil {
		return service.NewError(controller.ErrInternal, err)
	}

	err = s.encrypt.CompareHashAndPassword(oldPassword, req.OldPassword)
	if err != nil {
		return service.NewError(controller.ErrBadRequest, err)
	}

	newPassword, err := s.encrypt.EncryptPassword(req.NewPassword)
	if err != nil {
		return service.NewError(controller.ErrBadRequest, err)
	}

	err = s.repo.Consumers().UpdatePasswordByID(e.Request().Context(), id, newPassword)
	if err != nil {
		return service.NewError(controller.ErrInternal, err)
	}

	return nil
}

func (s *Service) DeleteConsumerByID(e echo.Context) error {
	id, err := s.getConsumerIDFromRequest(e.Request())
	if err != nil {
		return service.NewError(controller.ErrBadRequest, err)
	}

	err = s.repo.Consumers().DeleteByID(e.Request().Context(), id)
	if err != nil {
		return service.NewError(controller.ErrInternal, err)
	}
	return nil
}

func (s *Service) GetConsumerByID(c echo.Context) (*dto.Consumer, error) {
	id, err := s.getConsumerIDFromRequest(c.Request())
	if err != nil {
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	data, err := s.repo.Consumers().GetByID(c.Request().Context(), id)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}

	return data, nil

}

func (s *Service) Login(c echo.Context, req dto.Login) (*dto.Token, error) {
	data, err := s.repo.Consumers().GetByLogin(c.Request().Context(), req.Login)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}

	err = s.encrypt.CompareHashAndPassword(data.Password, req.Password)
	if err != nil {
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	access, err := s.provider.CreateAccessTokenForUser(data.ID)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}

	refresh, err := s.provider.CreateRefreshTokenForUser(data.ID)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}

	res := &dto.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}
	return res, nil
}

func (s *Service) RefreshToken(c echo.Context) (*dto.Token, error) {
	id, err := s.validRefreshToken(c.Request())
	if err != nil {
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	access, err := s.provider.CreateAccessTokenForUser(id)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}

	refresh, err := s.provider.CreateRefreshTokenForUser(id)
	if err != nil {
		return nil, service.NewError(controller.ErrInternal, err)
	}
	res := &dto.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}
	return res, nil
}
