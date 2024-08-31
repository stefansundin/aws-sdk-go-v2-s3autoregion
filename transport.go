package s3autoregion

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go"
)

// Custom error

type followXAmzBucketRegionError struct {
	NewRegion string
}

func (e followXAmzBucketRegionError) Error() string {
	return fmt.Sprintf("The bucket is located in %s.", e.NewRegion)
}

func (m followXAmzBucketRegionError) RetryableError() bool {
	return false
}

// RoundTripper

type followXAmzBucketRegion struct {
	tr http.RoundTripper
}

func (t followXAmzBucketRegion) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := t.tr.RoundTrip(r)
	if err != nil {
		return resp, err
	}

	switch resp.StatusCode {
	case 301, 302:
		if v := resp.Header.Get("x-amz-bucket-region"); len(v) > 0 {
			return nil, &followXAmzBucketRegionError{
				NewRegion: v,
			}
		}
	}

	return resp, err
}

// Helper function

func setRegionFn(region *string) func(o *s3.Options) {
	return func(o *s3.Options) {
		if region != nil {
			o.Region = *region
		}
	}
}

// Helper methods

func (c *Client) getBucketRegion(bucket *string) *string {
	if c.cache != nil {
		region, ok := c.cache.Get(*bucket)
		if ok {
			return &region
		}
	}
	return c.lastRegion
}

func (c *Client) followXAmzBucketRegion(bucket *string, region *string, err error) *string {
	if err == nil {
		c.lastRegion = region
		if c.cache != nil {
			c.cache.Add(*bucket, *region)
		}
		return nil
	}

	var e *followXAmzBucketRegionError
	if errors.As(err, &e) {
		c.lastRegion = &e.NewRegion
		if c.cache != nil {
			c.cache.Add(*bucket, e.NewRegion)
		}
		return &e.NewRegion
	} else {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "AuthorizationHeaderMalformed" {
			// This error can be returned when using S3 Accelerate since it doesn't use a regionalized endpoint
			// Example error message:
			// The authorization header is malformed; the region 'us-west-1' is wrong; expecting 'us-east-1'
			message := ae.ErrorMessage()
			if strings.HasPrefix(message, fmt.Sprintf("The authorization header is malformed; the region '%s' is wrong; expecting '", *region)) {
				correctRegion := message[strings.LastIndex(message, " ")+2 : len(message)-1]
				c.lastRegion = &correctRegion
				if c.cache != nil {
					c.cache.Add(*bucket, correctRegion)
				}
				return &correctRegion
			}
		}

		var notFoundError *types.NotFound
		if !errors.As(err, &notFoundError) {
			// There was an error but the bucket does indeed exist
			c.lastRegion = region
			if c.cache != nil {
				c.cache.Add(*bucket, *region)
			}
		}
	}

	return nil
}
