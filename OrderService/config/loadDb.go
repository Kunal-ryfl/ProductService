package config

import (
	"OrderService/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Db() *gorm.DB {
	dsn := "host=localhost user=kunal password=abc dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	err = db.AutoMigrate(&model.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migrated successfully")
	return db
}
