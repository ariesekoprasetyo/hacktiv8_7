package main

import (
	"github.com/ariesekoprasetyo/hacktiv8_7/db"
	"github.com/ariesekoprasetyo/hacktiv8_7/initializers"
	"log"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	db.SetupDB()
	err := db.DB.AutoMigrate(&db.Orders{}, &db.Items{})
	if err != nil {
		log.Fatal("Error Migrate")
	} else {
		log.Println("Success Migrate")
	}
}
