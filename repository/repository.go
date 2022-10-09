package repository

import (
	"hrswcksono/assignment2/dto"
	"hrswcksono/assignment2/entity"
)

type OrderRepository interface {
	CreateOrder(itemPayload []*entity.Item, customerName string) error
	GetOrder() ([]*dto.OrderHistoryResponse, error)
	UpdateOrder(orderId int, itemPayload []*entity.Item) (*dto.OrderHistoryResponse, error)
	DeleteOrder(orderId int)
}
