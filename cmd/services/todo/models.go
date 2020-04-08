package main

import "github.com/google/uuid"

// Todo todo model
type Todo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

// ErrorMessage error message for communication
type ErrorMessage struct {
	Message        string                 `json:"message"`
	Code           int                    `json:"code"`
	ErrorMessage   string                 `json:"errorMessage"`
	AdditionalData map[string]interface{} `json:"additionalData"`
}
