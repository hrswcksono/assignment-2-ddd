package rest

import (
	"fmt"
	"hrswcksono/assignment2/database"
	"hrswcksono/assignment2/docs"
	"hrswcksono/assignment2/repository/order_pg"
	"hrswcksono/assignment2/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

const port = ":8080"

func StartApp() {
	database.StartDB()

	db := database.GetDB()

	orderRepo := order_pg.NewOrderPG(db)
	orderService := service.NewOrderService(orderRepo)
	orderRestHandler := newOrderRestHandler(orderService)

	route := gin.Default()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Order API"
	docs.SwaggerInfo.Description = "Order API with DDD pattern"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	// docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	orderRoute := route.Group("/orders")
	{
		orderRoute.POST("/", orderRestHandler.MakeOrder)
		orderRoute.GET("/", orderRestHandler.ReadOrder)
		orderRoute.DELETE("/:orderId", orderRestHandler.RemoveOrder)
		orderRoute.PUT("/:orderId", orderRestHandler.EditOrder)
	}

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)
}
