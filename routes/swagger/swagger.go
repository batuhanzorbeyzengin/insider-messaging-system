package swagger

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// SwaggerConfig configures the Swagger integration for the Fiber application
func SwaggerConfig(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault) // Register Swagger handler
}
