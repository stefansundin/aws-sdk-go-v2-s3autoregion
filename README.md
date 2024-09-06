A wrapper for the aws-sdk-go-v2 S3 client that automatically follows the `x-amz-bucket-region` header so that you don't have to worry about which region the bucket is in.

It also attempts to infer the correct region from S3 Express bucket names, e.g. `bucketname--usw2-az1--x-s3` is in `us-west-2`. It also works with S3 Accelerate.

Last updated for [v1.61.0](https://pkg.go.dev/mod/github.com/aws/aws-sdk-go-v2/service/s3@v1.61.0).

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
		config.WithRegion(initialRegion),
		func(o *config.LoadOptions) error {
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
	// Cache up to 50 bucket-to-region mappings
	client := s3autoregion.NewFromConfig(cfg, &s3autoregion.ExtendedOptions{
		CacheSize: 50,
	})
	fmt.Fprintf(os.Stderr, "Client region before operation: %s\n\n", client.Options().Region)

	result, err := client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	region, ok := client.GetBucketRegion(bucket)
	if ok {
		fmt.Fprintf(os.Stderr, "Bucket region: %s\n\n", region)
	} else {
		fmt.Fprintf(os.Stderr, "Bucket region not found!\n\n")
	}

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

SDK 2024/08/31 13:02:11 DEBUG Request
HEAD /policygen.html HTTP/1.1
Host: awspolicygen.s3.us-west-1.amazonaws.com
User-Agent: m/E aws-sdk-go-v2/1.30.4 os/macos lang/go#1.23.0 md/GOOS#darwin md/GOARCH#arm64 api/s3#1.60.1
Accept-Encoding: identity
Amz-Sdk-Invocation-Id: 35ec0564-f60c-4c73-b415-7a008c3a78b3
Amz-Sdk-Request: attempt=1; max=3
X-Amz-Content-Sha256: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

SDK 2024/08/31 13:02:12 DEBUG Request
HEAD /policygen.html HTTP/1.1
Host: awspolicygen.s3.us-east-1.amazonaws.com
User-Agent: m/E aws-sdk-go-v2/1.30.4 os/macos lang/go#1.23.0 md/GOOS#darwin md/GOARCH#arm64 api/s3#1.60.1
Accept-Encoding: identity
Amz-Sdk-Invocation-Id: c791c50d-d6e3-4702-937c-3793f3f37654
Amz-Sdk-Request: attempt=1; max=3
X-Amz-Content-Sha256: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855

SDK 2024/08/31 13:02:13 DEBUG Response
HTTP/1.1 200 OK
Content-Length: 12352
Accept-Ranges: bytes
Content-Type: text/html
Date: Sat, 31 Aug 2024 20:02:14 GMT
Etag: "c6f16024402e4523b140dc5907227539"
Last-Modified: Fri, 30 Aug 2024 10:26:25 GMT
Server: AmazonS3
X-Amz-Id-2: mYpZaZ2GIDhqV8ehA/kw5aFaOHq7zVCKozCuDbxGKIP3mGFHgPtMjAknV/4KMV+jlto31l8pE1k=
X-Amz-Request-Id: 6S5TJ1HYNHZDPFK5
X-Amz-Server-Side-Encryption: AES256
X-Amz-Version-Id: rGk6PFnAkEXtvKpm0PfuuR1k.XCZg2sg

Bucket region: us-east-1

HeadObject ContentLength: 12352
```
