package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go-upchi/config"
	"go-upchi/s3client"
)

func main() {
	// Load the .env file
	config.LoadEnv()

	// Initialize Fiber app
	app := fiber.New()

	// Define a route to upload a file to S3
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
				"error": "Failed to open uploaded file",
			})
		}
		defer src.Close()

		// Upload the file to S3
		key, err := s3client.UploadFile(c.Context(), src, file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to upload file: %s", err.Error()),
			})
		}

		// Construct the file URL
		fileURL := fmt.Sprintf("/%s", key)

		return c.JSON(fiber.Map{
			"message": "File uploaded successfully",
			"url":     fileURL,
		})
	})

    // Route for downloading the file using ID
    app.Get("/:file_id", s3client.DownloadFileByID)

	// Listen on port 3000
	log.Fatal(app.Listen(":3000"))
}
