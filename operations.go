package s3autoregion

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Client) AbortMultipartUpload(ctx context.Context, params *s3.AbortMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.AbortMultipartUpload(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.AbortMultipartUpload(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) CompleteMultipartUpload(ctx context.Context, params *s3.CompleteMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.CompleteMultipartUpload(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.CompleteMultipartUpload(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.CopyObject(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.CopyObject(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.CreateBucket(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.CreateBucket(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) CreateMultipartUpload(ctx context.Context, params *s3.CreateMultipartUploadInput, optFns ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.CreateMultipartUpload(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.CreateMultipartUpload(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) CreateSession(ctx context.Context, params *s3.CreateSessionInput, optFns ...func(*s3.Options)) (*s3.CreateSessionOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.CreateSession(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.CreateSession(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucket(ctx context.Context, params *s3.DeleteBucketInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucket(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucket(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketAnalyticsConfiguration(ctx context.Context, params *s3.DeleteBucketAnalyticsConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketAnalyticsConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketAnalyticsConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketAnalyticsConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketCors(ctx context.Context, params *s3.DeleteBucketCorsInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketCorsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketCors(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketCors(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketEncryption(ctx context.Context, params *s3.DeleteBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketEncryptionOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketEncryption(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketEncryption(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketIntelligentTieringConfiguration(ctx context.Context, params *s3.DeleteBucketIntelligentTieringConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketIntelligentTieringConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketIntelligentTieringConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketIntelligentTieringConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketInventoryConfiguration(ctx context.Context, params *s3.DeleteBucketInventoryConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketInventoryConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketInventoryConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketInventoryConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketLifecycle(ctx context.Context, params *s3.DeleteBucketLifecycleInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketLifecycleOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketLifecycle(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketLifecycle(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketMetricsConfiguration(ctx context.Context, params *s3.DeleteBucketMetricsConfigurationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketMetricsConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketMetricsConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketMetricsConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketOwnershipControls(ctx context.Context, params *s3.DeleteBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketOwnershipControlsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketOwnershipControls(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketOwnershipControls(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketPolicy(ctx context.Context, params *s3.DeleteBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketPolicyOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketPolicy(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketPolicy(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketReplication(ctx context.Context, params *s3.DeleteBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketReplicationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketReplication(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketReplication(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketTagging(ctx context.Context, params *s3.DeleteBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketTaggingOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketTagging(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketTagging(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteBucketWebsite(ctx context.Context, params *s3.DeleteBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketWebsiteOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteBucketWebsite(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteBucketWebsite(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteObject(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteObject(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteObjectTagging(ctx context.Context, params *s3.DeleteObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectTaggingOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteObjectTagging(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteObjectTagging(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeleteObjects(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeleteObjects(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) DeletePublicAccessBlock(ctx context.Context, params *s3.DeletePublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.DeletePublicAccessBlockOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.DeletePublicAccessBlock(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.DeletePublicAccessBlock(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketAccelerateConfiguration(ctx context.Context, params *s3.GetBucketAccelerateConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketAccelerateConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketAccelerateConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketAccelerateConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketAcl(ctx context.Context, params *s3.GetBucketAclInput, optFns ...func(*s3.Options)) (*s3.GetBucketAclOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketAcl(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketAcl(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketAnalyticsConfiguration(ctx context.Context, params *s3.GetBucketAnalyticsConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketAnalyticsConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketAnalyticsConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketAnalyticsConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketCors(ctx context.Context, params *s3.GetBucketCorsInput, optFns ...func(*s3.Options)) (*s3.GetBucketCorsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketCors(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketCors(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketEncryption(ctx context.Context, params *s3.GetBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.GetBucketEncryptionOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketEncryption(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketEncryption(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketIntelligentTieringConfiguration(ctx context.Context, params *s3.GetBucketIntelligentTieringConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketIntelligentTieringConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketIntelligentTieringConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketIntelligentTieringConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketInventoryConfiguration(ctx context.Context, params *s3.GetBucketInventoryConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketInventoryConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketInventoryConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketInventoryConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketLifecycleConfiguration(ctx context.Context, params *s3.GetBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLifecycleConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketLifecycleConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketLifecycleConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketLocation(ctx context.Context, params *s3.GetBucketLocationInput, optFns ...func(*s3.Options)) (*s3.GetBucketLocationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketLocation(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketLocation(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketLogging(ctx context.Context, params *s3.GetBucketLoggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketLoggingOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketLogging(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketLogging(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketMetricsConfiguration(ctx context.Context, params *s3.GetBucketMetricsConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketMetricsConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketMetricsConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketMetricsConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketNotificationConfiguration(ctx context.Context, params *s3.GetBucketNotificationConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetBucketNotificationConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketNotificationConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketNotificationConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketOwnershipControls(ctx context.Context, params *s3.GetBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.GetBucketOwnershipControlsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketOwnershipControls(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketOwnershipControls(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketPolicy(ctx context.Context, params *s3.GetBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketPolicy(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketPolicy(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketPolicyStatus(ctx context.Context, params *s3.GetBucketPolicyStatusInput, optFns ...func(*s3.Options)) (*s3.GetBucketPolicyStatusOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketPolicyStatus(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketPolicyStatus(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketReplication(ctx context.Context, params *s3.GetBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.GetBucketReplicationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketReplication(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketReplication(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketRequestPayment(ctx context.Context, params *s3.GetBucketRequestPaymentInput, optFns ...func(*s3.Options)) (*s3.GetBucketRequestPaymentOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketRequestPayment(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketRequestPayment(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketTagging(ctx context.Context, params *s3.GetBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.GetBucketTaggingOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketTagging(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketTagging(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketVersioning(ctx context.Context, params *s3.GetBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.GetBucketVersioningOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketVersioning(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketVersioning(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetBucketWebsite(ctx context.Context, params *s3.GetBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.GetBucketWebsiteOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetBucketWebsite(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetBucketWebsite(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetObject(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetObject(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetObjectAcl(ctx context.Context, params *s3.GetObjectAclInput, optFns ...func(*s3.Options)) (*s3.GetObjectAclOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetObjectAcl(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetObjectAcl(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetObjectAttributes(ctx context.Context, params *s3.GetObjectAttributesInput, optFns ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetObjectAttributes(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetObjectAttributes(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetObjectLegalHold(ctx context.Context, params *s3.GetObjectLegalHoldInput, optFns ...func(*s3.Options)) (*s3.GetObjectLegalHoldOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetObjectLegalHold(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetObjectLegalHold(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetObjectLockConfiguration(ctx context.Context, params *s3.GetObjectLockConfigurationInput, optFns ...func(*s3.Options)) (*s3.GetObjectLockConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetObjectLockConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetObjectLockConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetObjectRetention(ctx context.Context, params *s3.GetObjectRetentionInput, optFns ...func(*s3.Options)) (*s3.GetObjectRetentionOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetObjectRetention(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetObjectRetention(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetObjectTagging(ctx context.Context, params *s3.GetObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetObjectTagging(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetObjectTagging(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetObjectTorrent(ctx context.Context, params *s3.GetObjectTorrentInput, optFns ...func(*s3.Options)) (*s3.GetObjectTorrentOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetObjectTorrent(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetObjectTorrent(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) GetPublicAccessBlock(ctx context.Context, params *s3.GetPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.GetPublicAccessBlockOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.GetPublicAccessBlock(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.GetPublicAccessBlock(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) HeadBucket(ctx context.Context, params *s3.HeadBucketInput, optFns ...func(*s3.Options)) (*s3.HeadBucketOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.HeadBucket(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.HeadBucket(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.HeadObject(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.HeadObject(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) ListBucketAnalyticsConfigurations(ctx context.Context, params *s3.ListBucketAnalyticsConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketAnalyticsConfigurationsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.ListBucketAnalyticsConfigurations(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.ListBucketAnalyticsConfigurations(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) ListBucketIntelligentTieringConfigurations(ctx context.Context, params *s3.ListBucketIntelligentTieringConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketIntelligentTieringConfigurationsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.ListBucketIntelligentTieringConfigurations(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.ListBucketIntelligentTieringConfigurations(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) ListBucketInventoryConfigurations(ctx context.Context, params *s3.ListBucketInventoryConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketInventoryConfigurationsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.ListBucketInventoryConfigurations(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.ListBucketInventoryConfigurations(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) ListBucketMetricsConfigurations(ctx context.Context, params *s3.ListBucketMetricsConfigurationsInput, optFns ...func(*s3.Options)) (*s3.ListBucketMetricsConfigurationsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.ListBucketMetricsConfigurations(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.ListBucketMetricsConfigurations(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) ListBuckets(ctx context.Context, params *s3.ListBucketsInput, optFns ...func(*s3.Options)) (*s3.ListBucketsOutput, error) {
	return c.client.ListBuckets(ctx, params, optFns...)
}

func (c *Client) ListDirectoryBuckets(ctx context.Context, params *s3.ListDirectoryBucketsInput, optFns ...func(*s3.Options)) (*s3.ListDirectoryBucketsOutput, error) {
	return c.client.ListDirectoryBuckets(ctx, params, optFns...)
}

func (c *Client) ListMultipartUploads(ctx context.Context, params *s3.ListMultipartUploadsInput, optFns ...func(*s3.Options)) (*s3.ListMultipartUploadsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.ListMultipartUploads(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.ListMultipartUploads(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) ListObjectVersions(ctx context.Context, params *s3.ListObjectVersionsInput, optFns ...func(*s3.Options)) (*s3.ListObjectVersionsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.ListObjectVersions(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.ListObjectVersions(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) ListObjects(ctx context.Context, params *s3.ListObjectsInput, optFns ...func(*s3.Options)) (*s3.ListObjectsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.ListObjects(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.ListObjects(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.ListObjectsV2(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.ListObjectsV2(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) ListParts(ctx context.Context, params *s3.ListPartsInput, optFns ...func(*s3.Options)) (*s3.ListPartsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.ListParts(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.ListParts(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketAccelerateConfiguration(ctx context.Context, params *s3.PutBucketAccelerateConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketAccelerateConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketAccelerateConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketAccelerateConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketAcl(ctx context.Context, params *s3.PutBucketAclInput, optFns ...func(*s3.Options)) (*s3.PutBucketAclOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketAcl(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketAcl(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketAnalyticsConfiguration(ctx context.Context, params *s3.PutBucketAnalyticsConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketAnalyticsConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketAnalyticsConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketAnalyticsConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketCors(ctx context.Context, params *s3.PutBucketCorsInput, optFns ...func(*s3.Options)) (*s3.PutBucketCorsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketCors(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketCors(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketEncryption(ctx context.Context, params *s3.PutBucketEncryptionInput, optFns ...func(*s3.Options)) (*s3.PutBucketEncryptionOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketEncryption(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketEncryption(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketIntelligentTieringConfiguration(ctx context.Context, params *s3.PutBucketIntelligentTieringConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketIntelligentTieringConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketIntelligentTieringConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketInventoryConfiguration(ctx context.Context, params *s3.PutBucketInventoryConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketInventoryConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketInventoryConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketInventoryConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketLifecycleConfiguration(ctx context.Context, params *s3.PutBucketLifecycleConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketLifecycleConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketLifecycleConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketLifecycleConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketLogging(ctx context.Context, params *s3.PutBucketLoggingInput, optFns ...func(*s3.Options)) (*s3.PutBucketLoggingOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketLogging(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketLogging(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketMetricsConfiguration(ctx context.Context, params *s3.PutBucketMetricsConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketMetricsConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketMetricsConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketMetricsConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketNotificationConfiguration(ctx context.Context, params *s3.PutBucketNotificationConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutBucketNotificationConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketNotificationConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketNotificationConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketOwnershipControls(ctx context.Context, params *s3.PutBucketOwnershipControlsInput, optFns ...func(*s3.Options)) (*s3.PutBucketOwnershipControlsOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketOwnershipControls(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketOwnershipControls(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketPolicy(ctx context.Context, params *s3.PutBucketPolicyInput, optFns ...func(*s3.Options)) (*s3.PutBucketPolicyOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketPolicy(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketPolicy(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketReplication(ctx context.Context, params *s3.PutBucketReplicationInput, optFns ...func(*s3.Options)) (*s3.PutBucketReplicationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketReplication(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketReplication(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketRequestPayment(ctx context.Context, params *s3.PutBucketRequestPaymentInput, optFns ...func(*s3.Options)) (*s3.PutBucketRequestPaymentOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketRequestPayment(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketRequestPayment(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketTagging(ctx context.Context, params *s3.PutBucketTaggingInput, optFns ...func(*s3.Options)) (*s3.PutBucketTaggingOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketTagging(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketTagging(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketVersioning(ctx context.Context, params *s3.PutBucketVersioningInput, optFns ...func(*s3.Options)) (*s3.PutBucketVersioningOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketVersioning(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketVersioning(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutBucketWebsite(ctx context.Context, params *s3.PutBucketWebsiteInput, optFns ...func(*s3.Options)) (*s3.PutBucketWebsiteOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutBucketWebsite(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutBucketWebsite(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutObject(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutObject(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutObjectAcl(ctx context.Context, params *s3.PutObjectAclInput, optFns ...func(*s3.Options)) (*s3.PutObjectAclOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutObjectAcl(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutObjectAcl(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutObjectLegalHold(ctx context.Context, params *s3.PutObjectLegalHoldInput, optFns ...func(*s3.Options)) (*s3.PutObjectLegalHoldOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutObjectLegalHold(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutObjectLegalHold(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutObjectLockConfiguration(ctx context.Context, params *s3.PutObjectLockConfigurationInput, optFns ...func(*s3.Options)) (*s3.PutObjectLockConfigurationOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutObjectLockConfiguration(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutObjectLockConfiguration(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutObjectRetention(ctx context.Context, params *s3.PutObjectRetentionInput, optFns ...func(*s3.Options)) (*s3.PutObjectRetentionOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutObjectRetention(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutObjectRetention(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutObjectTagging(ctx context.Context, params *s3.PutObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.PutObjectTaggingOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutObjectTagging(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutObjectTagging(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) PutPublicAccessBlock(ctx context.Context, params *s3.PutPublicAccessBlockInput, optFns ...func(*s3.Options)) (*s3.PutPublicAccessBlockOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.PutPublicAccessBlock(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.PutPublicAccessBlock(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) RestoreObject(ctx context.Context, params *s3.RestoreObjectInput, optFns ...func(*s3.Options)) (*s3.RestoreObjectOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.RestoreObject(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.RestoreObject(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) SelectObjectContent(ctx context.Context, params *s3.SelectObjectContentInput, optFns ...func(*s3.Options)) (*s3.SelectObjectContentOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.SelectObjectContent(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.SelectObjectContent(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) UploadPart(ctx context.Context, params *s3.UploadPartInput, optFns ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.UploadPart(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.UploadPart(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) UploadPartCopy(ctx context.Context, params *s3.UploadPartCopyInput, optFns ...func(*s3.Options)) (*s3.UploadPartCopyOutput, error) {
	region := c.getBucketRegion(params.Bucket)
	result, err := c.client.UploadPartCopy(ctx, params, append(optFns, setRegionFn(region))...)
	newRegion := c.followXAmzBucketRegion(params.Bucket, region, err)
	if newRegion != nil {
		result, err = c.client.UploadPartCopy(ctx, params, append(optFns, setRegionFn(newRegion))...)
	}
	return result, err
}

func (c *Client) WriteGetObjectResponse(ctx context.Context, params *s3.WriteGetObjectResponseInput, optFns ...func(*s3.Options)) (*s3.WriteGetObjectResponseOutput, error) {
	return c.client.WriteGetObjectResponse(ctx, params, optFns...)
}
