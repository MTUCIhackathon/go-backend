package production

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/controller"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/MTUCIhackathon/go-backend/internal/service"
	"github.com/MTUCIhackathon/go-backend/internal/store/pgx"
)

/*func (s *Service) CreateResolved(req dto.CreateResolved) (*dto.Resolved, error) {
	return nil, service.NewError(
		controller.ErrInternal,
		errors.Wrap(nil, "some err"),
	)
}*/

func (s *Service) CreateConsumer(ctx context.Context, req dto.CreateConsumer) (*dto.Token, error) {
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
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successful encrypt password", zap.String("password", password))

	data = dto.Consumer{
		ID:        uuid.New(),
		Login:     req.Login,
		Password:  password,
		CreatedAt: time.Now(),
	}

	exists, err := s.repo.Consumers().GetLoginAvailable(ctx, data.Login)
	if err != nil {
		if errors.Is(err, pgx.ErrAlreadyExists) {
			s.log.Debug("failed to check login accessibility: login already exists", zap.Error(err))
			return nil, service.NewError(
				controller.ErrAlreadyExist,
				errors.Wrap(err, "failed to check login accessibility"),
			)
		}
		s.log.Debug("failed to check login accessibility", zap.Error(err))
		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to check login accessibility"),
		)
	}

	if !exists {
		s.log.Debug("failed to fetch available consumers", zap.String("consumer", data.Login))
		return nil, service.NewError(controller.ErrAlreadyExist, ErrAlreadyExists)
	}

	s.log.Debug("successfully fetch available consumers")

	err = s.repo.Consumers().Create(ctx, data)
	if err != nil {
		if errors.Is(err, pgx.ErrAlreadyExists) {
			s.log.Debug("failed to create consumer: consumer already exists", zap.Error(err))
			return nil, service.NewError(
				controller.ErrAlreadyExist,
				errors.Wrap(err, "failed to create consumer"),
			)
		}
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

func (s *Service) UpdateConsumerPassword(ctx context.Context, req dto.UpdatePassword) error {
	var err error
	data, err := s.provider.GetDataFromToken(req.Token)
	if err != nil {
		s.log.Debug("failed to fetch consumer ID", zap.Error(err))
		return service.NewError(controller.ErrBadRequest, err)
	}

	if !data.IsAccess {
		s.log.Debug("not is access", zap.Any("data", data))
	}

	id := data.ID
	s.log.Debug("successfully fetch consumer ID", zap.Any("id", id))

	err = s.valid.ValidatePassword(req.NewPassword)
	if err != nil {
		s.log.Debug("failed to validate password", zap.Error(err))
		return service.NewError(controller.ErrBadRequest, err)
	}

	s.log.Debug("successfully validate password", zap.Any("password", req.NewPassword))

	oldPassword, err := s.repo.Consumers().GetPasswordByID(ctx, id)
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

	err = s.repo.Consumers().UpdatePasswordByID(ctx, id, newPassword)
	if err != nil {
		s.log.Debug("failed to update password", zap.Error(err))
		return service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully update password", zap.Any("oldPassword", oldPassword))
	return nil
}

func (s *Service) DeleteConsumerByID(ctx context.Context, token string) error {
	data, err := s.provider.GetDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to fetch consumer ID", zap.Error(err))
		return service.NewError(controller.ErrBadRequest, err)
	}

	if !data.IsAccess {
		s.log.Debug("not is access", zap.Any("data", data))
		return service.NewError(controller.ErrInternal, err)
	}

	id := data.ID

	s.log.Debug("successfully fetch consumer id", zap.Any("id", id))

	err = s.repo.Consumers().DeleteByID(ctx, id)
	if err != nil {
		s.log.Debug("failed to delete consumer", zap.Error(err))
		return service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully delete consumer", zap.Any("consumer", id))
	return nil
}

func (s *Service) GetConsumerByID(ctx context.Context, token string) (*dto.Consumer, error) {
	data, err := s.provider.GetDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to fetch consumer ID", zap.Error(err))
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	if !data.IsAccess {
		s.log.Debug("not is access", zap.Any("data", data))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	id := data.ID
	s.log.Debug("successfully fetch consumer ID", zap.Any("id", id))

	consumer, err := s.repo.Consumers().GetByID(ctx, id)
	if err != nil {
		s.log.Debug("failed to fetch consumer", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully fetch consumer", zap.Any("consumer", consumer))

	return consumer, nil

}

func (s *Service) Login(ctx context.Context, req dto.Login) (*dto.Token, error) {
	data, err := s.repo.Consumers().GetByLogin(ctx, req.Login)
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

func (s *Service) RefreshToken(_ context.Context, token string) (*dto.Token, error) {
	data, err := s.provider.GetDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to fetch refresh token", zap.Error(err))
		return nil, service.NewError(controller.ErrBadRequest, err)
	}
	if data.IsAccess {
		s.log.Debug("not is refresh", zap.Any("data", data))
		return nil, service.NewError(controller.ErrInternal, err)
	}
	id := data.ID

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

func (s *Service) GetAllTests(_ context.Context, token string) ([]dto.Test, error) {
	data, err := s.provider.GetDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to fetch test data", zap.Error(err))
		return nil, service.NewError(controller.ErrBadRequest, err)
	}

	if !data.IsAccess {
		s.log.Debug("not is access", zap.Any("data", data))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	tests, err := s.inmemory.GetAll()
	if err != nil {
		s.log.Debug("failed to fetch test list", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully fetch test list", zap.Any("tests", tests))

	return tests, nil
}

func (s *Service) GetTestByID(_ context.Context, token string, testID uuid.UUID) (*dto.Test, error) {
	data, err := s.provider.GetDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to fetch test data", zap.Error(err))
		return nil, service.NewError(controller.ErrBadRequest, err)
	}
	if !data.IsAccess {
		s.log.Debug("not is access", zap.Any("data", data))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully fetch test ID", zap.Any("consumer", testID))

	test, err := s.inmemory.Get(testID)
	if err != nil {
		s.log.Debug("failed to fetch test", zap.Error(err))
		return nil, service.NewError(controller.ErrInternal, err)
	}

	s.log.Debug("successfully fetch test", zap.Any("test", test))

	return test, nil
}
