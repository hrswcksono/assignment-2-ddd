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

// CreateOrder godoc
// @Tags order
// @Description Create New Order Data
// @ID create-new-order
// @Accept json
// @Produce json
// @Param RequestBody body order_dto.CreateOrderRequest true "request body json"
// @Success 201 {object} order_dto.MessageOrder
// @Router /orders [post]
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

	c.JSON(http.StatusCreated, order_dto.MessageOrder{
		Message: "order sukses terbuat",
	})
}

// GetOrder godoc
// @Tags order
// @Description Retrieving Order Data
// @ID get-all-order
// @Produce json
// @Success 200 {array} entity.Order
// @Router /orders [get]
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

// UpdateOrder godoc
// @Tags order
// @Description Update Order Data
// @ID update-order
// @Produce json
// @Accept json
// @Param orderId path int true "order's order_id"
// @Param RequestBody body order_dto.UpdateOrderRequest true "request body json"
// @Success 200 {object} entity.Order
// @Router /orders/{orderId} [put]
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

// DeleteOrder godoc
// @Tags order
// @Description Delete Order Data
// @ID delete-order
// @Produce json
// @Param orderId path int true "order's order_id"
// @Success 200 {object} order_dto.MessageOrder
// @Router /orders/{orderId} [delete]
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

	c.JSON(http.StatusCreated, order_dto.MessageOrder{
		Message: "order sukses terhapus",
	})
}
