package order_dto

import (
	"hrswcksono/assignment2/entity"
	"time"

	"github.com/asaskevich/govalidator"
)

type UpdateOrderRequest struct {
	CustomerName string              `json:"customerName" valid:"required~customer_name cannot be empty"`
	OrderedAt    time.Time           `json:"orderedAt" valid:"required~ordered_at cannot be empty"`
	Items        []UpdateItemRequest `json:"items" valid:"required~items cannot be empty"`
}

type UpdateItemRequest struct {
	LineItemId  uint   `json:"lineItemId" valid:"required~lineItemId cannot be empty"`
	ItemCode    string `json:"itemCode" valid:"required~itemCode cannot be empty"`
	Description string `json:"description" valid:"required~description cannot be empty"`
	Quantity    int    `json:"quantity" valid:"required~quantity cannot be empty"`
	// OrderFk     uint   `json:"orderFk" valid:"required~orderFk cannot be empty"`
}

func (from UpdateOrderRequest) BindToOrderUpdate(to *entity.Order) {
	to.CustomerName = from.CustomerName
	to.OrderedAt = from.OrderedAt
	to.Items = DataMapperUpdate(from.Items)
}

func (from UpdateItemRequest) BindToItemUpdate(to *entity.Item) {
	to.ItemID = from.LineItemId
	to.ItemCode = from.ItemCode
	to.Description = from.Description
	to.Quantity = from.Quantity
	// to.OrderFk = from.OrderFk
}

func DataMapperUpdate(input []UpdateItemRequest) []entity.Item {
	var output []entity.Item
	for _, index := range input {
		var newItem entity.Item
		index.BindToItemUpdate(&newItem)
		output = append(output, newItem)
	}
	return output
}

func (m *UpdateOrderRequest) Validate() error {
	_, err := govalidator.ValidateStruct(m)

	if err != nil {
		return err
	}

	return nil
}
