package handler

import (
	services "main/pkg/usecase/interface"
	models "main/pkg/utils/models"
	"main/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
}

func NewAdminHandler(usecase services.AdminUseCase) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
	}
}

// @Summary		Admin Login
// @Description	Login handler for admins
// @Tags			Admin
// @Accept			json
// @Produce		json
// @Param			admin	body		models.AdminLogin	true	"Admin login details"
// @Success		200		{object}	response.Response{}
// @Failure		500		{object}	response.Response{}
// @Router			/admin/adminlogin [post]
func (ad *AdminHandler) LoginHandler(c *gin.Context) { // login handler for the admin

	// var adminDetails models.AdminLogin
	var adminDetails models.AdminLogin
	if err := c.BindJSON(&adminDetails); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}

	admin, err := ad.adminUseCase.LoginHandler(adminDetails)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "cannot authenticate user", nil, err.Error())
		c.JSON(http.StatusBadRequest, errRes)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", admin.AccessToken, 0, "/", "", false, true)
	c.SetCookie("Refreshtoken", admin.RefreshToken, 0, "/", "", false, true)

	successRes := response.ClientResponse(http.StatusOK, "Admin authenticated successfully", admin, nil)
	c.JSON(http.StatusOK, successRes)

}

// Logout is a handler for admin logout
// @Summary		admin Logout
// @Description	admin can logout
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/logout [post]
func (ad *AdminHandler) Logout(c *gin.Context) {
	// Clear the access token and refresh token cookies
	c.SetCookie("Authorization", "", -1, "/", "", false, true)
	c.SetCookie("Refreshtoken", "", -1, "/", "", false, true)

	successRes := response.ClientResponse(http.StatusOK, "User successfully logged out", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Admin Orders
// @Description	Admin can view the orders according to status
// @Tags			Admin
// @Produce		    json
// @Param			page	query  string 	true	"page"
// @Param			limit	query  string 	true	"limit"
// @Param			status	query  string	true	"status"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/orders [get]
func (i *AdminHandler) AdminOrders(c *gin.Context) {
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
	status := c.Query("status")

	orders, err := i.adminUseCase.AdminOrders(page, limit, status)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", orders, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get Specific Order
// @Description	Get a specific order by its ID
// @Tags			Admin
// @Produce		json
// @Param			id	path	string	true	"Order ID"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Failure		404	{object}	response.Response{}
// @Router			/admin/orders/{id} [get]
func (i *AdminHandler) GetOrder(c *gin.Context) {

	orderIDStr := c.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "order ID not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	// Call your use case method to retrieve the order by ID
	order, err := i.adminUseCase.GetOrder(orderID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not retrieve order", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully retrieved order", order, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Change Order Status
// @Description	Change the status of a specific order
// @Tags			Admin
// @Accept		json
// @Produce		json
// @Param			id	path	string	true	"Order ID"
// @Param			status	query	string	true	"New status"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Failure		404	{object}	response.Response{}
// @Router			/admin/orders/{id}/status [put]
func (i *AdminHandler) ChangeOrderStatus(c *gin.Context) {
	orderID := c.Param("id")
	status := c.Query("status")

	// Call your use case method to change the order status
	err := i.adminUseCase.ChangeOrderStatus(orderID, status)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Failed to change order status", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	// Success case
	successRes := response.ClientResponse(http.StatusOK, "Order status changed successfully", nil, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Get User Statistics
// @Description	Fetch user statistics and return them as JSON.
// @Tags			Admin
// @Accept		json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/user/stats [get]
func (a *AdminHandler) UserStats(c *gin.Context) {
	// Fetch user statistics using the UserStats use case function
	userStats, err := a.adminUseCase.UserStats()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Message: "Could not fetch user Stats ", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Response{Message: "Fetched userStats", Data: userStats})
}

// @Summary		Get Order Statistics
// @Description	Fetch order statistics and return them as JSON.
// @Tags			Admin
// @Accept		json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/order/stats [get]
func (a *AdminHandler) OrderStats(c *gin.Context) {
	// Fetch order statistics using the OrderStats use case function
	orderStats, err := a.adminUseCase.OrderStats()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Message: "Could not fetch orderStats ", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Response{Message: "Fetched orderStats", Data: orderStats})
}

// @Summary		Get Inventory Statistics
// @Description	Fetch inventory statistics and return them as JSON.
// @Tags			Admin
// @Accept		json
// @Produce		json
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/admin/inventory/stats [get]
func (a *AdminHandler) InventoryStats(c *gin.Context) {
	// Fetch inventory statistics using the InventoryStats use case function
	inventoryStats, err := a.adminUseCase.InventoryStats()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Response{Message: "Could not fetch inventory Stats ", Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Response{Message: "Fetched inventoryStats", Data: inventoryStats})
}
