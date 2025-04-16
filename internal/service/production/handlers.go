package production

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/controller"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/style/kind"
	"github.com/MTUCIhackathon/go-backend/internal/service"
	"github.com/MTUCIhackathon/go-backend/internal/store/pgx"
)

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

	if !data.IsAccess {
		return service.NewError(
			controller.ErrUnauthorized,
			errors.Wrap(err, "token is not an access token"),
		)
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

	if !data.IsAccess {
		return service.NewError(
			controller.ErrUnauthorized,
			errors.Wrap(err, "token is not an access token"),
		)
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

	if !data.IsAccess {
		return nil, service.NewError(
			controller.ErrUnauthorized,
			errors.Wrap(err, "token is not an access token"),
		)
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

func (s *Service) CreateResolved(ctx context.Context, token string, req dto.ResolvedRequest) (*dto.Resolved, error) {
	var (
		err  error
		resp dto.Resolved
	)

	userData, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to fetch consumer data from token", zap.Error(err))
		return nil, service.NewError(
			controller.ErrUnauthorized,
			errors.Wrap(err, "failed to fetch consumer data from token"))
	}

	questions := make([]dto.Question, len(req.Questions))
	for i := 0; i < len(req.Questions); i++ {
		question := req.Questions[i]

		mark, err := s.determinator.MarkResult(question.QuestionAnswer)
		if err != nil {
			s.log.Debug("failed to determinate result", zap.Error(err))
			return nil, service.NewError(
				controller.ErrInternal,
				errors.Wrap(err, "failed to fetch consumer data from token"),
			)
		}

		questions[i] = dto.Question{
			ResolvedID:     question.ResolvedID,
			QuestionOrder:  question.QuestionOrder,
			Issue:          question.Issue,
			QuestionAnswer: question.QuestionAnswer,
			ImageLocation:  question.ImageLocation,
			Mark:           mark,
		}
	}

	resp = dto.Resolved{
		ID:           req.ID,
		UserID:       userData.ID,
		ResolvedType: req.ResolvedType,
		IsActive:     req.IsActive,
		CreatedAt:    req.CreatedAt,
		PassedAt:     req.PassedAt,
		Questions:    questions,
	}

	err = s.repo.Resolved().CreateResolved(ctx, resp)
	if err != nil {
		s.log.Debug("failed to create resolved", zap.Error(err))
		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to determinate test type"))
	}

	return &resp, nil
}

func (s *Service) CreateResultByFirstTest(ctx context.Context, token string, req dto.Resolved) (*dto.Result, error) {
	userData, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to fetch consumer data from token", zap.Error(err))
		return nil, service.NewError(
			controller.ErrUnauthorized,
			errors.Wrap(err, "failed to fetch consumer data from token"))
	}

	if req.ResolvedType != kind.FirstOrder {
		s.log.Debug("test type incorrect", zap.Any("result", req))
	}

	topMarks := make([]dto.Mark, len(req.Questions))

	for i := 0; i < len(req.Questions); i++ {
		mark := dto.Mark{
			Order: req.Questions[i].QuestionOrder,
			Mark:  req.Questions[i].Mark,
		}
		topMarks[i] = mark
	}

	areas, err := s.study.First().GetAreas(topMarks)
	if err != nil {
		s.log.Debug("failed to get marks", zap.Error(err))
		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get areas"))
	}

	professions, err := s.ml.HandlerSendResultsForFirstTest(areas)

	resp := dto.Result{
		ID:            uuid.New(),
		UserID:        userData.ID,
		ResolvedID:    req.ID,
		ImageLocation: nil,
		Profession:    professions,
		CreatedAt:     req.CreatedAt,
	}

	err = s.repo.Results().CreateResult(ctx, resp)
	if err != nil {
		s.log.Debug("failed to insert result", zap.Error(err))
		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to determinate test type"))
	}

	return &resp, nil
}

func (s *Service) CreateResultBySecondTest(ctx context.Context, token string, req dto.Resolved) (*dto.Result, error) {
	userData, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to fetch consumer data from token", zap.Error(err))
		return nil, service.NewError(
			controller.ErrUnauthorized,
			errors.Wrap(err, "failed to fetch consumer data from token"))
	}

	if req.ResolvedType != kind.SecondOrder {
		s.log.Debug("test type incorrect", zap.Any("result", req))
	}

	topMarks := make([]dto.Mark, len(req.Questions))

	for i := 0; i < len(req.Questions); i++ {
		mark := dto.Mark{
			Order: req.Questions[i].QuestionOrder,
			Mark:  req.Questions[i].Mark,
		}
		topMarks[i] = mark
	}

	personality, err := s.study.Second().GetPersonality(topMarks)
	if err != nil {
		s.log.Debug("failed to get marks", zap.Error(err))
		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get personality"))
	}

	mlReq, err := s.ml.HandlerSendResultsForSecondTest(personality)
	if err != nil {
		s.log.Debug("failed to insert result", zap.Error(err))
		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get personality"))
	}

	resp := dto.Result{
		ID:            uuid.New(),
		UserID:        userData.ID,
		ResolvedID:    req.ID,
		ImageLocation: nil,
		Profession:    mlReq.Professions,
		CreatedAt:     req.CreatedAt,
	}

	err = s.repo.Results().CreateResult(ctx, resp)

	if err != nil {
		s.log.Debug("failed to put result in db", zap.Error(err))
		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get personality"))
	}

	return &resp, nil
}

func (s *Service) GetResultByTestID(ctx context.Context, token string, testID uuid.UUID) (*dto.Result, error) {
	userData, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Debug("failed to fetch consumer data from token", zap.Error(err))
		return nil, service.NewError(
			controller.ErrUnauthorized,
			errors.Wrap(err, "failed to fetch consumer data from token"))
	}

	resp, err := s.repo.Results().GetResultByResolvedIDAndUserID(ctx, userData.ID, testID)
	if err != nil {
		s.log.Debug("failed to fetch result", zap.Error(err))
		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get result"))
	}

	return resp, err
}

func (s *Service) SaveResult(ctx context.Context, token string, req dto.ResultCreation) error {
	data, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Error(
			"failed to fetch consumer data from token",
			zap.Error(err),
		)

		return err
	}

	err = s.repo.Results().CreateResult(ctx, dto.Result{
		ID:            uuid.New(),
		UserID:        data.ID,
		ResolvedID:    req.ResolveID,
		ImageLocation: req.ImageLocation,
		Profession:    req.Professions,
		CreatedAt:     time.Now(),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrAlreadyExists) {
			s.log.Error(
				"failed to create result: already exists",
				zap.Error(err),
			)

			return service.NewError(
				controller.ErrAlreadyExist,
				errors.Wrap(err, "failed to create result"),
			)
		}
		s.log.Error(
			"failed to create result",
			zap.Error(err),
		)

		return service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to create result"),
		)
	}

	return nil
}

func (s *Service) GetResultByResolvedID(ctx context.Context, token string, resultID uuid.UUID) (*dto.Result, error) {
	data, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Error(
			"failed to fetch consumer data from token",
			zap.Error(err),
		)

		return nil, err
	}

	result, err := s.repo.Results().GetResultByResolvedIDAndUserID(ctx, data.ID, resultID)
	if err != nil {
		if errors.Is(err, pgx.ErrNotFound) {
			s.log.Error(
				"failed to get result by user id and result id: not found",
				zap.Error(err),
			)

			return nil, service.NewError(
				controller.ErrNotFound,
				errors.Wrap(err, "failed to get result by user id and result id"),
			)
		}
		s.log.Error(
			"failed to get result by user id and result id",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get result by user id and result id"),
		)
	}

	return result, nil
}

func (s *Service) GetResultsByUserID(ctx context.Context, token string) ([]dto.Result, error) {
	data, err := s.GetConsumerDataFromToken(token)
	if err != nil {
		s.log.Error(
			"failed to fetch consumer data from token",
			zap.Error(err),
		)

		return nil, err
	}

	results, err := s.repo.Results().GetResultByUserID(ctx, data.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNotFound) {
			s.log.Error(
				"failed to get results by user id",
				zap.Error(err),
			)

			return nil, service.NewError(
				controller.ErrNotFound,
				errors.Wrap(err, "failed to get results by user id"),
			)
		}
		s.log.Error(
			"failed to get results by user id",
			zap.Error(err),
		)

		return nil, service.NewError(
			controller.ErrInternal,
			errors.Wrap(err, "failed to get results by user id"),
		)
	}

	return results, nil
}
