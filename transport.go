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

func setRegionFn(region string) func(o *s3.Options) {
	return func(o *s3.Options) {
		o.Region = region
	}
}

// Helper methods

func isNumeric(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func expandRegionShortName(s string) (string, bool) {
	// This attempts to convert a region "short" name (e.g. usw2) to a region (e.g. us-west-2), without relying on a hard-coded list
	// We just copy the first two characters (e.g. us), add a dash, decode the cardinal-ish directions (e.g. w to west), add another dash and the remaining numbers
	// This just aims to support the "normal" aws partition. There is a check for GovCloud, so it probably works there too (but I can't test it). I have no idea about the other partitions!
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html
	region := s[:2] + "-"
	for i, c := range s[2:] {
		if isNumeric(c) {
			region += "-" + s[2+i:]
			break
		} else if c == 'g' {
			region += "gov-"
		} else if c == 'n' {
			region += "north"
		} else if c == 's' {
			region += "south"
		} else if c == 'e' {
			region += "east"
		} else if c == 'w' {
			region += "west"
		} else if c == 'c' {
			region += "central"
		} else {
			return "", false
		}
	}
	return region, true
}

func (c *Client) getBucketRegion(bucket *string) string {
	// Check for directory bucket suffix
	if prefix, ok := strings.CutSuffix(*bucket, "--x-s3"); ok {
		i := strings.LastIndex(prefix, "--") + 2
		j := strings.LastIndex(prefix, "-")
		regionShortName := prefix[i:j]
		if region, ok := expandRegionShortName(regionShortName); ok {
			return region
		}
	}

	if c.cache != nil {
		region, ok := c.cache.Get(*bucket)
		if ok {
			return region
		}
	}
	return c.lastRegion
}

func (c *Client) setBucketRegion(bucket *string, region string) {
	if region == "" {
		region = "us-east-1"
	} else if region == "EU" {
		region = "eu-west-1"
	}

	c.lastRegion = region

	// No need to cache directory buckets
	if strings.HasSuffix(*bucket, "--x-s3") {
		return
	}

	if c.cache != nil {
		c.cache.Add(*bucket, region)
	}
}

func (c *Client) followXAmzBucketRegion(bucket *string, region string, err error) (string, bool) {
	if err == nil {
		c.setBucketRegion(bucket, region)
		return "", false
	}

	var e *followXAmzBucketRegionError
	if errors.As(err, &e) {
		c.setBucketRegion(bucket, e.NewRegion)
		return e.NewRegion, true
	} else {
		var ae smithy.APIError
		if errors.As(err, &ae) && ae.ErrorCode() == "AuthorizationHeaderMalformed" {
			// This error can be returned when using S3 Accelerate since it doesn't use a regionalized endpoint
			// Example error message:
			// The authorization header is malformed; the region 'us-west-1' is wrong; expecting 'us-east-1'
			message := ae.ErrorMessage()
			if strings.HasPrefix(message, fmt.Sprintf("The authorization header is malformed; the region '%s' is wrong; expecting '", region)) {
				correctRegion := message[strings.LastIndex(message, " ")+2 : len(message)-1]
				c.setBucketRegion(bucket, correctRegion)
				return correctRegion, true
			}
		}

		var notFoundError *types.NotFound
		if !errors.As(err, &notFoundError) {
			// There was an error but the bucket does indeed exist
			c.setBucketRegion(bucket, region)
		}
	}

	return "", false
}
