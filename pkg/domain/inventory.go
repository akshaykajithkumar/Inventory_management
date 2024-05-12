package domain

// Inventory represents the Products in the website
type Inventory struct {
	ID          uint    `json:"id" gorm:"unique;not null"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
