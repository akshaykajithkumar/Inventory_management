package handler

import (
	services "main/pkg/usecase/interface"
	"main/pkg/utils/models"
	"main/pkg/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	InventoryUseCase services.InventoryUseCase
}

func NewInventoryHandler(usecase services.InventoryUseCase) *InventoryHandler {
	return &InventoryHandler{
		InventoryUseCase: usecase,
	}
}

// @Summary		Add Inventory
// @Description	Admin can add new  products
// @Tags			Admin
// @Accept			multipart/form-data
// @Produce		    json
// @Param			product_name	formData	string	true	"product_name"
// @Param			description		formData	string	true	"description"
// @Param			price	formData	string	true	"price"
// @Param			stock		formData	string	true	"stock"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/add [post]
func (i *InventoryHandler) AddInventory(c *gin.Context) {
	//change
	var inventory models.Inventory

	product_name := c.Request.FormValue("product_name")
	description := c.Request.FormValue("description")
	p, err := strconv.Atoi(c.Request.FormValue("price"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	price := float64(p)
	stock, err := strconv.Atoi(c.Request.FormValue("stock"))
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "form file error", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	// Save the uploaded

	inventory.ProductName = product_name
	inventory.Description = description
	inventory.Price = price
	inventory.Stock = stock

	InventoryResponse, err := i.InventoryUseCase.AddInventory(inventory)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Could not add the Inventory", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully added Inventory", InventoryResponse, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		Update Stock
// @Description	Admin can update inventories
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Param			updateinventory	body	models.UpdateInventory	true	"Update Inventory"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/update [patch]
func (i *InventoryHandler) UpdateInventory(c *gin.Context) {
	//change
	inventoryIDstr := c.Query("id")
	invID, err := strconv.Atoi(inventoryIDstr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "id is not valid", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	var invData models.UpdateInventory

	if err := c.BindJSON(&invData); err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	invRes, err := i.InventoryUseCase.UpdateInventory(invID, invData)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not update the inventory stock", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully updated the inventory stock", invRes, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Delete Inventory
// @Description	Admin can delete a product
// @Tags			Admin
// @Accept			json
// @Produce		    json
// @Param			id	query	string	true	"id"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		500	{object}	response.Response{}
// @Router			/admin/inventories/delete [delete]
func (i *InventoryHandler) DeleteInventory(c *gin.Context) {

	inventoryID := c.Query("id")
	err := i.InventoryUseCase.DeleteInventory(inventoryID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully deleted the inventory", nil, nil)
	c.JSON(http.StatusOK, successRes)

}

// @Summary		Search Products
// @Description	client can search with a key and get the list of products similar to that key
// @Tags			Products
// @Accept		json
// @Produce		json
// @Param		page	query  string 	true	"page"
// @Param		limit	query  string 	true	"limit"
// @Param		searchkey	query  string 	true	"searchkey"
// @Param		sortBY	query  string 	false	"sortBY (asc/desc) - Sort by price in ascending (asc) or descending (desc) order"
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router		/products/search [get]
func (i *InventoryHandler) SearchProducts(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "limit number not in the right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	searchKey := c.Query("searchkey")
	sortBY := c.Query("sortBY") // Add this line to get the sorting parameter

	results, err := i.InventoryUseCase.SearchProducts(searchKey, page, limit, sortBY)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not retrieve the records", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	successRes := response.ClientResponse(http.StatusOK, "Successfully got all records", results, nil)
	c.JSON(http.StatusOK, successRes)
}

// @Summary		View Inventory
// @Description	View details of an inventory by ID
// @Tags			Inventory
// @Accept			json
// @Produce		json
// @Param			id	path	string	true	"Inventory ID"
// @Security		Bearer
// @Success		200	{object}	response.Response{}
// @Failure		400	{object}	response.Response{}
// @Router			/users/inventories/view/{id} [get]
func (i *InventoryHandler) ViewInventory(c *gin.Context) {
	// Extract inventory ID from path parameters
	inventoryID := c.Param("id")

	// Use the inventory ID to fetch details from the use case layer
	inventory, err := i.InventoryUseCase.GetInventoryByID(inventoryID)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "Failed to fetch inventory details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}

	// Respond with the fetched inventory details
	successRes := response.ClientResponse(http.StatusOK, "Successfully fetched inventory details", inventory, nil)
	c.JSON(http.StatusOK, successRes)
}
