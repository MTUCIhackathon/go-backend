package migrate

import (
	"context"

	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/migrations"
	"github.com/MTUCIhackathon/go-backend/pkg/migrator/tern"
)

func Migrate(cfg *config.Config, log *zap.Logger) error {
	if !cfg.Postgres.RunMigrations {
		return nil
	}
	ctx := context.Background()
	m, err := tern.New(ctx, cfg, log, migrations.Migrations)
	if err != nil {
		return err
	}
	defer m.Close(ctx)
	return m.MigrateUp(ctx)
}
