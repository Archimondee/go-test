package utils

import (
	"cloud.google.com/go/storage"
	"context"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
	"io"
	"time"
)

func HandleFileUploadToBucket(c *fiber.Ctx) error {
	bucket := "uber-61648.appspot.com" // your bucket name

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile("keys.json"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"error":   true,
		})
	}

	file, err := c.FormFile("image")

	if file != nil {
		contentType := file.Header.Get("Content-Type")
		if contentType != "image/jpeg" && contentType != "image/png" {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(
				fiber.Map{
					"message": "Error file extension",
					"status":  fiber.StatusUnprocessableEntity,
				})
		}
		fileContent, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"error":   true,
			})
		}

		defer fileContent.Close()

		ctx, cancel := context.WithTimeout(ctx, time.Second*50)
		defer cancel()

		object := storageClient.Bucket(bucket).Object(file.Filename)
		sw := object.NewWriter(ctx)

		if _, err := io.Copy(sw, fileContent); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"error":   true,
			})
		}

		if err := sw.Close(); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
				"error":   true,
			})
		}

		u := "https://firebasestorage.googleapis.com/v0/b/" + bucket + "/o/" + object.ObjectName() + "?alt=media"

		c.Locals("filename", u)
	}

	//return c.Status(fiber.StatusOK).JSON(fiber.Map{
	//	"message":  "file uploaded successfully",
	//	"pathname": u,
	//})

	return c.Next()
}
