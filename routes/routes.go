package routes

import (
	"github/batuhanzorbeyzengin/insider-messaging-system/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {

	// Middleware
	app.Use(logger.New())

	api := app.Group("/api/v1")

	messageGroup := api.Group("/messages")
	messageGroup.Post("/send", controllers.StartStopMessagingSending)
	messageGroup.Get("/sent", controllers.GetSentMessages)
}
