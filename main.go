package main

import (
	"go-upchi/config"
	"go-upchi/s3client"
	"log"

	"github.com/gofiber/fiber/v2"
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

    // Fiber instance
    app := fiber.New()

    // Route
    app.Get("/list", func(c *fiber.Ctx) error {
        objects, err := s3client.ListBucketObjects(svc, bucket, customEndpoint)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(objects)
    })

    // Listen on port 3000
    log.Fatal(app.Listen(":3000"))
}
