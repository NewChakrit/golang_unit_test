package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "mydatabase" // as defined in docker-compose.yml
)

type UserPostgres struct {
	gorm.Model
	Fullname string
	Email    string `gorm:"unique"`
	Age      int
}

// InitializeDBPostgres initializes the database and automigrates the UserPostgres model.
func InitializeDBPostgres() *gorm.DB {
	// Configure your PostgreSQL database details here
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&UserPostgres{})
	return db
}

// AddUser adds a new user to the database.
func AddUserPostgres(db *gorm.DB, fullname, email string, age int) error {
	user := UserPostgres{Fullname: fullname, Email: email, Age: age}

	// Check if email already exists
	var count int64
	db.Model(&UserPostgres{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		return errors.New("email already exists")
	}

	// Save the new user
	result := db.Create(&user)
	return result.Error
}

func Postgres() {
	db := InitializeDBPostgres()

	AddUserPostgres(db, "John Doe", "jane.doe@example.com", 30)
}
