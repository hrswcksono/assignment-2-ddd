package rest

import (
	"hrswcksono/assignment2/dto/order_dto"
	"hrswcksono/assignment2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type orderRestHandler struct {
	service service.OrderService
}

func newOrderRestHandler(service service.OrderService) orderRestHandler {
	return orderRestHandler{
		service: service,
	}
}

func (o orderRestHandler) MakeOrder(c *gin.Context) {
	var orderRequest *order_dto.CreateOrderRequest

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	err := o.service.MakeOrder(orderRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, order_dto.SuccessOrder{
		Message: "order sukses terbuat",
	})
}
