package main

import (
	"fmt"
	"gorm/db"
	"gorm/models"
	"log"

	"gorm.io/gorm"
)

func main() {
	db := db.ConnectDB()
	if db != nil {
		log.Println("db connected")
	}

	// createUser("admin@hacktiv8.com", db)
	// getUsers(db)

	createProduct(db)
	// getOneUserWithProducts(db)
	// getProductWithUser(db)
}

func createUser(email string, db *gorm.DB) {
	user := models.User{
		Email: email,
	}

	err := db.Create(&user).Error
	if err != nil {
		panic(err.Error())
	}

	log.Println("created", email, "success")

}

func getUsers(db *gorm.DB) {
	var users []models.User

	err := db.Find(&users, "email=?", "admin@hacktiv8.com").Error
	if err != nil {
		panic(err.Error())
	}

	for _, user := range users {
		fmt.Println("email :", user.Email)
	}
}

func getOneUserWithProducts(db *gorm.DB) {
	var user models.User

	err := db.Preload("Products").First(&user, "id=?", 5).Error
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v \n", user)
}

func getProductWithUser(db *gorm.DB) {
	var product models.Product

	err := db.Joins("User").First(&product, "id=?", 1).Error
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(product)
}

func createProduct(db *gorm.DB) {
	product := models.Product{
		Name:   "Laptop",
		Brand:  "Asus",
		UserID: 5,
	}

	err := db.Create(&product).Error
	if err != nil {
		panic(err.Error())
	}

	log.Println("create product success")
}
