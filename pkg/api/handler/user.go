package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "main/pkg/usecase/interface"
	"main/pkg/utils/models"
	"main/pkg/utils/response"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// Login is a handler for user login
// @Summary		User Login
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			login  body  models.UserLogin  true	"login"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/login [post]
func (i *UserHandler) Login(c *gin.Context) {
	var user models.UserLogin

	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	userDetails, err := i.userUseCase.Login(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not be logged in", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "User successfully logged in", userDetails, nil)

	c.SetCookie("Authorization", userDetails.AccessToken, 0, "/", "", false, true)
	c.SetCookie("Refreshtoken", userDetails.RefreshToken, 0, "/", "", false, true)
	c.JSON(http.StatusOK, successRes)
}

// Signup is a handler for user Registration
// @Summary		User Signup
// @Description	user can signup by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			signup  body  models.UserDetails  true	"signup"
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/signup [post]
func (i *UserHandler) SignUp(c *gin.Context) {

	var user models.UserDetails
	// bind the user details to the struct
	if err := c.BindJSON(&user); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	// business logic goes inside this function
	userCreated, err := i.userUseCase.SignUp(user)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "User could not signed up", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	successRes := response.ClientResponse(http.StatusCreated, "User successfully signed up", userCreated, nil)
	c.JSON(http.StatusCreated, successRes)

}

// // @Summary		Add quantity in cart by one
// // @Description	user can add 1 quantity of product to their cart
// // @Tags			User
// // @Accept			json
// // @Produce		    json
// // @Param			inventory	query	string	true	"inventory id"
// // @Security		Bearer
// // @Success		200	{object}	response.Response{}
// // @Failure		500	{object}	response.Response{}
// // @Router			/users/cart/updateQuantity/plus [put]
// func (i *UserHandler) UpdateQuantityAdd(c *gin.Context) {
// 	id, err := helper.GetUserID(c)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	cartID, err := i.userUseCase.GetCartID(id)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	inv, err := strconv.Atoi(c.Query("inventory"))
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	if err := i.userUseCase.UpdateQuantityAdd(cartID, inv); err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "could not Add the quantity", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Successfully added quantity", nil, nil)
// 	c.JSON(http.StatusOK, successRes)
// }

// // @Summary		Subtract quantity in cart by one
// // @Description	user can subtract 1 quantity of product from their cart
// // @Tags			User
// // @Accept			json
// // @Produce		    json
// // @Param			inventory	query	string	true	"inventory id"
// // @Security		Bearer
// // @Success		200	{object}	response.Response{}
// // @Failure		500	{object}	response.Response{}
// // @Router			/users/cart/updateQuantity/minus [put]
// func (i *UserHandler) UpdateQuantityLess(c *gin.Context) {
// 	id, err := helper.GetUserID(c)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	cartID, err := i.userUseCase.GetCartID(id)
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	inv, err := strconv.Atoi(c.Query("inventory"))
// 	if err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "check parameters properly", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	if err := i.userUseCase.UpdateQuantityLess(cartID, inv); err != nil {
// 		errorRes := response.ClientResponse(http.StatusBadRequest, "could not  subtract quantity", nil, err.Error())
// 		c.JSON(http.StatusBadRequest, errorRes)
// 		return
// 	}

// 	successRes := response.ClientResponse(http.StatusOK, "Successfully subtracted quantity", nil, nil)
// 	c.JSON(http.StatusOK, successRes)
// }

// Login is a handler for user login
// @Summary		User Logout
// @Description	user can log in by giving their details
// @Tags			User
// @Accept			json
// @Produce		    json
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/logout [post]
func (i *UserHandler) Logout(c *gin.Context) {

	successRes := response.ClientResponse(http.StatusOK, "User successfully logged out", nil, nil)
	//c.SetCookie("Authorization", userDetails.Token, 3600, "/", "main.online", true, false)
	c.SetCookie("Authorization", "", -1, "", "", true, false)
	c.JSON(http.StatusOK, successRes)
}
