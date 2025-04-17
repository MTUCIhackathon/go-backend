package main

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/MTUCIhackathon/go-backend/internal/cache"
	"github.com/MTUCIhackathon/go-backend/internal/cache/inmemory"
	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/controller"
	"github.com/MTUCIhackathon/go-backend/internal/controller/http"
	"github.com/MTUCIhackathon/go-backend/internal/ml"
	"github.com/MTUCIhackathon/go-backend/internal/ml/client"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/assay"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/assay/study"
	encrytpor "github.com/MTUCIhackathon/go-backend/internal/pkg/encryptor"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/encryptor/hash"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/mark"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/mark/determinator"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/migrate"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/token/jwt"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/validator"
	"github.com/MTUCIhackathon/go-backend/internal/pkg/validator/valid"
	"github.com/MTUCIhackathon/go-backend/internal/service"
	"github.com/MTUCIhackathon/go-backend/internal/service/production"
	"github.com/MTUCIhackathon/go-backend/internal/store"
	storage "github.com/MTUCIhackathon/go-backend/internal/store/pgx"
	"github.com/MTUCIhackathon/go-backend/pkg/logger"
	"github.com/MTUCIhackathon/go-backend/pkg/migrator"
	"github.com/MTUCIhackathon/go-backend/pkg/migrator/tern"
	"github.com/MTUCIhackathon/go-backend/pkg/pgx"
	"github.com/MTUCIhackathon/go-backend/pkg/s3"
	"github.com/MTUCIhackathon/go-backend/pkg/s3/webcloud"
)

func CreateApp() fx.Option {
	return fx.Options(
		fx.WithLogger(fxLogger),
		fx.Provide(
			logger.New,
			config.New,
			createPgx,
			createS3,
			fx.Annotate(tern.New, fx.As(new(migrator.Interface))),
			fx.Annotate(cacheCreate, fx.As(new(cache.Cache))),
			fx.Annotate(jwt.NewProvider, fx.As(new(token.Provider))),
			fx.Annotate(determinator.NewMark, fx.As(new(mark.Marker))),
			fx.Annotate(valid.NewValidator, fx.As(new(validator.Interface))),
			fx.Annotate(storage.New, fx.As(new(store.Interface))),
			fx.Annotate(hash.New, fx.As(new(encrytpor.Interface))),
			fx.Annotate(production.New, fx.As(new(service.Interface))),
			fx.Annotate(http.New, fx.As(new(controller.Controller))),
			fx.Annotate(study.New, fx.As(new(assay.Interface))),
			fx.Annotate(client.New, fx.As(new(ml.Interface))),
		),
		fx.Invoke(
			controller.RunControllerFx,
			migrate.Migrate,
		),
	)
}

func createS3(log *zap.Logger, cfg *config.Config) (s3.Interface, error) {
	log = log.Named("aws")
	log.Debug("starting s3 with config", zap.Any("config", cfg.AWS))
	aws, err := webcloud.New(cfg)
	if err != nil {
		return nil, err
	}

	l, err := aws.GenerateLink(context.Background(), "")
	log.Debug("created link", zap.String("l", l))

	if aws == nil {
		return nil, errors.New("aws is nil")
	}

	return aws, nil
}

func createPgx(log *zap.Logger, cfg *config.Config) (*pgxpool.Pool, error) {
	pool, err := pgx.New(cfg, log, pgx.AddUUIDSupport, pgx.WithEnumTypeSupport())
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func fxLogger(log *zap.Logger) fxevent.Logger {
	return &fxevent.ZapLogger{Logger: log.Named("fx")}
}

func cacheCreate(cfg *config.Config, log *zap.Logger) (cache.Cache, error) {
	return inmemory.New(cfg, log, inmemory.WithLoader())
}

func main() {
	fx.New(CreateApp()).Run()
}
