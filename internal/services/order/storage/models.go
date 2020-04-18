package storage

import "github.com/google/uuid"

// Order order model for storage
type Order struct {
	ID          uuid.UUID
	Code        string
	Description string
	Name        string
}

// Orders list of orders
type Orders []Order
