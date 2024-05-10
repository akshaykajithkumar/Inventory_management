package interfaces

import (
	"main/pkg/utils/models"
)

type UserRepository interface {
	CheckUserAvailability(email string) bool
	UserBlockStatus(email string) (bool, error)
	FindUserByEmail(user models.UserLogin) (models.UserResponse, error)
	SignUp(user models.UserDetails) (models.UserResponse, error)

	GetUserDetails(id int) (models.UserResponse, error)
	FindUserIDByOrderID(orderID int) (int, error)

	FindIdFromPhone(phone string) (int, error)
	EditName(id int, name string) error
	EditEmail(id int, email string) error
	EditPhone(id int, phone string) error
	EditUsername(id int, username string) error

	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error

	FindProductNames(inventory_id int) (string, error)

	FindPrice(inventory_id int) (float64, error)
}
