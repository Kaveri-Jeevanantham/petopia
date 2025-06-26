package main

import (
	"log"
	"petopia-be/config"
	"petopia-be/db"
	"petopia-be/server"
	"petopia-be/messaging/rabbitmq"
	"petopia-be/seed"

	_ "petopia-be/docs"

	"github.com/joho/godotenv"
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
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load configuration
	cfg := config.Load()

	// Initialize RabbitMQ
	if err := rabbitmq.InitRabbitMQ(); err != nil {
		log.Fatalf("Could not initialize RabbitMQ: %v", err)
	}
	defer rabbitmq.CloseConnection()

	if err := rabbitmq.DeclareTestQueue(); err != nil {
		log.Fatalf("Could not declare test queue: %v", err)
	}

	if err := rabbitmq.StartTestConsumer(); err != nil {
		log.Fatalf("Could not start test consumer: %v", err)
	}

	// Initialize database
	database, err := db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Seed the database
	if err := seed.SeedDatabase(database); err != nil {
		log.Fatalf("Error seeding the database: %v", err)
	}

	// Start server
	log.Fatal(server.Start(cfg, database))
}
