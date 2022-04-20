package entities

type Product struct {
	ProCod            int64   `gorm:"primaryKey" json:"id"`
	ProName           string  `json:"name"`
	ProPrice          float32 `json:"price"`
	ProAvailableStock uint    `json:"stock"`
}
