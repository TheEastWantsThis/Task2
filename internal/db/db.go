package db

import (
	"awesomeProject1/internal/servesTask"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не работает БД: %v", err)
	}
	if err := db.AutoMigrate(&servesTask.TaskNew{}); err != nil {
		log.Fatalf("Не работает миграция: %v", err)
	}
	return db, nil
}
