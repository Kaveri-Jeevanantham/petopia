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

// StartV2 starts the HTTP server using clean architecture
func StartV2(cfg *config.Config, db *gorm.DB) error {
	router := mux.NewRouter()

	// Add CORS middleware
	corsHandler := middleware.CORS(cfg.CORSAllowedOrigin)

	// Setup routes using clean architecture approach
	routes.SetupV2(router, db)

	// Start server
	log.Printf("Starting server on port %s", cfg.Port)
	return http.ListenAndServe(":"+cfg.Port, corsHandler(router))
}
