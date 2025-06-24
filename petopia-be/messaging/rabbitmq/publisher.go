package rabbitmq

import (
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

// Message represents a generic message structure
type Message struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

// PublishMessage publishes a message to the specified queue
func PublishMessage(queueName string, message Message) error {
	// Convert message to JSON
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	// Publish the message
	err = channel.Publish(
		"",        // exchange
		queueName, // routing key (queue name)
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBytes,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}

// PublishTestMessage publishes a test message to the test queue
func PublishTestMessage(message string) error {
	testMessage := Message{
		Type: "test",
		Payload: map[string]string{
			"message": message,
		},
	}

	return PublishMessage("test_queue", testMessage)
}
