package order_dto

import "time"

type OrderedItem struct {
	LineItemId  int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type OrderHistoryResponse struct {
	OrderId      int           `json:"orderId"`
	CustomerName string        `json:"customerName"`
	OrderedAt    time.Time     `json:"orderedAt"`
	Items        []OrderedItem `json:"items"`
}

type SuccessOrder struct {
	Message string `json:"message" example:"order sukses terbuat"`
}
