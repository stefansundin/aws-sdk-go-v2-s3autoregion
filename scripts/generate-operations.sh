#!/bin/bash -e

if [ $# -eq 0 ]; then
  echo "Please provide a path to the S3 module and pipe the output to a file. Example:"
  echo
  echo "$0 ~/go/pkg/mod/github.com/aws/aws-sdk-go-v2/service/s3@v1.60.1 > new-operations.go"
  echo
  exit 1
fi

cat <<EOS
package s3autoregion

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

EOS


find "$1" -type f -name 'api_op_*' -maxdepth 1 | sort | while read path
do
  op=$(basename "$path" .go | cut -d '_' -f 3)
  cat <<EOS
func (c *Client) ${op}(ctx context.Context, params *s3.${op}Input, optFns ...func(*s3.Options)) (*s3.${op}Output, error) {
	result, err := c.client.${op}(ctx, params, optFns...)
	if c.followXAmzBucketRegion(err) {
		result, err = c.client.${op}(ctx, params, optFns...)
	}
	return result, err
}

EOS
done
