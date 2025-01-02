package main

import (
	"log"
	"net/http"
	"os"

	"gitlab.com/nusakti/golang-api-boilerplate/internal/app"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get RabbitMQ URL from environment
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		log.Fatal("RABBITMQ_URL is not set in .env")
	}

	// Connect to RabbitMQ
	rabbitConn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitConn.Close()

	// Define a simple callback function for processing messages
	callback := func(message []byte) error {
		// Logic to process the message
		log.Printf("Received message main: %s", message)
		return nil
	}

	// Initialize the app with RabbitMQ connection and listener
	app := app.NewApp(rabbitConn, "request_notifications", callback)

	log.Println("Server is running on port 8123")
	log.Fatal(http.ListenAndServe(":8123", app.Router))
}
