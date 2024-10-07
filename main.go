package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"orderbook-system/src/modules/users"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	_, connectionError := InitializeDBConnection()
	if connectionError != nil {
		log.Fatalf("Db connection error")
	}

	users.InitController()

	port := fmt.Sprintf(":%s", os.Getenv("WEB_SERVER_PORT"))

	log.Printf("Starting web server at prop %v\n", port)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Failed to start web server: %v\n", err)
	}
}
