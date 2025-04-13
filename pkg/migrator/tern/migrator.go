package tern

import (
	"context"
	"io/fs"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/tern/v2/migrate"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/MTUCIhackathon/go-backend/internal/config"
	"github.com/MTUCIhackathon/go-backend/pkg/migrator"
)

var _ migrator.Interface = (*Migrator)(nil)

type Migrator struct {
	cfg        *config.Config
	log        *zap.Logger
	migrations fs.FS
	conn       *pgx.Conn
	migrator   tern
}

type tern interface {
	Migrate(ctx context.Context) error
	MigrateTo(ctx context.Context, version int32) error
	GetCurrentVersion(ctx context.Context) (int32, error)
	LoadMigrations(migrations fs.FS) error
}

func New(ctx context.Context, cfg *config.Config, log *zap.Logger, migrations fs.FS) (*Migrator, error) {
	if cfg == nil {
		return nil, errNilConfig
	}

	if log == nil {
		log = zap.NewNop()
		log.Named("migrator")
		log.Warn("logger is nil, configuring with global logger")
	}

	m := &Migrator{
		cfg:        cfg,
		log:        log,
		migrations: migrations,
		conn:       nil,
	}

	err := multierr.Combine(
		m.initConn(ctx),
		m.initMigrator(ctx),
		m.loadMigrations(),
	)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Migrator) MigrateUp(ctx context.Context) error {
	return m.migrator.Migrate(ctx)
}

func (m *Migrator) MigrateDown(ctx context.Context) error {
	return m.migrator.MigrateTo(ctx, 0)
}

func (m *Migrator) MigrateTo(ctx context.Context, version int32) error {
	return m.migrator.MigrateTo(ctx, version)
}

func (m *Migrator) GetVersion(ctx context.Context) (int32, error) {
	return m.migrator.GetCurrentVersion(ctx)
}

func (m *Migrator) loadMigrations() (err error) {
	err = m.migrator.LoadMigrations(m.migrations)
	if err != nil {
		m.log.Error("migrator failed to load migrations", zap.Error(err))
		return errLoadingMigrations
	}
	return nil
}

func (m *Migrator) initConn(ctx context.Context) (err error) {
	m.conn, err = pgx.Connect(ctx, m.cfg.Postgres.GetURI())
	if err != nil {
		m.log.Error("failed to create connection to database", zap.Error(err))
		return errFailedToConnect
	}
	m.log.Debug("connected to database")
	return nil
}

func (m *Migrator) initMigrator(ctx context.Context) (err error) {
	m.migrator, err = migrate.NewMigrator(ctx, m.conn, m.cfg.Postgres.VersionTableName)
	if err != nil {
		return errCreationMigrator
	}
	return nil
}
