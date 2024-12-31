package handler

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "gitlab.com/nusakti/golang-api-boilerplate/internal/service"
    "gitlab.com/nusakti/golang-api-boilerplate/internal/domain/request_notification/entity"
)

type RequestNotificationHandler struct {
    RequestNotificationService *service.RequestNotificationService
}

func NewRequestNotificationHandler(request_notificationService *service.RequestNotificationService) *RequestNotificationHandler {
    return &RequestNotificationHandler{ RequestNotificationService: request_notificationService }
}

func (h *RequestNotificationHandler) CreateRequestNotification(w http.ResponseWriter, r *http.Request) {
    var request_notification entity.RequestNotification
    json.NewDecoder(r.Body).Decode(&request_notification)
    err := h.RequestNotificationService.CreateRequestNotification(&request_notification)
    if err != nil {
        http.Error(w, "Failed to create RequestNotification", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(request_notification)
}

func (h *RequestNotificationHandler) GetRequestNotificationByID(w http.ResponseWriter, r *http.Request) {
    id := mux.Vars(r)["id"]
    request_notification, err := h.RequestNotificationService.GetRequestNotificationByID(id)
    if err != nil {
        http.Error(w, "RequestNotification not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(request_notification)
}

func (h *RequestNotificationHandler) GetAllRequestNotifications(w http.ResponseWriter, r *http.Request) {
    request_notifications, err := h.RequestNotificationService.GetAllRequestNotifications()
    if err != nil {
        http.Error(w, "Failed to fetch RequestNotifications", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(request_notifications)
}
