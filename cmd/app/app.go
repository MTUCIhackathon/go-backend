package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/cache"
	"github.com/MTUCIhackathon/go-backend/internal/cache/inmemory"
	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/internal/controller"
	"github.com/MTUCIhackathon/go-backend/internal/controller/http"
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
			pgx.New,
			fx.Annotate(webcloud.New, fx.As(new(s3.Interface))),
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
		),
		fx.Invoke(
			controller.RunControllerFx,
			migrate.Migrate,
		),
	)
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
