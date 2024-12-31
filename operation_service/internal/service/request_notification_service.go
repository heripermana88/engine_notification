package service

import (
    "gitlab.com/nusakti/golang-api-boilerplate/internal/domain/request_notification/entity"
    "gitlab.com/nusakti/golang-api-boilerplate/internal/domain/request_notification/repository"
)

type RequestNotificationService struct {
    RequestNotificationRepo repository.RequestNotificationRepository
}

func NewRequestNotificationService(repo repository.RequestNotificationRepository) *RequestNotificationService {
    return &RequestNotificationService{
        RequestNotificationRepo: repo,
    }
}

func (s *RequestNotificationService) CreateRequestNotification(request_notification *entity.RequestNotification) error {
    return s.RequestNotificationRepo.CreateRequestNotification(request_notification)
}

func (s *RequestNotificationService) GetRequestNotificationByID(id string) (*entity.RequestNotification, error) {
    return s.RequestNotificationRepo.GetRequestNotificationByID(id)
}

func (s *RequestNotificationService) GetAllRequestNotifications() ([]*entity.RequestNotification, error) {
    return s.RequestNotificationRepo.GetAllRequestNotifications()
}
