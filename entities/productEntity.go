package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name           string  `json:"name"`
	Price          float32 `json:"price"`
	AvailableStock uint    `json:"stock"`
}
