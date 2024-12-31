package routes

import (
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/mongo"
    "gitlab.com/nusakti/golang-api-boilerplate/internal/handler"
    "gitlab.com/nusakti/golang-api-boilerplate/internal/repository"
    "gitlab.com/nusakti/golang-api-boilerplate/internal/service"
)

func RegisterRequestNotificationRoutes(router *mux.Router, db *mongo.Database) {
    request_notificationRepo := repository.NewRequestNotificationRepository(db)
    request_notificationService := service.NewRequestNotificationService(request_notificationRepo)
    request_notificationHandler := handler.NewRequestNotificationHandler(request_notificationService)

    router.HandleFunc("/request_notifications", request_notificationHandler.CreateRequestNotification).Methods("POST")
    router.HandleFunc("/request_notifications/{id}", request_notificationHandler.GetRequestNotificationByID).Methods("GET")
    router.HandleFunc("/request_notifications", request_notificationHandler.GetAllRequestNotifications).Methods("GET")
}
