package order_pg

import (
	"hrswcksono/assignment2/entity"
	"hrswcksono/assignment2/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (o *orderPG) GetAllOrder() ([]entity.Order, error) {
	tx := o.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var order = []entity.Order{}

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Preload("Items").Find(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return order, tx.Commit().Error
}

func (o *orderPG) UpdateOrder(orderId int, orderPayload *entity.Order) (*entity.Order, error) {
	tx := o.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var order = entity.Order{}
	orderPayload.OrderID = uint(orderId)

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Session(&gorm.Session{FullSaveAssociations: true}).Updates(orderPayload).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Preload("Items").Where("order_id = ?", orderId).Find(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return &order, tx.Commit().Error
}

func (o *orderPG) DeleteOrder(orderId int) error {
	tx := o.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	order := &entity.Order{}

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Select(clause.Associations).Delete(order, orderId).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
