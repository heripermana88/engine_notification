package service

import (
	"encoding/json"
	"log"

	"gitlab.com/nusakti/golang-api-boilerplate/internal/domain/request_notification/entity"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/domain/request_notification/repository"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/infrastructure/adapter"
)

type RequestNotificationService struct {
	RequestNotificationRepo repository.RequestNotificationRepository
	RabbitPublisher         *adapter.RabbitMQPublisher
}

func NewRequestNotificationService(repo repository.RequestNotificationRepository, publisher *adapter.RabbitMQPublisher) *RequestNotificationService {
	return &RequestNotificationService{
		RequestNotificationRepo: repo,
		RabbitPublisher:         publisher,
	}
}

func (s *RequestNotificationService) CreateRequestNotification(request_notification *entity.RequestNotification) error {
	err := s.RequestNotificationRepo.CreateRequestNotification(request_notification)
	if err != nil {
		return err
	}

	// Publish ke RabbitMQ
	message, err := json.Marshal(request_notification)
	if err != nil {
		log.Printf("Failed to marshal request notification: %v", err)
		return err
	}

	err = s.RabbitPublisher.Publish("request_notifications", message)
	if err != nil {
		log.Printf("Failed to publish message to RabbitMQ: %v", err)
		return err
	}

	log.Println("Request notification published to RabbitMQ")
	return nil
}

func (s *RequestNotificationService) GetRequestNotificationByID(id string) (*entity.RequestNotification, error) {
	return s.RequestNotificationRepo.GetRequestNotificationByID(id)
}

func (s *RequestNotificationService) GetAllRequestNotifications() ([]*entity.RequestNotification, error) {
	return s.RequestNotificationRepo.GetAllRequestNotifications()
}
