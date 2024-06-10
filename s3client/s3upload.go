package s3client

import (
    "context"
    "fmt"
    "mime/multipart"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
)

// UploadFile uploads a file to the specified S3 bucket
func UploadFile(client *s3.Client, bucket, key string, file multipart.File, fileHeader *multipart.FileHeader) error {
    defer file.Close()

    uploader := manager.NewUploader(client)
    input := &s3.PutObjectInput{
        Bucket:        aws.String(bucket),
        Key:           aws.String(key),
        Body:          file,
        ContentLength: aws.Int64(fileHeader.Size),
        ContentType:   aws.String(fileHeader.Header.Get("Content-Type")),
    }

    // Upload the file to S3
    _, err := uploader.Upload(context.TODO(), input)
    if err != nil {
        return fmt.Errorf("unable to upload file to bucket %q, %v", bucket, err)
    }

    return nil
}
