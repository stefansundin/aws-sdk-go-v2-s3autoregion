package s3autoregion

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"unsafe"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Custom error

type followXAmzBucketRegionError struct {
	NewRegion string
}

func (e followXAmzBucketRegionError) Error() string {
	return fmt.Sprintf("This bucket is located in %s.", e.NewRegion)
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

// Helper method

func (c *Client) followXAmzBucketRegion(err error) bool {
	var e *followXAmzBucketRegionError
	if errors.As(err, &e) {
		field := reflect.ValueOf(c.client).Elem().FieldByName("options")
		s := GetUnexportedField(field).(s3.Options)
		// fmt.Printf("Updating S3 Client Region from %s to %s.\n\n", s.Region, e.NewRegion)
		s.Region = e.NewRegion
		SetUnexportedField(field, s)
		return true
	}
	return false
}

// https://stackoverflow.com/a/60598827/2730380
func GetUnexportedField(field reflect.Value) interface{} {
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

func SetUnexportedField(field reflect.Value, value interface{}) {
	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Set(reflect.ValueOf(value))
}
