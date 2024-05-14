package interfaces

import (
	"main/pkg/utils/models"
)

type UserRepository interface {
	CheckUserAvailability(email string) bool
	FindUserByEmail(user models.UserLogin) (models.UserResponse, error)
	SignUp(user models.UserDetails) (models.UserResponse, error)
	GetUserDetails(id int) (models.UserResponse, error)
	FindUserIDByOrderID(orderID int) (int, error)
	FindProductNames(inventory_id int) (string, error)
	FindPrice(inventory_id int) (float64, error)
}
