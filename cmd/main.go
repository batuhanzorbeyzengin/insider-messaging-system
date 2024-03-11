package main

import (
	"github/batuhanzorbeyzengin/insider-messaging-system/configs"
	"github/batuhanzorbeyzengin/insider-messaging-system/database"
	"github/batuhanzorbeyzengin/insider-messaging-system/routes"
	"github/batuhanzorbeyzengin/insider-messaging-system/utils"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	err := database.Connect()
	if err != nil {
		panic(err)
	}

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize Redis connection
	err = utils.InitRedis(config.Redis)
	if err != nil {
		log.Fatalf("Error initializing Redis connection: %v", err)
	}

	routes.SetupRoutes(app)

	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
