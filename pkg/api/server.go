package http

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// _ "main/cmd/api/docs"
	handler "main/pkg/api/handler"
	"main/pkg/routes"
)

// ServerHTTP represents an HTTP server for the web application.
type ServerHTTP struct {
	engine *gin.Engine // engine is the core of the Gin web framework, responsible for routing HTTP requests and handling middleware.
}

/*
NewServerHTTP creates a new instance of ServerHTTP.

Parameters:

- inventoryHandler: A handler for inventory-related operations.
- userHandler: A handler for user-related operations.

- adminHandler: A handler for admin-related operations.

- orderHandler: A handler for order-related operations.

Returns:
- *ServerHTTP: A pointer to the newly created ServerHTTP instance.
*/
func NewServerHTTP(inventoryHandler *handler.InventoryHandler, userHandler *handler.UserHandler, adminHandler *handler.AdminHandler, orderHandler *handler.OrderHandler) *ServerHTTP {
	engine := gin.New()
	engine.Use(gin.Logger())

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.UserRoutes(engine.Group("/users"), userHandler, orderHandler, inventoryHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, inventoryHandler)

	return &ServerHTTP{engine: engine}
}

/*
Start starts the HTTP server and listens on port 1243.
*/
func (sh *ServerHTTP) Start() {
	sh.engine.Run(":1233")

}
