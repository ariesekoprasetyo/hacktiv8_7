package main

import (
	"fmt"
	"github.com/ariesekoprasetyo/hacktiv8_7/db"
	"github.com/ariesekoprasetyo/hacktiv8_7/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db.SetupDB()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router.InitializeRouter(),
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
