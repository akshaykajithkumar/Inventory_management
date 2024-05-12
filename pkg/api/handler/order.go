package handler

import (
	"main/pkg/helper"
	services "main/pkg/usecase/interface"
	"main/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase services.OrderUseCase
}

func NewOrderHandler(useCase services.OrderUseCase) *OrderHandler {
	return &OrderHandler{
		orderUseCase: useCase,
	}
}

// @Summary		Get Orders
// @Description	user can view the details of the orders
// @Tags			User
// @Accept			json
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders [get]
func (i *OrderHandler) GetOrders(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "limit number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	id, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	orders, err := i.orderUseCase.GetOrders(id, page, limit)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Place Order
// @Description	user can place orders
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			productid	query	int	true	"product id"
// @Param			quantity	query	int	true	"quantity"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/users/profile/orders/place [post]
func (i *OrderHandler) PlaceOrder(c *gin.Context) {
	// Retrieve user ID from the context
	userID, err := helper.GetUserID(c)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not get userID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	// Parse product ID and quantity from query parameters
	productID, err := strconv.Atoi(c.Query("productid"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Failed to parse product ID", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	quantity, err := strconv.Atoi(c.Query("quantity"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Failed to parse quantity", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	// Call the use case to place the order
	if err := i.orderUseCase.PlaceOrder(userID, productID, quantity); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Failed to place order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	// Respond with success
	successRes := response.ClientResponse(http.StatusOK, "Order placed successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}
