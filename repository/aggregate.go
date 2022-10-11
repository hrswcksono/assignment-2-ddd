package repository

import (
	// "hrswcksono/assignment2/dto/order_dto"
	"hrswcksono/assignment2/entity"
)

type OrderHistory struct {
	Item  entity.Item
	Order entity.Order
}

// func (o OrderHistory) ToOrderHistoryDTO() order_dto.OrderHistoryResponse {
// 	return order_dto.OrderHistoryResponse{
// 		CustomerName: o.Order.CustomerName,
// 		OrderedAt:    o.Order.OrderedAt,
// 		Items:        []order_dto.OrderedItem{},
// 	}
// }
