package database

import (
	"fmt"
	"jwt/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	POSTGRES_HOST   = "localhost"
	POSTGRES_PORT   = "9999"
	POSTGRES_USER   = "hacktiv8"
	POSTGRES_PASS   = "hacktiv8"
	POSTGRES_DBNAME = "hacktiv8-35"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST,
		POSTGRES_PORT,
		POSTGRES_USER,
		POSTGRES_PASS,
		POSTGRES_DBNAME,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(models.User{}, models.Product{})

	return db
}
