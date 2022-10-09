package service

import (
	"hrswcksono/assignment2/dto"
	"hrswcksono/assignment2/entity"
	"hrswcksono/assignment2/repository"
)

type OrderService interface {
	MakeOrder(itemPayload []*entity.Item, customerName string) error
	ReadOrder() ([]*dto.OrderHistoryResponse, error)
	EditOrder(orderId int, itemPayload []*entity.Item) (*dto.OrderHistoryResponse, error)
	RemoveOrder(orderId int)
}

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) MakeOrder(itemPayload []*entity.Item, customerName string) error {
	return nil
}

func (o *orderService) ReadOrder() ([]*dto.OrderHistoryResponse, error) {
	orderHistories := []*dto.OrderHistoryResponse{}

	return orderHistories, nil
}

func (o *orderService) EditOrder(orderId int, itemPayload []*entity.Item) (*dto.OrderHistoryResponse, error) {
	return &dto.OrderHistoryResponse{}, nil
}

func (o *orderService) RemoveOrder(orderId int) {

}
