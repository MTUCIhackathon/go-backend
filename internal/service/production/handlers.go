package production

import (
	"github.com/pkg/errors"

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
