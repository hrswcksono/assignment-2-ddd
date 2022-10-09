package dto

import "time"

type OrderedItem struct {
	LineItemId  int
	ItemCode    string
	Description string
	Quantity    int
}

type OrderHistoryResponse struct {
	CustomerName string
	OrderedAt    time.Time
	Items        OrderedItem
}
