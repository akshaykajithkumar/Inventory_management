package interfaces

import (
	"main/pkg/domain"
	"time"
)

type OrderRepository interface {
	GetOrders(id, page, limit int) ([]domain.Order, error)
	GetProductsQuantity() ([]domain.ProductReport, error)
	GetOrdersInRange(startDate, endDate time.Time) ([]domain.Order, error)
	GetProductNameFromID(id int) (string, error)
	EditOrderStatus(status string, id int) error
	FindUserIdFromOrderID(orderID int) (int, error)
	FindAmountFromOrderID(orderID int) (float64, error)
	GetProductByID(productID int) (bool, error)
	CheckProductAvailability(productID int, quantity int) (bool, error)
	FindAmountFromProductID(productID int) (float64, error)
	PlaceOrder(userID, productID, quantity int, price float64) error
}
