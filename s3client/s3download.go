// s3download.go

package s3client

import (
	"context"
	"io"
	"net/http"
	"os"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

// DownloadFileByID fetches a file from S3 by its ID, returns it as a response,
// and deletes the object from S3 after successful download
func DownloadFileByID(c *fiber.Ctx) error {
	fileID := c.Params("file_id")

	obj, err := s3client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String("tmp/" + fileID),
	})
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	// Read the file content
	fileContent, err := io.ReadAll(obj.Body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to read file content"})
	}

	// Set response headers
	c.Set("Content-Type", "application/octet-stream")
	c.Set("Content-Disposition", "attachment; filename="+obj.Metadata["file_name"])

	// Send the file content as response
	if err := c.Send(fileContent); err != nil {
		return err
	}

	// Delete the object from S3
	_, delErr := s3client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:    aws.String("tmp/" + fileID),
	})
	if delErr != nil {
		// Log the error, but continue since the file was successfully downloaded
		log.Printf("Failed to delete object from S3: %v\n", delErr)
	}

	return nil
}
