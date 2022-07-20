package models

type Item struct {
	ItemId      uint   `gorm:"primary_key;autoIncrement"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}
