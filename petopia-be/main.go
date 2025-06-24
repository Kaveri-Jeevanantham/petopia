package main

import (
	"log"
	"net/http"
	"petopia-be/controller"
	"petopia-be/db"
	"petopia-be/messaging/rabbitmq"
	"petopia-be/seed"

	"github.com/rs/cors"

	"os"
	_ "petopia-be/docs"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Petopia API
// @version 1.0
// @description This is a sample server for Petopia.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize RabbitMQ
	if err := rabbitmq.InitRabbitMQ(); err != nil {
		log.Fatalf("Could not initialize RabbitMQ: %v", err)
	}
	defer rabbitmq.CloseConnection()

	// Declare test queue
	if err := rabbitmq.DeclareTestQueue(); err != nil {
		log.Fatalf("Could not declare test queue: %v", err)
	}

	// Start test consumer
	if err := rabbitmq.StartTestConsumer(); err != nil {
		log.Fatalf("Could not start test consumer: %v", err)
	}

	// Initialize database connection
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Seed the database
	if err := seed.SeedDatabase(database); err != nil {
		log.Fatalf("Error seeding the database: %v", err)
	}

	controller.InitDB(database)

	// Set up CORS
	corsAllowedOrigin := os.Getenv("CORS_ALLOWED_ORIGIN")
	c := cors.New(cors.Options{
		AllowedOrigins: []string{corsAllowedOrigin},
	})

	// Set up router
	router := mux.NewRouter()

	// Health check endpoint
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy"}`))
	}).Methods("GET")

	// Define routes
	router.HandleFunc("/api/products", controller.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products", controller.ListProducts).Methods("GET")
	router.HandleFunc("/api/products/{id:[0-9]+}", controller.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id:[0-9]+}", controller.DeleteProduct).Methods("DELETE")

	router.HandleFunc("/api/test/message", controller.SendTestMessage).Methods("POST")

	// Swagger route
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Start server with CORS
	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
	// log.Fatal(http.ListenAndServe(":8080", router))
}
