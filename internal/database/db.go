package database

import (
	"log"

	"github.com/Jordy-6/CRM-Go/internal/storage"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbName := "CRM.db"

	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database! '%s' : %v", dbName, err)
	}

	DB = database

	err = database.AutoMigrate(&storage.Contact{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connection successful")
}
