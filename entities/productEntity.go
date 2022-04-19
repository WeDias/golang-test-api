package entities

type Product struct {
	Id             uint64  `gorm:"primaryKey" json:"id"`
	Name           string  `json:"name"`
	Price          float32 `json:"price"`
	AvailableStock uint    `json:"stock"`
}
