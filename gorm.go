package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserGorm struct {
	gorm.Model
	FullName string
	Email    string `gorm:"unique"`
	Age      int
}

// InitialzeDB initializes the database and automigrates the User
func InitializeDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&UserGorm{})
	return db
}

func AddUser(db *gorm.DB, fullname, email string, age int) error {
	user := UserGorm{FullName: fullname, Email: email, Age: age}

	// Check if email already exists
	var count int64
	db.Model(&UserGorm{}).Where("email = ?", email).Count(&count) // count email from db
	if count > 0 {
		return errors.New("email already exists")
	}

	// Save the new user
	result := db.Create(&user)
	return result.Error
}

func Gorm() {
	db := InitializeDB()
	// Your application code

	err := AddUser(db, "John Doe", "janedoe@gmail.com", 44)
	if err != nil {
		fmt.Println("err DB:", err)
	}
}
