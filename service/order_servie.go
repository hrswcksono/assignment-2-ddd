package service

import (
	"hrswcksono/assignment2/dto/order_dto"
	"hrswcksono/assignment2/entity"
	"hrswcksono/assignment2/repository"
)

type OrderService interface {
	MakeOrder(orderPayload *order_dto.CreateOrderRequest) error
	// ReadOrder() ([]*order_dto.OrderHistoryResponse, error)
	// EditOrder(orderId int, itemPayload []*entity.Item) (*order_dto.OrderHistoryResponse, error)
	// RemoveOrder(orderId int)
}

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (m *orderService) MakeOrder(orderPayload *order_dto.CreateOrderRequest) error {
	if err := orderPayload.Validate(); err != nil {
		return err
	}

	movieRequest := &entity.Order{}

	orderPayload.BindToOrder(movieRequest)

	err := m.orderRepo.CreateOrder(movieRequest)

	if err != nil {
		return err
	}

	return nil
}

// func (o *orderService) ReadOrder() ([]*order_dto.OrderHistoryResponse, error) {
// 	orderHistories := []*order_dto.OrderHistoryResponse{}

// 	return orderHistories, nil
// }

// func (o *orderService) EditOrder(orderId int, itemPayload []*entity.Item) (*order_dto.OrderHistoryResponse, error) {
// 	return &order_dto.OrderHistoryResponse{}, nil
// }

// func (o *orderService) RemoveOrder(orderId int) {

// }
