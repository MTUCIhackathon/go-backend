package migrate

import (
	"context"

	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/migrations"
	"github.com/MTUCIhackathon/go-backend/pkg/migrator/tern"
)

func Migrate(cfg *config.Config, log *zap.Logger) error {
	if log == nil {
		log = zap.NewNop()
	}
	log.Named("migrate").Info("migrating")

	if !cfg.Postgres.RunMigrations {
		return nil
	}
	ctx := context.Background()
	m, err := tern.New(ctx, cfg, log, migrations.Migrations)
	if err != nil {
		return err
	}
	defer m.Close(ctx)

	v, err := m.GetVersion(ctx)
	if err != nil {
		return err
	}

	log.Info("current version", zap.Int32("version", v))

	return m.MigrateUp(ctx)
}
