package models

type Order struct {
	OrderId      uint   `gorm:"primary_key" json:"order_id"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
	Items        []Item
}
