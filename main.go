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
        objects, err := s3client.ListBucketObjects(svc, bucket)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.JSON(objects)
    })

    // Define a route to upload files to S3
    app.Post("/upload", func(c *fiber.Ctx) error {
        // Get the uploaded file
        file, err := c.FormFile("file")
        if err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "File is required",
            })
        }

        // Open the uploaded file
        src, err := file.Open()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to open file",
            })
        }

        // Define the S3 object key (you can modify this as needed)
        key := file.Filename

        // Upload the file to S3
        err = s3client.UploadFile(svc, bucket, key, src, file)
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.JSON(fiber.Map{
            "message": "File uploaded successfully",
            "file_url": s3client.ConstructObjectURL(key),
        })
    })

    // Listen on port 3000
    log.Fatal(app.Listen("0.0.0.0:3000"))
}
