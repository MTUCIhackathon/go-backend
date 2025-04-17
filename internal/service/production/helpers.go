package production

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/controller"
	"github.com/MTUCIhackathon/go-backend/internal/model/dto"
	"github.com/MTUCIhackathon/go-backend/internal/service"
)

func (s *Service) GetConsumerDataFromToken(token string) (*dto.ConsumerDataInToken, error) {
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

	switch {
	case data == nil:
		s.log.Error("failed to get consumer data from token: collected nil data")

		return nil, service.NewError(
			controller.ErrForbidden,
			errors.New("failed to get consumer data from token: collected nil data"),
		)
	case !data.IsAccess:
		s.log.Error("failed to get consumer data from token: should be access")

		return nil, service.NewError(
			controller.ErrForbidden,
			errors.New("failed to get consumer data from token: should be access"),
		)
	default:
		return data, nil
	}
}

func (s *Service) UploadImage(ctx context.Context, profession string, imageKey string) (*string, error) {
	rawImage, err := s.ml.HandlerGenerateImage(profession)
	if err != nil {
		s.log.Error("failed to generate image", zap.Error(err))
		return nil, err
	}

	s.log.Debug("raw image", zap.Any("rawImage", rawImage[:50]), zap.String("imageKey", imageKey))

	err = s.s3.PutObject(ctx, imageKey, rawImage)
	if err != nil {
		s.log.Error("failed to upload image", zap.Error(err))
		return nil, err
	}

	s.log.Debug("image uploaded", zap.String("imageKey", imageKey))

	imageLink, err := s.s3.GenerateLink(ctx, imageKey)
	if err != nil {
		s.log.Error("failed to generate link", zap.Error(err))
		return nil, err
	}

	s.log.Debug("image link generated", zap.Any("link", imageLink))

	return &imageLink, nil
}
