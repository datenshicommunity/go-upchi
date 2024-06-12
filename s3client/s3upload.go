package s3client

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

// PutObject uploads a file to the specified S3 bucket
func PutObject(ctx context.Context, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	client := GetS3Client() // Get the initialized S3 client

	// Upload the file to S3
	resp, err := client.PutObject(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("unable to upload file to bucket %q, %v", *input.Bucket, err)
	}

	return resp, nil
}

// UploadFile uploads a file to the specified S3 bucket
func UploadFile(ctx context.Context, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	defer file.Close()

	fileID, err := gonanoid.Generate("abcdefghijklmnopqrstuvwxyz1234567890", 12)
	if err != nil {
		return "", fmt.Errorf("failed to generate file ID: %v", err)
	}

	expAt := time.Now().Add(time.Hour * 1)

	key := fmt.Sprint(fileID)
	_, err = PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:         aws.String(key),
		Body:        file,
		Metadata:    map[string]string{"file_name": fileHeader.Filename},
		Expires:     aws.Time(expAt),
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
	})
	if err != nil {
		return "", fmt.Errorf("unable to upload file to S3: %v", err)
	}

	return key, nil
}
