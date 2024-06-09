package s3client

import (
    "fmt"
    "log"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/s3"
)

// ListBucketObjects lists all objects in the specified S3 bucket
func ListBucketObjects(svc *s3.S3, bucket, customEndpoint string) {
    input := &s3.ListObjectsV2Input{
        Bucket: aws.String(bucket),
    }

    result, err := svc.ListObjectsV2(input)
    if err != nil {
        log.Fatalf("Unable to list items in bucket %q, %v", bucket, err)
    }

    fmt.Println("Objects in bucket", bucket)
    for _, item := range result.Contents {
        url := ConstructObjectURL(customEndpoint, bucket, *item.Key)
        fmt.Printf("Name: %s, Size: %d, LastModified: %s, URL: %s\n", *item.Key, *item.Size, item.LastModified, url)
    }
}

// ConstructObjectURL constructs the URL for an S3 object given the endpoint, bucket, and object key
func ConstructObjectURL(customEndpoint, bucket, key string) string {
    return fmt.Sprintf("%s/%s/%s", customEndpoint, bucket, key)
}
