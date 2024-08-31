package s3autoregion

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
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

	// The wrapper client needs to override the HTTPClient in order to function
	// If you need to modify the transport in any way, use this function to do so
	TransportOptionsFn func(*http.Transport)
}

func New(options s3.Options, extendedOptions *ExtendedOptions, optFns ...func(*s3.Options)) *Client {
	httpClientBuilder := awshttp.NewBuildableClient()

	var cache *lru.Cache[string, string]
	if extendedOptions != nil {
		if extendedOptions.CacheSize > 0 {
			var err error
			cache, err = lru.New[string, string](extendedOptions.CacheSize)
			if err != nil {
				panic(fmt.Errorf("error creating LRU cache: %w", err))
			}
		}
		if extendedOptions.TransportOptionsFn != nil {
			httpClientBuilder = httpClientBuilder.WithTransportOptions(extendedOptions.TransportOptionsFn)
		}
	}

	options.HTTPClient = &http.Client{
		Transport: followXAmzBucketRegion{
			tr: httpClientBuilder.GetTransport(),
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
	httpClientBuilder := awshttp.NewBuildableClient()

	var cache *lru.Cache[string, string]
	if extendedOptions != nil {
		if extendedOptions.CacheSize > 0 {
			var err error
			cache, err = lru.New[string, string](extendedOptions.CacheSize)
			if err != nil {
				panic(fmt.Errorf("error creating LRU cache: %w", err))
			}
		}
		if extendedOptions.TransportOptionsFn != nil {
			httpClientBuilder = httpClientBuilder.WithTransportOptions(extendedOptions.TransportOptionsFn)
		}
	}

	cfg.HTTPClient = &http.Client{
		Transport: followXAmzBucketRegion{
			tr: httpClientBuilder.GetTransport(),
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
	if c.cache == nil {
		return "", false
	}
	return c.cache.Get(bucket)
}
