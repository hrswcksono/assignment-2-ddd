package order_pg

import (
	"hrswcksono/assignment2/dto/order_dto"
	"hrswcksono/assignment2/entity"
	"hrswcksono/assignment2/repository"

	"gorm.io/gorm"
)

type orderPG struct {
	db *gorm.DB
}

func NewOrderPG(db *gorm.DB) repository.OrderRepository {
	return &orderPG{
		db: db,
	}
}

func (o *orderPG) CreateOrder(orderPayload *entity.Order) error {

	tx := o.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	orderInput := orderPayload

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(orderInput).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (o *orderPG) GetOrder() ([]*order_dto.OrderHistoryResponse, error) {

	orderHistories := []*order_dto.OrderHistoryResponse{}

	return orderHistories, nil
}

func (o *orderPG) UpdateOrder(orderId int, orderPayload *entity.Order) (*order_dto.OrderHistoryResponse, error) {
	orderHistories := &order_dto.OrderHistoryResponse{}

	return orderHistories, nil
}

func (o *orderPG) DeleteOrder(orderId int) {

}
