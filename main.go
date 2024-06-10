package main

import (
    "fmt"
    "log"
    "time"

    "go-upchi/config"
    "go-upchi/s3client"

    "github.com/gofiber/fiber/v2"
)

func main() {
    // Load the .env file
    config.LoadEnv()

    // Retrieve the environment variables
    accessKeyID := config.GetEnv("AWS_ACCESS_KEY_ID")
    secretAccessKey := config.GetEnv("AWS_SECRET_ACCESS_KEY")
    bucket := config.GetEnv("S3_BUCKET")
    accountId := config.GetEnv("ACCOUNT_ID")

    // Create a new S3 client
    client, err := s3client.CreateS3Client(accessKeyID, secretAccessKey, accountId)
    if err != nil {
        log.Fatalf("Failed to create S3 client: %v", err)
    }

    // Initialize Fiber app
    app := fiber.New()

    // Define a route to list S3 bucket objects
    app.Get("/list", func(c *fiber.Ctx) error {
        objects, err := s3client.ListBucketObjects(client, bucket, fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId))
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(objects)
    })

    // Define a route to generate presigned URLs for uploading files
    app.Post("/upload", func(c *fiber.Ctx) error {
        // Get the uploaded file
        file, err := c.FormFile("file")
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "File is required",
            })
        }

        fileName := file.Filename

        // Generate a presigned URL for the file upload
        presignedURL, err := s3client.GeneratePresignedURL(client, bucket, fileName, 1*time.Minute)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.JSON(fiber.Map{
            "presigned_url": presignedURL,
            "file_url": s3client.ConstructObjectURL(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId), bucket, fileName),
        })
    })

    // Listen on port 3000
    log.Fatal(app.Listen(":3000"))
}
