package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := mux.NewRouter()

	database, connectionErr := InitializeDatabase()
	if connectionErr != nil {
		log.Fatalf("Error Database connection error")
	}

	userModule, userErr := InitializeUserModule(router, database)
	if userErr != nil {
		log.Fatalf("Error User module error")
	}
	_, orderErr := InitializeOrderModule(router, database, userModule.UserService)
	if orderErr != nil {
		log.Fatalf("Error Order module error")
	}
	port := fmt.Sprintf(":%s", os.Getenv("WEB_SERVER_PORT"))
	log.Printf("Starting web server at prop %v\n", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Failed to start web server: %v\n", err)
	}
}
