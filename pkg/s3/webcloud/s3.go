package webcloud

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

type Client struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	config        *config.AWS
}

func New(ctx context.Context, cfg *config.AWS) (*Client, error) {
	loadAWSConfig, err := awsConfig.LoadDefaultConfig(
		ctx,
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKey,
				cfg.SecretKey,
				cfg.Region,
			),
		),
		awsConfig.WithBaseEndpoint(cfg.Host),
		awsConfig.WithRegion(cfg.Region),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(loadAWSConfig)
	presigned := s3.NewPresignClient(client)

	return &Client{
		client:        client,
		presignClient: presigned,
		config:        cfg,
	}, nil
}

func (c *Client) ListObjects(ctx context.Context, bucket string) ([]string, error) {
	objects, err := c.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(objects.Contents))
	for _, object := range objects.Contents {
		result = append(result, aws.ToString(object.Key))
	}

	return result, nil
}

func (c *Client) PutObject(ctx context.Context, bucket string, key string, data []byte) error {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(bytes.NewBuffer(data).Bytes()),
	})
	return err
}

func (c *Client) GenerateLink(ctx context.Context, bucket string, key string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	request, err := c.presignClient.PresignGetObject(ctx, input, func(opts *s3.PresignOptions) {
		opts.Expires = c.config.LinkLifeTime
	})
	if err != nil {
		return "", err
	}

	return request.URL, err
}
