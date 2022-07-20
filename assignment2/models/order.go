package models

type Order struct {
	OrderId      uint    `gorm:"primary_key" json:"orderId"`
	CustomerName string  `json:"customerName"`
	OrderedAt    string  `json:"orderedAt"`
	Items        []*Item `gorm:"foreignkey:OrderId;association_foreignkey:OrderId"`
}
