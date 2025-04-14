package migrator

import (
	"context"
)

type Interface interface {
	MigrateUp(ctx context.Context) error
	MigrateDown(ctx context.Context) error
	MigrateTo(ctx context.Context, version int32) error
	GetVersion(ctx context.Context) (int32, error)
}
