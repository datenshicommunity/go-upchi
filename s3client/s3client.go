package s3client

import (
	"context"
	"log"
    "os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	configlokal "go-upchi/config" // configlokal is used to load environment variables
)

var s3client *s3.Client

func init() {
	// Load environment variables
	configlokal.LoadEnv()

	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: os.Getenv("S3_ENDPOINT_URL"),
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("S3_ACCESS_KEY"), os.Getenv("S3_SECRET_KEY"), "")),
        config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	s3client = s3.NewFromConfig(cfg)
}

// GetS3Client returns the initialized S3 client
func GetS3Client() *s3.Client {
	return s3client
}
