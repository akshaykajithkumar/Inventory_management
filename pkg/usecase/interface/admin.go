package interfaces

import (
	"main/pkg/domain"
	"main/pkg/utils/models"
)

type AdminUseCase interface {
	LoginHandler(adminDetails models.AdminLogin) (models.TokenAdmin, error)
	GetUsers(page int, limit int) ([]models.UserDetailsAtAdmin, error)
	AdminOrders(page, limit int, status string) ([]domain.Order, error)
	GetOrder(id int) (domain.Order, error)
	ChangeOrderStatus(orderID, status string) error
	UserStats() (models.UserStats, error)
	OrderStats() (models.OrderStats, error)
	InventoryStats() (models.InventoryStats, error)
}
