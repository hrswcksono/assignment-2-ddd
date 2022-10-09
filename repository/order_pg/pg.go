package order_pg

import (
	"hrswcksono/assignment2/dto"
	"hrswcksono/assignment2/entity"
	"hrswcksono/assignment2/repository"

	"gorm.io/gorm"
)

// const (
// 	createOrderItemQuery = `
// 		insert into orders
// 		(

// 		)
// 	`
// )

type orderPG struct {
	db *gorm.DB
}

func NewOrderPG(db *gorm.DB) repository.OrderRepository {
	return &orderPG{
		db: db,
	}
}

func (o *orderPG) CreateOrder(itemPayload []*entity.Item, customerName string) error {

	return nil
}

func (o *orderPG) GetOrder() ([]*dto.OrderHistoryResponse, error) {

	orderHistories := []*dto.OrderHistoryResponse{}

	return orderHistories, nil
}

func (o *orderPG) UpdateOrder(orderId int, itemPayload []*entity.Item) (*dto.OrderHistoryResponse, error) {
	orderHistories := &dto.OrderHistoryResponse{}

	return orderHistories, nil
}

func (o *orderPG) DeleteOrder(orderId int) {

}
