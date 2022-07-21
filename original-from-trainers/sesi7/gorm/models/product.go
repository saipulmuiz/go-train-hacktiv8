package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"varchar(100)"`
	Brand  string `gorm:"varchar(100)"`
	UserID uint

	// User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("Before create called ...")
	if len(p.Name) <= 5 {
		err = fmt.Errorf("length of name must be greater than 5")
	}
	return
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	log.Println("After create called ...")
	return
}
