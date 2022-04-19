package services

import (
	"errors"

	"github.com/WeDias/golang-test-api/database"
	"github.com/WeDias/golang-test-api/entities"
)

func NewProduct(product *entities.Product) (*entities.Product, error) {
	var err error
	database.Conn.Create(&product)
	return product, err
}

func GetProduct(productId string) (*entities.Product, error) {
	var err error
	var product *entities.Product
	database.Conn.Find(&product, productId)
	if product.ID == 0 {
		err = errors.New("product not found")
	}
	return product, err
}

func GetProducts() (*[]entities.Product, error) {
	var err error
	var products *[]entities.Product
	database.Conn.Find(&products)
	return products, err
}

func UpdateProduct(newProduct *entities.Product, productId string) (*entities.Product, error) {
	product, err := GetProduct(productId)
	database.Conn.Model(product).Updates(newProduct)
	return product, err
}

func DeleteProduct(productId string) (*entities.Product, error) {
	product, err := GetProduct(productId)
	database.Conn.Delete(product)
	return product, err
}
