package entity

import "time"

type Order struct {
	OrderID      uint      `gorm:"primaryKey" json:"orderId" example:"1"`
	CustomerName string    `json:"customerName" example:"haris"`
	OrderedAt    time.Time `json:"orderedAt" example:"2022-10-07T15:54:24.575005+07:00"`
	Items        []Item    `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"items"`
}
