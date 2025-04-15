package s3

import (
	"context"
)

type Interface interface {
	ListObjects(ctx context.Context, bucket string) ([]string, error)
	PutObject(ctx context.Context, bucket string, key string, data []byte) error
	GenerateLink(ctx context.Context, bucket string, key string) (string, error)
}
