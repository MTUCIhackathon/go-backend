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
		consumer dto.Consumer
		err      error
	)

	err = s.valid.ValidatePassword(req.Password)
	if err != nil {
		s.log.Error(
			"failed to validate password",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrBadRequest,
			errors.Wrap(err, "password does not match the standard"))
	}

	password, err := s.encrypt.EncryptPassword(req.Password)
	if err != nil {
		s.log.Error(
			"failed to encrypt password",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to encrypt password"),
		)
	}

	consumer = dto.Consumer{
		ID:        uuid.New(),
		Login:     req.Login,
		Password:  password,
		CreatedAt: time.Now(),
	}

	exists, err := s.repo.Consumers().GetLoginAvailable(ctx, consumer.Login)
	if err != nil {
		s.log.Error(
			"failed to check login accessibility",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to check login accessibility"),
		)
	}

	if !exists {
		s.log.Error(
			"login already exists",
			zap.String("login", consumer.Login),
		)

		return nil, service.NewError(
			controller.ErrAlreadyExist,
			ErrAlreadyExists,
		)
	}

	err = s.repo.Consumers().Create(ctx, consumer)
	if err != nil {
		if errors.Is(err, pgx.ErrAlreadyExists) {
			s.log.Error(
				"failed to create consumer: consumer already exists",
				zap.Error(err),
			)

			return nil, service.NewError(
				controller.ErrAlreadyExist,
				errors.Wrap(err, "failed to create consumer"),
			)
		}
		s.log.Error(
			"failed to create consumer",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to create consumer"),
		)
	}

	access, refresh, err := s.provider.CreateAccessAndRefreshTokenForUser(consumer.ID)
	if err != nil {
		s.log.Error(
			"failed to create a couples of tokens",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to create a couples of tokens"),
		)
	}

	s.log.Debug("successfully created consumer", zap.String("consumer_id", consumer.ID.String()))

	return &dto.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *Service) UpdateConsumerPassword(ctx context.Context, req dto.UpdatePassword) error {
	data, err := s.GetConsumerDataFromToken(req.Token)
	if err != nil {
		s.log.Error(
			"failed to get consumer data from token",
			zap.Error(err),
		)

		return err
	}

	err = s.valid.ValidatePassword(req.NewPassword)
	if err != nil {
		s.log.Error(
			"failed to validate password",
			zap.Error(err),
		)

		return service.NewError(
			controller.ErrBadRequest,
			errors.Wrap(err, "password does not match the standard"),
		)
	}

	oldPassword, err := s.repo.Consumers().GetPasswordByID(ctx, data.ID)
	if err != nil {
		if errors.Is(pgx.ErrNotFound, err) {
			s.log.Error(
				"failed to fetch old password: not found",
				zap.Error(err),
			)

			return service.NewError(
				controller.ErrNotFound,
				errors.Wrap(err, "failed to fetch old password"),
			)
		}
		s.log.Debug(
			"failed to fetch old password",
			zap.Error(err),
		)

		return service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to fetch old password"),
		)
	}

	err = s.encrypt.CompareHashAndPassword(oldPassword, req.OldPassword)
	if err != nil {
		s.log.Error(
			"failed to compare old passwords",
			zap.Error(err),
		)

		return service.NewError(
			controller.ErrForbidden,
			errors.Wrap(err, "failed to compare old passwords"),
		)
	}

	newPassword, err := s.encrypt.EncryptPassword(req.NewPassword)
	if err != nil {
		s.log.Error(
			"failed to encrypt password",
			zap.Error(err),
		)

		return service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to encrypt password"),
		)
	}

	err = s.repo.Consumers().UpdatePasswordByID(ctx, data.ID, newPassword)
	if err != nil {
		if errors.Is(pgx.ErrNotFound, err) {
			s.log.Error(
				"failed to update password: not found",
				zap.Error(err),
			)

			return service.NewError(
				controller.ErrNotFound,
				errors.Wrap(err, "failed to update password"),
			)
		}
		s.log.Error(
			"failed to update password",
			zap.Error(err),
		)

		return service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to update password"),
		)
	}

	s.log.Debug("successfully updated password for consumer", zap.String("consumer_id", data.ID.String()))

	return nil
}

func (s *Service) DeleteConsumerByID(ctx context.Context, token string) error {
	data, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Error(
			"failed to get consumer data from token",
			zap.Error(err),
		)

		return err
	}

	err = s.repo.Consumers().DeleteByID(ctx, data.ID)
	if err != nil {
		if errors.Is(pgx.ErrNotFound, err) {
			s.log.Error(
				"failed to delete consumer by id: not found",
				zap.Error(err),
			)

			return service.NewError(
				controller.ErrNotFound,
				errors.Wrap(err, "failed to delete consumer by id"),
			)
		}
		s.log.Error(
			"failed to delete consumer by id",
			zap.Error(err),
		)

		return service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to delete consumer by id"),
		)
	}

	s.log.Debug("successfully deleted consumer", zap.String("consumer_id", data.ID.String()))

	return nil
}

func (s *Service) GetConsumerByID(ctx context.Context, token string) (*dto.Consumer, error) {
	data, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Error(
			"failed to get consumer data from token",
			zap.Error(err),
		)

		return nil, err
	}

	consumer, err := s.repo.Consumers().GetByID(ctx, data.ID)
	if err != nil {
		if errors.Is(pgx.ErrNotFound, err) {
			s.log.Error(
				"failed to get consumer by id: not found",
				zap.Error(err),
			)

			return nil, service.NewError(
				controller.ErrNotFound,
				errors.Wrap(err, "failed to get consumer"),
			)
		}
		s.log.Error(
			"failed to get consumer by id",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get consumer"),
		)
	}

	s.log.Debug("successfully got consumer by id", zap.String("consumer_id", consumer.ID.String()))

	return consumer, nil
}

func (s *Service) Login(ctx context.Context, req dto.Login) (*dto.Token, error) {
	consumer, err := s.repo.Consumers().GetByLogin(ctx, req.Login)
	if err != nil {
		if errors.Is(pgx.ErrNotFound, err) {
			s.log.Error(
				"failed to get consumer by login: not found",
				zap.Error(err),
			)

			return nil, service.NewError(
				controller.ErrNotFound,
				errors.Wrap(err, "failed to get consumer by login"),
			)
		}
		s.log.Error(
			"failed to get consumer by login",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get consumer by login"),
		)
	}

	err = s.encrypt.CompareHashAndPassword(consumer.Password, req.Password)
	if err != nil {
		s.log.Error(
			"failed to compare password",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrForbidden,
			errors.Wrap(err, "failed to compare old passwords"),
		)
	}

	access, refresh, err := s.provider.CreateAccessAndRefreshTokenForUser(consumer.ID)
	if err != nil {
		s.log.Error(
			"failed to create a couples of tokens",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to create a couples of tokens"),
		)
	}

	s.log.Debug("successfully logged in", zap.String("consumer_id", consumer.ID.String()))

	return &dto.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *Service) RefreshToken(_ context.Context, token string) (*dto.Token, error) {
	data, err := s.provider.GetDataFromToken(token)
	if err != nil {
		s.log.Error(
			"failed to get consumer data from token",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrUnauthorized,
			errors.Wrap(err, "failed to get consumer data from token"),
		)
	}

	if data.IsAccess {
		s.log.Error(
			"not acceptable token",
			zap.Bool("is_access", data.IsAccess),
		)

		return nil, service.NewError(
			controller.ErrForbidden,
			ErrTokenNotAcceptable,
		)
	}

	access, refresh, err := s.provider.CreateAccessAndRefreshTokenForUser(data.ID)
	if err != nil {
		s.log.Error(
			"failed to create a couples of tokens",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to create a couples of tokens"),
		)
	}

	s.log.Debug("successfully refreshed tokens for consumer", zap.Any("consumer_id", data.ID.String()))

	return &dto.Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (s *Service) GetAllTests(_ context.Context, token string) ([]dto.Test, error) {
	data, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Error(
			"failed to get consumer data from token",
			zap.Error(err),
		)

		return nil, err
	}

	tests, err := s.inmemory.GetAll()
	if err != nil {
		s.log.Error(
			"failed to get test list",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get test list"),
		)
	}

	s.log.Debug(
		"successfully got tests for consumer",
		zap.String("consumer_id", data.ID.String()),
	)

	return tests, nil
}

func (s *Service) GetTestByID(_ context.Context, token string, testID uuid.UUID) (*dto.Test, error) {
	data, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Error(
			"failed to fetch consumer data from token",
			zap.Error(err),
		)

		return nil, err
	}

	test, err := s.inmemory.Get(testID)
	if err != nil {
		s.log.Error(
			"failed to get test by id",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get test by id"),
		)
	}

	s.log.Debug(
		"successfully got test for consumer",
		zap.String("test_id", test.ID.String()),
		zap.String("consumer_id", data.ID.String()),
	)

	return test, nil
}
