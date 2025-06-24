package controller

import (
	"encoding/json"
	"net/http"
	"petopia-be/messaging/rabbitmq"
)

// TestMessageRequest represents the request body for test message
type TestMessageRequest struct {
	Message string `json:"message"`
}

// SendTestMessage handles sending a test message to RabbitMQ
func SendTestMessage(w http.ResponseWriter, r *http.Request) {
	var req TestMessageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := rabbitmq.PublishTestMessage(req.Message); err != nil {
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Test message sent successfully",
	})
}
