package service

import (
	"hrswcksono/assignment2/dto/order_dto"
	"hrswcksono/assignment2/entity"
	"hrswcksono/assignment2/repository"
)

type OrderService interface {
	MakeOrder(orderPayload *order_dto.CreateOrderRequest) error
	ReadOrder() ([]entity.Order, error)
	EditOrder(orderId int, itemPayload *order_dto.UpdateOrderRequest) (*entity.Order, error)
	RemoveOrder(orderId int) error
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

	orderRequest := &entity.Order{}

	orderPayload.BindToOrderCreate(orderRequest)

	err := m.orderRepo.CreateOrder(orderRequest)

	if err != nil {
		return err
	}

	return nil
}

func (m *orderService) ReadOrder() ([]entity.Order, error) {
	orders, err := m.orderRepo.GetAllOrder()

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o *orderService) EditOrder(orderId int, orderPayload *order_dto.UpdateOrderRequest) (*entity.Order, error) {
	if err := orderPayload.Validate(); err != nil {
		return nil, err
	}

	orderRequest := &entity.Order{}

	orderPayload.BindToOrderUpdate(orderRequest)

	// fmt.Println(orderPayload)

	editedOrder, err := o.orderRepo.UpdateOrder(orderId, orderRequest)

	if err != nil {
		return nil, err
	}

	return editedOrder, nil
}

func (m *orderService) RemoveOrder(orderId int) error {

	err := m.orderRepo.DeleteOrder(orderId)

	if err != nil {
		return err
	}

	return nil
}
