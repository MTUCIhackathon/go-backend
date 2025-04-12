package pgx

import (
	"context"

	"github.com/jackc/pgx/v5"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"
)

type AfterConnect func(ctx context.Context, conn *pgx.Conn) error

func AddUUIDSupport(_ context.Context, conn *pgx.Conn) error {
	pgxUUID.Register(conn.TypeMap())
	return nil
}
func WithEnumTypeSupport(t string) AfterConnect {
	return func(ctx context.Context, conn *pgx.Conn) error {
		dt, err := conn.LoadType(ctx, t)
		if err != nil {
			return err
		}
		conn.TypeMap().RegisterType(dt)
		return nil
	}
}
