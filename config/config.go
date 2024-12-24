package config

import (
    "log"
	"fmt"
	"os"

    "github.com/TobiAdeniji94/ecommerce_api/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// get environment variables
	host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

	// data source name
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		host, user, password, dbName, port,
	)

	// connect to db
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // db instance to global DB
    DB = database

    // migrate all models
    err = DB.AutoMigrate(
        &models.User{},
        &models.Product{},
        &models.Order{},
        &models.OrderItem{},
    )
    if err != nil {
        log.Fatalf("Failed to auto-migrate: %v", err)
    }

    log.Println("Database connected and migrated successfully!")
}
