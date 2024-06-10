package s3client

import (
    "context"
    "fmt"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/credentials"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

// CreateS3Client initializes an S3 client using the provided credentials and custom endpoint
func CreateS3Client(accessKeyID, secretAccessKey, accountId string) (*s3.Client, error) {
    r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
        if service == s3.ServiceID {
            return aws.Endpoint{
                URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId),
            }, nil
        }
        return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
    })

    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion("auto"),
        config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
        config.WithEndpointResolverWithOptions(r2Resolver),
    )
    if err != nil {
        return nil, fmt.Errorf("failed to create S3 client: %v", err)
    }

    return s3.NewFromConfig(cfg), nil
}
