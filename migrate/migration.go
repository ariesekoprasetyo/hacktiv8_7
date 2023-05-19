package main

import (
	"github.com/ariesekoprasetyo/hacktiv8_7/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db.SetupDB()
	err = db.DB.AutoMigrate(&db.Orders{}, &db.Items{})
	if err != nil {
		log.Fatal("Error Migrate")
	} else {
		log.Println("Success Migrate")
	}
}
