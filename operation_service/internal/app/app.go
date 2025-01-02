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
	RabbitMQListener  *adapter.RabbitMQListener // Menambahkan listener RabbitMQ
	RabbitMQPublisher *adapter.RabbitMQPublisher
}

func NewApp(rabbitConn *amqp.Connection, queueName string, callback func([]byte) error) *App {
	router := mux.NewRouter()
	db := database.NewMongoDB() // Inisialisasi database

	// Menyiapkan Publisher RabbitMQ
	rabbitPublisher, err := adapter.NewRabbitMQPublisher(rabbitConn)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ Publisher: %v", err)
	}

	// Menyiapkan Listener RabbitMQ
	rabbitListener, err := adapter.NewRabbitMQListener(rabbitConn, queueName, callback)
	if err != nil {
		log.Fatalf("Failed to initialize RabbitMQ Listener: %v", err)
	}

	// Mulai mendengarkan pesan dari RabbitMQ
	go func() {
		if err := rabbitListener.StartListening(); err != nil {
			log.Fatalf("Failed to start RabbitMQ Listener: %v", err)
		}
	}()

	routes.RegisterRequestNotificationRoutes(router, db, rabbitPublisher) // Daftarkan routes user

	return &App{
		Router:            router,
		RabbitMQListener:  rabbitListener,
		RabbitMQPublisher: rabbitPublisher,
	}
}
