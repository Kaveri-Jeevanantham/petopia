package rabbitmq

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel

func InitRabbitMQ() error {
	// Get connection settings from environment variables
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	vhost := os.Getenv("RABBITMQ_VHOST")

	// Create connection URL
	url := fmt.Sprintf("amqp://%s:%s@%s:%s%s", user, password, host, port, vhost)

	// Connect to RabbitMQ
	var err error
	conn, err = amqp.Dial(url)
	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	// Create channel
	channel, err = conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %v", err)
	}

	log.Println("Successfully connected to RabbitMQ")
	return nil
}

// GetChannel returns the RabbitMQ channel
func GetChannel() *amqp.Channel {
	return channel
}

// CloseConnection closes the RabbitMQ connection and channel
func CloseConnection() {
	if channel != nil {
		channel.Close()
	}
	if conn != nil {
		conn.Close()
	}
}

// DeclareTestQueue declares a test queue for our initial setup
func DeclareTestQueue() error {
	_, err := channel.QueueDeclare(
		"test_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,         // arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}
	return nil
}
