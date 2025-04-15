package production

import (
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/cache"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/assay"
	encrytpor "github.com/MTUCIhackathon/go-backend/internal/pkg/encryptor"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/mark"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/validator"

	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token"
	"github.com/MTUCIhackathon/go-backend/internal/store"
)

type Service struct {
	log          *zap.Logger
	repo         store.Interface
	provider     token.Provider
	config       *config.Config
	encrypt      encrytpor.Interface
	valid        validator.Interface
	inmemory     cache.Cache
	determinator mark.Marker
	study        assay.Interface
}

func New(
	log *zap.Logger,
	repo store.Interface,
	provider token.Provider,
	config *config.Config,
	encrypt encrytpor.Interface,
	valid validator.Interface,
	inmemory cache.Cache,
	determinator mark.Marker,
	study assay.Interface,
) (*Service, error) {
	if log == nil {
		log = zap.L().Named("service.production")
		log.Warn(
			"provided nil logger, initializing with global logger",
		)
	}

	if repo == nil || provider == nil || config == nil || encrypt == nil || inmemory == nil || valid == nil || study == nil {
		log.Warn(
			"provided nil service dependency",
		)

		// TODO
		return nil, ErrNilReference
	}

	s := &Service{
		log:          log,
		repo:         repo,
		provider:     provider,
		config:       config,
		encrypt:      encrypt,
		valid:        valid,
		inmemory:     inmemory,
		determinator: determinator,
		study:        study,
	}

	return s, nil
}
