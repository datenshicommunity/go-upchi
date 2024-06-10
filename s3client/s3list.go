package s3client

import (
    "context"
    "fmt"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Object holds information about an object stored in S3
type S3Object struct {
    Name         string    `json:"name"`
    Size         int64     `json:"size"`
    LastModified time.Time `json:"last_modified"`
    URL          string    `json:"url"`
}

// ListBucketObjects lists all objects in the specified S3 bucket
func ListBucketObjects(client *s3.Client, bucket, customEndpoint string) ([]S3Object, error) {
    input := &s3.ListObjectsV2Input{
        Bucket: aws.String(bucket),
    }

    result, err := client.ListObjectsV2(context.TODO(), input)
    if err != nil {
        return nil, fmt.Errorf("unable to list items in bucket %q, %v", bucket, err)
    }

    var objects []S3Object
    for _, item := range result.Contents {
        url := ConstructObjectURL(customEndpoint, bucket, *item.Key)
        obj := S3Object{
            Name:         *item.Key,
            Size:         aws.ToInt64(item.Size), // Use aws.ToInt64 to safely dereference the pointer
            LastModified: *item.LastModified,
            URL:          url,
        }
        objects = append(objects, obj)
    }

    return objects, nil
}

// ConstructObjectURL constructs the URL for an S3 object given the endpoint, bucket, and object key
func ConstructObjectURL(customEndpoint, bucket, key string) string {
    return fmt.Sprintf("%s/%s/%s", customEndpoint, bucket, key)
}
