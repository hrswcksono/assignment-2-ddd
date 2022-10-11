package rest

import (
	"hrswcksono/assignment2/dto/order_dto"
	"hrswcksono/assignment2/pkg/helper"
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

func (o orderRestHandler) ReadOrder(c *gin.Context) {
	orders, err := o.service.ReadOrder()

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (o orderRestHandler) EditOrder(c *gin.Context) {
	orderId, err := helper.GetParamId(c, "orderId")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	var updateRequest *order_dto.UpdateOrderRequest

	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	order, err := o.service.EditOrder(orderId, updateRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (o orderRestHandler) RemoveOrder(c *gin.Context) {

	orderId, err := helper.GetParamId(c, "orderId")

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err = o.service.RemoveOrder(orderId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, order_dto.SuccessOrder{
		Message: "order sukses terhapus",
	})
}
