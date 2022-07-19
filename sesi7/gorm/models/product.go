package models

import "time"

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"varchar(100)"`
	Brand     string `gorm:"varchar(100)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
