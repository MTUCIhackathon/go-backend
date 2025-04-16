package production

import (
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

func (s *Service) unmarshalPointer(str *string) (string, error) {
	if str == nil {
		return "", nil
	}
	return *str, nil
}
