package repository

import (
	"hrswcksono/assignment2/dto"
	"hrswcksono/assignment2/entity"
)

type OrderHistory struct {
	Item  entity.Item
	Order entity.Order
}

func (o OrderHistory) ToOrderHistoryDTO() dto.OrderHistoryResponse {
	return dto.OrderHistoryResponse{
		CustomerName: o.Order.CustomerName,
		OrderedAt:    o.Order.OrderedAt,
		Items: dto.OrderedItem{
			LineItemId:  o.Item.Id,
			ItemCode:    o.Item.ItemCode,
			Description: o.Item.Description,
			Quantity:    o.Item.Quantity,
		},
	}
}
