package webcloud

import (
	"bytes"
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/MTUCIhackathon/go-backend/internal/config"
)

const imageKeyDir = "images"

type Client struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	config        *config.AWS
}

func New(ctx context.Context, cfg *config.Config) (*Client, error) {
	loadAWSConfig, err := awsConfig.LoadDefaultConfig(
		ctx,
		awsConfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AWS.AccessKey,
				cfg.AWS.SecretKey,
				"",
			),
		),
		awsConfig.WithBaseEndpoint(cfg.AWS.Host),
		awsConfig.WithRegion(cfg.AWS.Region),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(loadAWSConfig)
	presigned := s3.NewPresignClient(client)

	return &Client{
		client:        client,
		presignClient: presigned,
		config:        cfg.AWS,
	}, nil
}

func (c *Client) ListObjects(ctx context.Context) ([]string, error) {
	objects, err := c.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(c.config.Bucket),
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

func (c *Client) PutObject(ctx context.Context, key string, data []byte) error {
	_, err := c.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(c.config.Bucket),
		Key:    aws.String(imageKeyDir + "/" + key),
		Body:   bytes.NewReader(bytes.NewBuffer(data).Bytes()),
	})
	return err
}

func (c *Client) GenerateLink(ctx context.Context, key string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(c.config.Bucket),
		Key:    aws.String(imageKeyDir + "/" + key),
	}
	request, err := c.presignClient.PresignGetObject(ctx, input, func(opts *s3.PresignOptions) {
		opts.Expires = c.config.LinkLifeTime * time.Minute
	})
	if err != nil {
		return "", err
	}

	return request.URL, err
}
