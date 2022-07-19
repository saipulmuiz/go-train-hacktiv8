package models

type Order struct {
	OrderId      uint   `gorm:"primaryKey;autoIncrement"`
	CustomerName string `json:"customerName"`
	OrderedAt    string `json:"orderedAt"`
	Items        []Item
}
