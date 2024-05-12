package domain

import "time"

type Order struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID      int       `json:"user_id" gorm:"not null"`
	InventoryID int       `json:"inventory_id"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	User        User      `json:"-" gorm:"foreignkey:UserID"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	OrderedAt   time.Time `json:"orderedAt"`
	OrderStatus string    `json:"order_status" gorm:"order_status:4;default:'PENDING';check:order_status IN ('PENDING', 'SHIPPED','DELIVERED','CANCELED','RETURNED')"`
}

// OrderItem represents the product details of the order
type OrderItem struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderID     int       `json:"order_id"`
	Order       Order     `json:"-" gorm:"foreignkey:OrderID"`
	InventoryID int       `json:"inventory_id"`
	Inventory   Inventory `json:"-" gorm:"foreignkey:InventoryID"`
	Quantity    int       `json:"quantity"`
	TotalPrice  float64   `json:"total_price"`
}

type ProductReport struct {
	InventoryID int
	Quantity    int
}
