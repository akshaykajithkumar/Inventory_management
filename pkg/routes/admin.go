package routes

import (
	"main/pkg/api/handler"
	"main/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, inventoryHandler *handler.InventoryHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		engine.POST("/logout", adminHandler.Logout)

		inventorymanagement := engine.Group("/inventories")
		{
			inventorymanagement.POST("/add", inventoryHandler.AddInventory)
			inventorymanagement.PATCH("/update", inventoryHandler.UpdateInventory)
			inventorymanagement.DELETE("/delete", inventoryHandler.DeleteInventory)
		}

		orders := engine.Group("/orders")
		{
			orders.GET("", adminHandler.AdminOrders)
			orders.GET("/:id", adminHandler.GetOrder)
			orders.PUT("/:id/status", adminHandler.ChangeOrderStatus)
		}
	}
}
