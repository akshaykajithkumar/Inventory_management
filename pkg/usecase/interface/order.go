package interfaces

import (
	"main/pkg/domain"
)

type OrderUseCase interface {
	PlaceOrder(userID, productID, quantity int) error
	GetOrders(id, page, limit int) ([]domain.Order, error)
}
