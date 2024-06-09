package s3client

import (
    "log"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

// CreateS3Session initializes a new S3 session with the provided credentials and endpoint
func CreateS3Session(accessKeyID, secretAccessKey, region, customEndpoint string) *s3.S3 {
    sess, err := session.NewSession(&aws.Config{
        Region:           aws.String(region),
        Endpoint:         aws.String(customEndpoint),
        Credentials:      credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
        S3ForcePathStyle: aws.Bool(true),
    })

    if err != nil {
        log.Fatalf("Failed to create session: %v", err)
    }

    return s3.New(sess)
}
