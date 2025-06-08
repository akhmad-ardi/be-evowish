// tests/testdb.go
package tests

import (
	"log"

	"github.com/akhmad-ardi/be-evowish/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to in-memory SQLite: %v", err)
	}

	// Auto migrate semua model yang dibutuhkan untuk test
	db.AutoMigrate(&models.User{})

	return db
}
