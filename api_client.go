package s3autoregion

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsHttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	lru "github.com/hashicorp/golang-lru/v2"
)

type Client struct {
	client     *s3.Client
	cache      *lru.Cache[string, string]
	lastRegion *string
}

type ExtendedOptions struct {
	// The size of the LRU cache that remembers what region buckets reside in
	CacheSize int
}

func New(options s3.Options, extendedOptions *ExtendedOptions, optFns ...func(*s3.Options)) *Client {
	var cache *lru.Cache[string, string]
	if extendedOptions != nil {
		if extendedOptions.CacheSize > 0 {
			var err error
			cache, err = lru.New[string, string](extendedOptions.CacheSize)
			if err != nil {
				panic(fmt.Errorf("error creating LRU cache: %w", err))
			}
		}
	}
	options.HTTPClient = &http.Client{
		Transport: followXAmzBucketRegion{
			tr: awsHttp.NewBuildableClient().GetTransport(),
		},
	}
	client := s3.New(options, optFns...)
	effectiveOptions := client.Options()
	return &Client{
		client:     client,
		cache:      cache,
		lastRegion: &effectiveOptions.Region,
	}
}

func NewFromConfig(cfg aws.Config, extendedOptions *ExtendedOptions, optFns ...func(*s3.Options)) *Client {
	var cache *lru.Cache[string, string]
	if extendedOptions != nil {
		if extendedOptions.CacheSize > 0 {
			var err error
			cache, err = lru.New[string, string](extendedOptions.CacheSize)
			if err != nil {
				panic(fmt.Errorf("error creating LRU cache: %w", err))
			}
		}
	}
	cfg.HTTPClient = &http.Client{
		Transport: followXAmzBucketRegion{
			tr: awsHttp.NewBuildableClient().GetTransport(),
		},
	}
	client := s3.NewFromConfig(cfg, optFns...)
	effectiveOptions := client.Options()
	return &Client{
		client:     client,
		cache:      cache,
		lastRegion: &effectiveOptions.Region,
	}
}

func (c *Client) Options() s3.Options {
	return c.client.Options()
}

func (c *Client) GetBucketRegion(bucket string) (region string, ok bool) {
	return c.cache.Get(bucket)
}
