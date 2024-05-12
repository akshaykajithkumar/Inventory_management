package interfaces

import (
	"main/pkg/domain"
	"main/pkg/utils/models"
)

type AdminRepository interface {
	LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error)
	GetUserByID(id string) (domain.User, error)

	GetUsers(page int, limit int) ([]models.UserDetailsAtAdmin, error)
	AdminGetOrder(id int) (domain.Order, error)
	ChangeOrderStatus(orderID, status string) error
	AdminOrders(page, limit int, status string) ([]domain.Order, error)
	CountUsers() (int, error)
	GetTotalRevenue() (float32, error)
	GetTotalRevenueToday() (float32, error)
	CountOrders() (int, error)
	CountProducts() (int, error)
	CountStock() (int, error)
	GetTotalPrice() (float32, error)
	GetMostSoldProduct() (string, error)
	GetTrendingProduct() (string, error)
}
