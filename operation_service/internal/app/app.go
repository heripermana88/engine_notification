package app

import (
	"gitlab.com/nusakti/golang-api-boilerplate/internal/infrastructure/database"
	"gitlab.com/nusakti/golang-api-boilerplate/internal/infrastructure/routes"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func NewApp() *App {
	router := mux.NewRouter()
	db := database.NewMongoDB() // Inisialisasi database

	routes.RegisterRequestNotificationRoutes(router, db) // Daftarkan routes user

	return &App{Router: router}
}
