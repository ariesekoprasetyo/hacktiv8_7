package main

import (
	"fmt"
	"github.com/ariesekoprasetyo/hacktiv8_7/initializers"
	"github.com/ariesekoprasetyo/hacktiv8_7/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func init() {
	initializers.LoadEnv()
	setupDB()
	autoMigrate()
}

func main() {

}
func setupDB() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed To Connect")
		return
	}
	fmt.Println("Database is Connected")
}

func autoMigrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Address{})
	if err != nil {
		log.Fatal("Error Migrate")
	} else {
		fmt.Println("Success Migrate")
	}
}
