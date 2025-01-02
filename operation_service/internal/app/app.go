package app

import (
	"log"

	"gitlab.com/nusakti/golang-api-boilerplate/internal/infrastructure/adapter"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/infrastructure/database"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/infrastructure/routes"

	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

type App struct {
	Router            *mux.Router
	RabbitMQListener  *adapter.RabbitMQListener
	RabbitMQPublisher *adapter.RabbitMQPublisher
}

func NewApp(rabbitConn *amqp.Connection, exchange, routingKey string, callback func([]byte) error) *App {
	router := mux.NewRouter()
	db := database.NewMongoDB() // Initialize database

	// Set up RabbitMQ Publisher
	rabbitPublisher, err := adapter.NewRabbitMQPublisher(rabbitConn, exchange)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ Publisher: %v", err)
	}

	// Set up RabbitMQ Listener
	rabbitListener, err := adapter.NewRabbitMQListener(rabbitConn, exchange, routingKey, callback)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ Listener: %v", err)
	}

	// Start listening for RabbitMQ messages
	go func() {
		if err := rabbitListener.StartListening(); err != nil {
			log.Fatalf("Failed to start RabbitMQ Listener: %v", err)
		}
	}()

	routes.RegisterRequestNotificationRoutes(router, db, rabbitPublisher) // Register routes

	return &App{
		Router:            router,
		RabbitMQListener:  rabbitListener,
		RabbitMQPublisher: rabbitPublisher,
	}
}
