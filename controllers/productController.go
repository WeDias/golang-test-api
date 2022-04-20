package controllers

import (
	"github.com/WeDias/golang-test-api/entities"
	"github.com/WeDias/golang-test-api/services"
	"github.com/gofiber/fiber/v2"
)

func NewProduct(ctx *fiber.Ctx) error {
	var newProduct *entities.Product
	if err := ctx.BodyParser(&newProduct); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := services.NewProduct(newProduct)
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(product)
}

func GetProduct(ctx *fiber.Ctx) error {
	product, err := services.GetProduct(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(product)
}

func GetProducts(ctx *fiber.Ctx) error {
	products, err := services.GetProducts()
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(products)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	var product *entities.Product
	if err := ctx.BodyParser(&product); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := services.UpdateProduct(product, ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(product)
}

func DeleteProduct(ctx *fiber.Ctx) error {
	product, err := services.DeleteProduct(ctx.Params("id"))
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(product)
}
