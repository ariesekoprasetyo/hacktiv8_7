package main

import (
	"fmt"
	"github.com/ariesekoprasetyo/hacktiv8_7/db"
	"github.com/ariesekoprasetyo/hacktiv8_7/initializers"
	"github.com/ariesekoprasetyo/hacktiv8_7/router"
	"log"
	"net/http"
	"os"
)

func init() {
	initializers.LoadEnv()
	db.SetupDB()
}

func main() {
	s := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: router.InitializeRouter(),
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
