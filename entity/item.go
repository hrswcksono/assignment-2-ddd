package entity

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"lineItemId" example:"1"`
	ItemCode    string `json:"itemCode" example:"12345"`
	Description string `json:"description" example:"barang antik"`
	Quantity    int    `json:"quantity" example:"45"`
	OrderID     uint   `json:"orderId" example:"1"`
}
