package rest

import (
	"fmt"
	"hrswcksono/assignment2/database"
	"hrswcksono/assignment2/repository/order_pg"
	"hrswcksono/assignment2/service"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

func StartApp() {
	database.StartDB()

	db := database.GetDB()

	orderRepo := order_pg.NewOrderPG(db)
	orderService := service.NewOrderService(orderRepo)
	orderRestHandler := newOrderRestHandler(orderService)

	route := gin.Default()

	orderRoute := route.Group("/order")
	{
		orderRoute.POST("/", orderRestHandler.MakeOrder)
	}

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)
}
