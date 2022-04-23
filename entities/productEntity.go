package entities

type Product struct {
	Cod            int64   `json:"id"    gorm:"column:pro_cod; primaryKey; <-:create"`
	Name           string  `json:"name"  gorm:"column:pro_name"`
	Price          float32 `json:"price" gorm:"column:pro_price"`
	AvailableStock int32   `json:"stock" gorm:"column:pro_available_stock"`
}
