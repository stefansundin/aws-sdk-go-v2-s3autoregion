A wrapper for the aws-sdk-go-v2 S3 client that automatically follows the `x-amz-bucket-region` header so that you don't have to worry about which region the bucket is in.

Last updated for [v1.60.1](https://pkg.go.dev/mod/github.com/aws/aws-sdk-go-v2/service/s3@v1.60.1).

## Example usage

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3autoregion "github.com/stefansundin/aws-sdk-go-v2-s3autoregion"
)

func main() {
	initialRegion := "us-west-1"
	bucket := "awspolicygen"
	key := "policygen.html"
	debug := true

	// Initialize the AWS SDK
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		func(o *config.LoadOptions) error {
			o.Region = initialRegion
			o.Credentials = aws.AnonymousCredentials{}
			if debug {
				var lm aws.ClientLogMode = aws.LogRequest | aws.LogResponse
				o.ClientLogMode = &lm
			}
			return nil
		},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Create a client using the s3autoregion wrapper client
	client := s3autoregion.NewFromConfig(cfg)

	fmt.Fprintf(os.Stderr, "Client region before operation: %s\n\n", client.Options().Region)
	result, err := client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	fmt.Fprintf(os.Stderr, "Client region after operation: %s\n\n", client.Options().Region)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "HeadObject ContentLength: %d\n", aws.ToInt64(result.ContentLength))
}
```

Output should be:

```
Client region before operation: us-west-1

SDK 2024/08/23 23:15:41 DEBUG Request
HEAD /policygen.html HTTP/1.1
Host: awspolicygen.s3.us-west-1.amazonaws.com
User-Agent: m/E aws-sdk-go-v2/1.30.4 os/macos lang/go#1.23.0 md/GOOS#darwin md/GOARCH#arm64 api/s3#1.60.1
Accept-Encoding: identity
Amz-Sdk-Invocation-Id: 07e0feb2-7180-45dd-8d4c-cfa139fe78ea
Amz-Sdk-Request: attempt=1; max=3
X-Amz-Content-Sha256: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

SDK 2024/08/23 23:15:43 DEBUG Request
HEAD /policygen.html HTTP/1.1
Host: awspolicygen.s3.us-east-1.amazonaws.com
User-Agent: m/E aws-sdk-go-v2/1.30.4 os/macos lang/go#1.23.0 md/GOOS#darwin md/GOARCH#arm64 api/s3#1.60.1
Accept-Encoding: identity
Amz-Sdk-Invocation-Id: 5b409f04-0abb-4282-9f78-e24ee087e371
Amz-Sdk-Request: attempt=1; max=3
X-Amz-Content-Sha256: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

SDK 2024/08/23 23:15:44 DEBUG Response
HTTP/1.1 200 OK
Content-Length: 12352
Accept-Ranges: bytes
Content-Type: text/html
Date: Sat, 24 Aug 2024 06:15:45 GMT
Etag: "c6f16024402e4523b140dc5907227539"
Last-Modified: Wed, 21 Aug 2024 23:27:19 GMT
Server: AmazonS3
X-Amz-Id-2: wnnN+3ObRVZ9TknD/13Vg7WxIpzBIkRWWj/5++LiKp266Q7gY7r2vT7SnJNKQsQrTRnKcHWxGaA=
X-Amz-Request-Id: 7S74BW0GCS0R0SCN
X-Amz-Server-Side-Encryption: AES256
X-Amz-Version-Id: iOBfn9K1JspCgEVzL4JSAD9bWsjee1hG

Client region after operation: us-east-1

HeadObject ContentLength: 12352
```
