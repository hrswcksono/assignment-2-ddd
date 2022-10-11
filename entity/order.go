package entity

import "time"

type Order struct {
	OrderId      uint      `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:OrderFk;references:OrderId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"items"`
}
