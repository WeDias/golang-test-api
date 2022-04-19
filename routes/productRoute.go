package routes

import (
	"github.com/WeDias/golang-test-api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(api fiber.Router) {
	apiProduct := api.Group("/product")
	apiProduct.Post("/", controllers.NewProduct)
	apiProduct.Get("/", controllers.GetProducts)
	apiProduct.Get("/:id", controllers.GetProduct)
	apiProduct.Put("/:id", controllers.UpdateProduct)
	apiProduct.Delete("/:id", controllers.DeleteProduct)
}
