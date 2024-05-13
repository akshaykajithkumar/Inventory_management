package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "main/pkg/config"
	domain "main/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Inventory{})

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Admin{})

	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.OrderItem{})

	// Set up database triggers
	if err := SetUpDBTriggers(db); err != nil {
		return nil, err
	}
	// //setup the indexes
	if err := CreateIndexes(db); err != nil {
		log.Printf("failed to setup database indexes")
	}
	return db, dbErr
}
