// controllers/controllers.go
package controllers

import (
	"net/http"

	"github/batuhanzorbeyzengin/insider-messaging-system/services"
	"github/batuhanzorbeyzengin/insider-messaging-system/utils"

	"github.com/gofiber/fiber/v2"
)

func StartStopMessagingSending(c *fiber.Ctx) error {
	status := c.Query("status")

	switch status {
	case "start":
		services.StartMessagingService()
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Start Messaging Sending",
		})

	case "stop":
		services.StopMessagingService()
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Stop Messaging Sending",
		})

	default:
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid status. Use 'start' or 'stop'.",
		})
	}
}

func GetSentMessages(c *fiber.Ctx) error {
	cachedMessages, err := utils.GetCachedMessageDetails()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve cached messages",
		})
	}

	return c.Status(fiber.StatusOK).JSON(cachedMessages)
}
