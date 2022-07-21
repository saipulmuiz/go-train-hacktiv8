package database

import (
	"fmt"
	"restapi/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
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

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Person{})

	return db
}
