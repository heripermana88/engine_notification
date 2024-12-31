package repository

import (
	"context"

	"gitlab.com/nusakti/golang-api-boilerplate/internal/domain/request_notification/entity"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/domain/request_notification/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestNotificationRepositoryImpl struct {
	db *mongo.Collection
}

func NewRequestNotificationRepository(db *mongo.Database) repository.RequestNotificationRepository {
	return &RequestNotificationRepositoryImpl{
		db: db.Collection("request_notifications"),
	}
}

func (r *RequestNotificationRepositoryImpl) CreateRequestNotification(request_notification *entity.RequestNotification) error {
	request_notification.Status = "pending"
	_, err := r.db.InsertOne(context.Background(), request_notification)
	return err
}

func (r *RequestNotificationRepositoryImpl) GetRequestNotificationByID(id string) (*entity.RequestNotification, error) {
	var request_notification entity.RequestNotification
	err := r.db.FindOne(context.Background(), bson.M{"_id": id}).Decode(&request_notification)
	return &request_notification, err
}

func (r *RequestNotificationRepositoryImpl) GetAllRequestNotifications() ([]*entity.RequestNotification, error) {
	var request_notifications []*entity.RequestNotification
	cursor, err := r.db.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.Background()) {
		var request_notification entity.RequestNotification
		cursor.Decode(&request_notification)
		request_notifications = append(request_notifications, &request_notification)
	}
	return request_notifications, nil
}
