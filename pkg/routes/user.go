package routes

import (
	"main/pkg/api/handler"
	"main/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(engine *gin.RouterGroup, userHandler *handler.UserHandler, inventoryHandler *handler.InventoryHandler, orderHandler *handler.OrderHandler) {
	engine.POST("/login", userHandler.Login)
	engine.POST("/signup", userHandler.SignUp)

	// Auth middleware
	engine.Use(middleware.UserAuthMiddleware)
	{
		engine.POST("/logout", userHandler.Logout)

		profile := engine.Group("/profile")
		{

			orders := profile.Group("/orders")
			{
				orders.GET("", orderHandler.GetOrders)
				orders.POST("/cancel", orderHandler.CancelOrder)
				//orders.POST("/return", orderHandler.ReturnOrder)

			}
		}

		// cart := engine.Group("/cart")
		// {

		// cart.PUT("/updateQuantity/plus", userHandler.UpdateQuantityAdd)
		// cart.PUT("/updateQuantity/minus", userHandler.UpdateQuantityLess)

		// }

		// checkout := engine.Group("/check-out")
		// {
		// 	// checkout.GET("", cartHandler.CheckOut)
		// 	// checkout.POST("/order", orderHandler.OrderItemsFromCart)
		// 	// checkout.GET("/order/download-invoice", orderHandler.DownloadInvoice)
		// }

	}
}
