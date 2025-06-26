package routes

import (
	"net/http"
	"petopia-be/controller"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

func Setup(router *mux.Router, db *gorm.DB) {
	controller.InitDB(db)

	// Health check
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	}).Methods("GET")

	// Product routes
	router.HandleFunc("/api/products", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products", controller.ListProducts).Methods("GET")
	router.HandleFunc("/api/products/{id:[0-9]+}", controller.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id:[0-9]+}", controller.DeleteProduct).Methods("DELETE")

	// Test routes
	router.HandleFunc("/api/test/message", controller.SendTestMessage).Methods("POST")

	// Swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
