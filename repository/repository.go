package repository

import (
	"hrswcksono/assignment2/entity"
)

type OrderRepository interface {
	CreateOrder(orderPayload *entity.Order) error
	GetAllOrder() ([]entity.Order, error)
	UpdateOrder(orderId int, orderPayload *entity.Order) (*entity.Order, error)
	DeleteOrder(orderId int) error
}
