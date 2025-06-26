package server

import (
	"log"
	"net/http"
	"petopia-be/config"
	"petopia-be/middleware"
	"petopia-be/routes"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Start(cfg *config.Config, db *gorm.DB) error {
	router := mux.NewRouter()

	// Setup routes
	routes.Setup(router, db)

	// Setup CORS using middleware
	corsHandler := middleware.CORS(cfg.CORSAllowedOrigin)

	log.Printf("Server starting on port %s", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, corsHandler(router))
}
