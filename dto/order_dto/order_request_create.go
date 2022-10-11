package order_dto

import (
	"hrswcksono/assignment2/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateOrderRequest struct {
	CustomerName string              `json:"customerName" valid:"required~customerName cannot be empty"`
	OrderedAt    time.Time           `json:"orderedAt" valid:"required~orderedAt cannot be empty"`
	Items        []CreateItemRequest `json:"items" valid:"required~items cannot be empty"`
}

type CreateItemRequest struct {
	ItemCode    string `json:"itemCode" valid:"required~itemCode cannot be empty"`
	Description string `json:"description" valid:"required~description cannot be empty"`
	Quantity    int    `json:"quantity" valid:"required~quantity cannot be empty"`
}

func (from CreateOrderRequest) BindToOrderCreate(to *entity.Order) {
	to.CustomerName = from.CustomerName
	to.OrderedAt = from.OrderedAt
	to.Items = DataMapperCreate(from.Items)
}

func (from CreateItemRequest) BindToItemCreate(to *entity.Item) {
	to.ItemCode = from.ItemCode
	to.Description = from.Description
	to.Quantity = from.Quantity
}

func DataMapperCreate(input []CreateItemRequest) []entity.Item {
	var output []entity.Item
	for _, index := range input {
		var newItem entity.Item
		index.BindToItemCreate(&newItem)
		output = append(output, newItem)
	}
	return output
}

func (m *CreateOrderRequest) Validate() error {
	_, err := govalidator.ValidateStruct(m)

	if err != nil {
		return err
	}

	return nil
}
