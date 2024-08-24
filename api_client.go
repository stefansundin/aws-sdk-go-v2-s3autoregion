package s3autoregion

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsHttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	client *s3.Client
}

func New(options s3.Options, optFns ...func(*s3.Options)) *Client {
	options.HTTPClient = &http.Client{
		Transport: followXAmzBucketRegion{
			tr: awsHttp.NewBuildableClient().GetTransport(),
		},
	}
	return &Client{
		client: s3.New(options, optFns...),
	}
}

func NewFromConfig(cfg aws.Config, optFns ...func(*s3.Options)) *Client {
	cfg.HTTPClient = &http.Client{
		Transport: followXAmzBucketRegion{
			tr: awsHttp.NewBuildableClient().GetTransport(),
		},
	}
	return &Client{
		client: s3.NewFromConfig(cfg, optFns...),
	}
}

func (c *Client) Options() s3.Options {
	return c.client.Options()
}
