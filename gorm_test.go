package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to open database: %v", err))
	}
	db.AutoMigrate(&UserGorm{})
	return db
}

func TestAddUser(t *testing.T) {
	db := setupTestDB()

	t.Run("successfully add user", func(t *testing.T) {
		// Setting ใหม่เสมอ ** Not recommended
		err := AddUser(db, "John Doe", "john.doe@gmail.com", 44)
		assert.NoError(t, err)

		var user UserGorm
		db.First(&user, "email = ?", "john.doe@gmail.com")
		assert.Equal(t, "John Doe", user.FullName)
	})

	t.Run("fail to add user with existing email", func(t *testing.T) {
		err := AddUser(db, "Jane Doe", "john.doe@gmail.com", 28)
		assert.EqualError(t, err, "email already exists")
	})
}
