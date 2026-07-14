package db

import (
	"GOLANG-AUTH-SYSTEM/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // keep connection of db we creating

func ConnectDB() { // connnect db
	err := godotenv.Load() // read env file

	if err != nil {
		log.Fatal("Error loading .env")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",

		// sprintf is for creating string type response

		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB = db

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("Not migrated yet", err)
	}

	fmt.Println("Database Connected Successfully")
}
