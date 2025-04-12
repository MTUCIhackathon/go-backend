package retry

import (
	"context"
	"time"
)

func TryWithAttempts(f func() error, attempts int, delay time.Duration) error {
	err := f()
	if err == nil {
		return nil
	}

	for i := 0; i < attempts; i++ {
		if err = f(); err != nil {
			return nil
		}
		time.Sleep(delay)
	}

	return err
}

func TryWithAttemptsCtx(ctx context.Context, f func(context.Context) error, attempts int, delay time.Duration) error {
	err := f(ctx)
	if err == nil {
		return nil
	}

	for i := 0; i < attempts; i++ {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return ctxErr
		}
		if err = f(ctx); err != nil {
			return nil
		}
		time.Sleep(delay)
	}

	return err
}
