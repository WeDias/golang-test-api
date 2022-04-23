package services

import (
	"errors"

	"github.com/WeDias/golang-test-api/database"
	"github.com/WeDias/golang-test-api/entities"
	"github.com/WeDias/golang-test-api/utils"
)

func NewProduct(product *entities.Product) (*entities.Product, error) {
	var err error
	tx := database.Conn.Create(&product)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	}
	return product, err
}

func GetProduct(productId string) (*entities.Product, error) {
	var err error
	var product *entities.Product
	tx := database.Conn.Find(&product, productId)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if product.Cod == 0 {
		err = errors.New("product not found")
	}
	return product, err
}

func GetProducts() (*[]entities.Product, error) {
	var err error
	var products *[]entities.Product
	tx := database.Conn.Order("pro_cod").Find(&products)
	if tx.Error != nil {
		err = tx.Error
		utils.WarningLog.Println(err.Error())
	} else if *products == nil {
		err = errors.New("products not found")
	}
	return products, err
}

func UpdateProduct(newProduct *entities.Product, productId string) (*entities.Product, error) {
	product, err := GetProduct(productId)
	if err == nil {
		tx := database.Conn.Model(product).Updates(newProduct)
		if tx.Error != nil {
			err = tx.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return product, err
}

func DeleteProduct(productId string) (*entities.Product, error) {
	product, err := GetProduct(productId)
	if err == nil {
		tx := database.Conn.Delete(product)
		if tx.Error != nil {
			err = tx.Error
			utils.WarningLog.Println(err.Error())
		}
	}
	return product, err
}
