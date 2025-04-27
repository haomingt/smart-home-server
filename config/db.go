package config

import (
	"fmt"
	"log"
	"smart-home-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbConfig := AppConfig.Database

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// 自动迁移
	err = DB.AutoMigrate(
		&models.User{},
		&models.Pet{}, 
		&models.FeedingPlan{}, 
		&models.PetFoodInventory{}, 
		&models.FeedingRecord{}, 
		&models.MedicationReminder{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	fmt.Println("Database connection established and models migrated.")
}
