package domain

// Inventory represents the Products in the website
type Inventory struct {
	ID uint `json:"id" gorm:"unique;not null"`

	ProductName string `json:"product_name"`
	Description string `json:"description"`

	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}

// Category represents the category of product
type Category struct {
	ID       uint   `json:"id" gorm:"unique;not null"`
	Category string `json:"category" gorm:"unique;not null"`
}

type Image struct {
	ID          int       `json:"id" gorm:"unique;not null"`
	InventoryID int       `json:"inventory_id"`
	Inventory   Inventory `json:"-" gorm:"foriegnkey:InventoryID"`
	ImageURL    string    `json:"imageurl"`
}
