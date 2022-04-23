package main

import (
	"fmt"
	"os"

	"github.com/WeDias/golang-test-api/routes"
	"github.com/WeDias/golang-test-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New())

	app.Use(func(ctx *fiber.Ctx) error {
		if err := ctx.Next(); err != nil {
			return ctx.Status(404).JSON(fiber.Map{
				"error": fmt.Sprintf("invalid endpoint '%s'", ctx.Path()),
			})
		}
		return nil
	})

	api := app.Group("/api/v1")
	routes.SetupProductRoutes(api)

	utils.InfoLog.Println("starting the server on port", os.Getenv("API_PORT"))
	if err := app.Listen(os.Getenv("API_PORT")); err != nil {
		utils.ErrorLog.Panic(err.Error())
	}
}
