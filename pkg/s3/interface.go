package s3

import (
	"context"
)

type Interface interface {
	ListObjects(ctx context.Context) ([]string, error)
	PutObject(ctx context.Context, key string, data []byte) error
	GenerateLink(ctx context.Context, key string) (string, error)
}
