package storage

import "github.com/google/uuid"

// Order order model for storage
type Order struct {
	ID          uuid.UUID
	Description string
	Code        string
	Name        string
}

// Activities list of activities
type Activities []Order
