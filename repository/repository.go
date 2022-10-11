package repository

import (
	"hrswcksono/assignment2/dto/order_dto"
	"hrswcksono/assignment2/entity"
)

type OrderRepository interface {
	CreateOrder(orderPayload *entity.Order) error
	GetAllOrder() ([]entity.Order, error)
	UpdateOrder(orderId int, orderPayload *entity.Order) (*order_dto.OrderHistoryResponse, error)
	DeleteOrder(orderId int) error
}
