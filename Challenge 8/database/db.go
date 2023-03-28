package database

import (
	"Challenge_8/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "postgres"
	DB_PASSWORD = "Bolaitubundar123"
	DB_NAME     = "challenge7"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(models.Book{})
	return db
}
