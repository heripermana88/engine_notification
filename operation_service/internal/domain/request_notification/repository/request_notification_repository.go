package repository

import "gitlab.com/nusakti/golang-api-boilerplate/internal/domain/request_notification/entity"

type RequestNotificationRepository interface {
    CreateRequestNotification(request_notification *entity.RequestNotification) error
    GetRequestNotificationByID(id string) (*entity.RequestNotification, error)
    GetAllRequestNotifications() ([]*entity.RequestNotification, error)
}
