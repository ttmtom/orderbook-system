package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database, connectionErr := InitializeDatabase()
	if connectionErr != nil {
		log.Fatalf("Error Database connection error")
	}

	usersModule, userErr := InitializeUserModule(database)
	if userErr != nil {
		log.Fatalf("Error User module error")
	}

	http.Handle("/user", usersModule.Router())

	port := fmt.Sprintf(":%s", os.Getenv("WEB_SERVER_PORT"))

	log.Printf("Starting web server at prop %v\n", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Failed to start web server: %v\n", err)
	}
}
