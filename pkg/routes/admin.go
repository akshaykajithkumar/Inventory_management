package routes

import (
	"main/pkg/api/handler"
	"main/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(engine *gin.RouterGroup, adminHandler *handler.AdminHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler) {
	engine.POST("/adminlogin", adminHandler.LoginHandler)

	engine.Use(middleware.AdminAuthMiddleware)
	{
		engine.POST("/logout", adminHandler.Logout)

		inventorymanagement := engine.Group("/inventories")
		{
			//inventorymanagement.GET("", inventoryHandler.ListProducts)
			//inventorymanagement.GET("/details", inventoryHandler.ShowIndividualProducts)
			inventorymanagement.POST("/add", inventoryHandler.AddInventory)
			inventorymanagement.PATCH("/update", inventoryHandler.UpdateInventory)

			inventorymanagement.DELETE("/delete", inventoryHandler.DeleteInventory)
		}
		orders := engine.Group("/orders")
		{
			//orders.PATCH("/edit/status", orderHandler.EditOrderStatus)
			// orders.PATCH("/edit/mark-as-paid", orderHandler.MarkAsPaid)
			orders.GET("", orderHandler.AdminOrders)
		}

		sales := engine.Group("/sales")
		{
			sales.GET("/daily", orderHandler.AdminSalesDailyReport)
			sales.GET("/weekly", orderHandler.AdminSalesWeeklyReport)
			sales.GET("/monthly", orderHandler.AdminSalesMonthlyReport)
			sales.GET("/annual", orderHandler.AdminSalesAnnualReport)
			sales.POST("/custom", orderHandler.AdminSalesCustomReport)
		}
		products := engine.Group("/products")
		{
			// products.GET("/details", inventoryHandler.ShowIndividualProducts)
			products.GET("/search", inventoryHandler.SearchProducts)

		}

	}
}
