package s3client

import (
    "context"
    "fmt"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

// GeneratePresignedURL generates a presigned URL for uploading a file to S3
func GeneratePresignedURL(client *s3.Client, bucket, key string, expiration time.Duration) (string, error) {
    presignClient := s3.NewPresignClient(client)
    params := &s3.PutObjectInput{
        Bucket: aws.String(bucket),
        Key:    aws.String(key),
    }

    presignedURL, err := presignClient.PresignPutObject(context.TODO(), params, func(opts *s3.PresignOptions) {
        opts.Expires = expiration
    })
    if err != nil {
        return "", fmt.Errorf("failed to generate presigned URL: %v", err)
    }

    return presignedURL.URL, nil
}
