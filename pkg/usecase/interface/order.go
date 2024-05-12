package interfaces

import (
	"main/pkg/domain"
)

type OrderUseCase interface {
	PlaceOrder(userID, productID, quantity int) error
	GetOrders(id, page, limit int) ([]domain.Order, error)
	//OrderItemsFromCart(userid int, order models.Order, coupon string) (string, error)

	//EditOrderStatus(status string, id int) error
	//MarkAsPaid(orderID int) error

	//CustomDateOrders(dates models.CustomDates) (domain.SalesReport, error)
	//ReturnOrder(id int) error

}
