package interfaces

import (
	"main/pkg/utils/models"
)

type UserUseCase interface {
	Login(user models.UserLogin) (models.TokenUser, error)
	SignUp(user models.UserDetails) (models.TokenUser, error)

	EditUser(id int, userData models.EditUser) error

	UpdateQuantityAdd(id, inv_id int) error
	UpdateQuantityLess(id, inv_id int) error
}
