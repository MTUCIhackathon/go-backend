package production

import (
	"github.com/MTUCIhackathon/go-backend/internal/pkg/encrytpor"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token"
	"github.com/MTUCIhackathon/go-backend/internal/store"
)

type Service struct {
	log      *zap.Logger
	repo     store.Interface
	provider token.Provider
	config   *config.Config
	encrypt  encrytpor.Interface
}

func New(
	log *zap.Logger,
	repo store.Interface,
	provider token.Provider,
	config *config.Config,
	encrypt encrytpor.Interface,
) (*Service, error) {
	if log == nil {
		log = zap.L().Named("service.production")
		log.Warn(
			"provided nil logger, initializing with global logger",
		)
	}

	if repo == nil || provider == nil || config == nil || encrypt == nil {
		log.Warn(
			"provided nil service dependency",
		)

		// TODO
		return nil, ErrNilReference
	}

	s := &Service{
		log:      log,
		repo:     repo,
		provider: provider,
		config:   config,
		encrypt:  encrypt,
	}

	return s, nil
}
