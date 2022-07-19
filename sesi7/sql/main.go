package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
)

const (
	POSTGRES_HOST = "localhost"
	POSTGRES_PORT = "5432"
	POSTGRES_DB   = "hacktiv8-35"
	POSTGRES_USER = "postgres"
	POSTGRES_PASS = "root"
)

type Employee struct {
	ID       int
	FullName string
	Email    string
	Age      int
	Division string
}

func (e *Employee) print() {
	fmt.Println("ID :", e.ID)
	fmt.Println("FullName :", e.FullName)
	fmt.Println("Email :", e.Email)
	fmt.Println("Age :", e.Age)
	fmt.Println("Division :", e.Division)
}

func main() {
	db := connectDB()

	if db != nil {
		fmt.Println("db connected")
		// createEmp(db)
		getEmps(db)
	}
}

func createEmp(db *sql.DB) {
	emp := Employee{
		FullName: "Hacktiv8",
		Email:    "admin@hacktive8.com",
		Age:      20,
		Division: "IT",
	}

	query := `INSERT INTO employees (full_name, email, age, division) VALUES ($1, $2, $3, $4)`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		emp.FullName,
		emp.Email,
		emp.Age,
		emp.Division,
	)

	if err != nil {
		panic(err.Error())
	}

	log.Println("Created employee successful")
}

func getEmps(db *sql.DB) {
	var results []Employee

	query := `SELECT id, full_name, email, age, division FROM employees`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var emp Employee
		err = rows.Scan(
			&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division,
		)
		if err != nil {
			panic(err.Error())
		}
		results = append(results, emp)
	}

	for _, emp := range results {
		emp.print()
		fmt.Println(strings.Repeat("=", 20))
	}
}

func connectDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASS, POSTGRES_DB,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
