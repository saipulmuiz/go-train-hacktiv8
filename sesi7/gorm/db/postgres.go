package db

import (
	"fmt"
	"gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	POSTGRES_HOST = "localhost"
	POSTGRES_PORT = "5432"
	POSTGRES_DB   = "hacktiv8-35"
	POSTGRES_USER = "postgres"
	POSTGRES_PASS = "root"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASS, POSTGRES_DB,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})

	return db
}
