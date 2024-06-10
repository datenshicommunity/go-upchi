package s3client

import (
    "fmt"
    "mime/multipart"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/s3"
	"time"
)

// UploadFile uploads a file to the specified S3 bucket
func UploadFile(svc *s3.S3, bucket, key string, file multipart.File, fileHeader *multipart.FileHeader) error {
    defer file.Close()

	// Calculate expiration time (1 minute from now)
    expirationTime := time.Now().Add(1 * time.Minute)

    // Prepare the file for upload
    input := &s3.PutObjectInput{
        Bucket:        aws.String(bucket),
        Key:           aws.String(key),
        Body:          file,
        ContentLength: aws.Int64(fileHeader.Size),
        ContentType:   aws.String(fileHeader.Header.Get("Content-Type")),
		Expires:       &expirationTime,
    }

    // Upload the file to S3
    _, err := svc.PutObject(input)
    if err != nil {
        return fmt.Errorf("unable to upload file to bucket %q, %v", bucket, err)
    }

    return nil
}
