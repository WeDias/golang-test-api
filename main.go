package main

import (
	"github.com/WeDias/golang-test-api/database"
	"github.com/WeDias/golang-test-api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.InitDatabase()

	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api/v1")
	routes.SetupProductRoutes(api)

	app.Listen(":3000")
}
