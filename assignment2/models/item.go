package models

type Item struct {
	ItemId      uint   `gorm:"primaryKey;autoIncrement"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}
