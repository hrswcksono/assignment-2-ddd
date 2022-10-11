package order_dto

import "time"

type CreatedItemResponse struct {
	ItemId      int    `json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type CreatedOrderResponse struct {
	OrderId      int                   `json:"orderId"`
	CustomerName string                `json:"customerName"`
	OrderedAt    time.Time             `json:"orderedAt"`
	Items        []CreatedItemResponse `json:"items"`
}

type SuccessOrder struct {
	Message string `json:"message" example:"order sukses terbuat"`
}
