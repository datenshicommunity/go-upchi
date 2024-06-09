package main

import (
    "go-upchi/config"
    "go-upchi/s3client"
)

func main() {
    // Load the .env file
    config.LoadEnv()

    // Retrieve the environment variables
    accessKeyID := config.GetEnv("AWS_ACCESS_KEY_ID")
    secretAccessKey := config.GetEnv("AWS_SECRET_ACCESS_KEY")
    region := config.GetEnv("AWS_REGION")
    bucket := config.GetEnv("S3_BUCKET")
    customEndpoint := config.GetEnv("S3_ENDPOINT")

    // Create a new S3 session
    svc := s3client.CreateS3Session(accessKeyID, secretAccessKey, region, customEndpoint)

    // List the objects in the specified bucket
    s3client.ListBucketObjects(svc, bucket, customEndpoint)
}
