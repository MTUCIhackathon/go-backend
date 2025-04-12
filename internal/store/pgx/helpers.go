package pgx

import (
	"github.com/pkg/errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/zap"
)

func (s *Store) pgErr(err error) error {
	if err == nil {
		return nil
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case pgerrcode.UniqueViolation:
			return ErrAlreadyExists
		case pgerrcode.ForeignKeyViolation:
			return ErrNotFound

		default:
			s.log.Error(
				"got unexpected error",
				zap.Error(err),
				zap.String("pg_err_code", pgErr.Code),
				zap.String("pg_err_message", pgErr.Message),
				zap.String("pg_err_where", pgErr.Where),
				zap.String("pg_err_schema_name", pgErr.SchemaName),
				zap.String("pg_err_table_name", pgErr.TableName),
			)

			return errors.Wrapf(
				ErrUnknown,
				"unknown error: message=%q code=%q", pgErr.Message, pgErr.Code,
			)
		}
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}

	s.log.Error("got unexpected error", zap.Error(err))

	return ErrUnknown
}
