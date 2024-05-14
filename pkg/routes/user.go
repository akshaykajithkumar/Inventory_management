package routes

import (
	"main/pkg/api/handler"
	"main/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, orderHandler *handler.OrderHandler, inventoryHandler *handler.InventoryHandler) {
	engine.POST("/login", userHandler.Login)
	engine.POST("/signup", userHandler.SignUp)
	engine.POST("/logout", userHandler.Logout)

	engine.Use(middleware.UserAuthMiddleware)
	{
		engine.GET("/inventories/view/:id", inventoryHandler.ViewInventory)

		engine.GET("/profile/orders", orderHandler.GetOrders)
		engine.POST("/profile/orders/place", orderHandler.PlaceOrder)
	}

	engine.GET("/products/search", inventoryHandler.UserSearchProducts)
}
