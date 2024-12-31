package main

import (
	"log"
	"net/http"

	"gitlab.com/nusakti/golang-api-boilerplate/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := app.NewApp()
	log.Println("Server is running on port 8123")
	log.Fatal(http.ListenAndServe(":8123", app.Router))
}
