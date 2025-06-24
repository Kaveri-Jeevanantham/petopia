package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
)

// MessageHandler is a function type that handles received messages
type MessageHandler func(Message) error

// ConsumeMessages starts consuming messages from the specified queue
func ConsumeMessages(queueName string, handler MessageHandler) error {
	msgs, err := channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}

	go func() {
		for msg := range msgs {
			var message Message
			if err := json.Unmarshal(msg.Body, &message); err != nil {
				log.Printf("Error unmarshaling message: %v", err)
				continue
			}

			if err := handler(message); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}()

	log.Printf("Started consuming messages from queue: %s", queueName)
	return nil
}

// StartTestConsumer starts a consumer for the test queue
func StartTestConsumer() error {
	return ConsumeMessages("test_queue", func(msg Message) error {
		log.Printf("Received test message: %+v", msg)
		return nil
	})
}
