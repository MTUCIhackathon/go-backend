package tern

import (
	"errors"
)

var (
	errNilConfig         = errors.New("config is nil")
	errFailedToConnect   = errors.New("failed to create connection to database")
	errCreationMigrator  = errors.New("failed to create tern migrator")
	errLoadingMigrations = errors.New("failed to load migrations")
)
