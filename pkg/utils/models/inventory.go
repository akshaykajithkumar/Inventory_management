package models

type InventoryResponse struct {
	ProductID int
	//Stock     int
}

type Inventory struct {
	ID          uint    `json:"id"`
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}

type UpdateInventory struct {
	ProductName string  `json:"productName"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
	Price       float64 `json:"price"`
}
