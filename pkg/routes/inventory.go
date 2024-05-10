package routes

import (
	"main/pkg/api/handler"

	"github.com/gin-gonic/gin"
)

func InventoryRoutes(engine *gin.RouterGroup, inventoryHandler *handler.InventoryHandler) {

	// engine.GET("/details", inventoryHandler.ShowIndividualProducts)
	engine.GET("/search", inventoryHandler.SearchProducts)

}
