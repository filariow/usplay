package main

import (
	"github.com/google/uuid"
)

// Repository repository for todo
type Repository interface {
	Create(Todo) (Todo, error)
	Update(Todo) (*Todo, error)
	Delete(uuid.UUID) (*Todo, error)
	ReadAll() ([]Todo, error)
	Read(uuid.UUID) (*Todo, error)
}
