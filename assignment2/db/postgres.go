package db

import (
	"assignment2/models"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	POSTGRES_HOST = "localhost"
	POSTGRES_PORT = "5432"
	POSTGRES_DB   = "db_assignment2"
	POSTGRES_USER = "postgres"
	POSTGRES_PASS = "root"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST,
		POSTGRES_PORT,
		POSTGRES_USER,
		POSTGRES_PASS,
		POSTGRES_DB,
	)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})

	return db
}
