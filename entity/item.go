package entity

type Item struct {
	ItemId      uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderFk     uint   `json:"orderFk"`
}
