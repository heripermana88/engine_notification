package routes

import (
	"github.com/gorilla/mux"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/handler"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/infrastructure/adapter"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/repository"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRequestNotificationRoutes(router *mux.Router, db *mongo.Database, rabbitPublisher *adapter.RabbitMQPublisher) {
	request_notificationRepo := repository.NewRequestNotificationRepository(db)
	request_notificationService := service.NewRequestNotificationService(request_notificationRepo, rabbitPublisher)
	request_notificationHandler := handler.NewRequestNotificationHandler(request_notificationService)

	router.HandleFunc("/request_notifications", request_notificationHandler.CreateRequestNotification).Methods("POST")
	router.HandleFunc("/request_notifications/{id}", request_notificationHandler.GetRequestNotificationByID).Methods("GET")
	router.HandleFunc("/request_notifications", request_notificationHandler.GetAllRequestNotifications).Methods("GET")
}
