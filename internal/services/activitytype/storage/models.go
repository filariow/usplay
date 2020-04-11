package storage

import "github.com/google/uuid"

// ActivityType ActivityType model for storage
type ActivityType struct {
	ID          uuid.UUID
	Code        int32
	Description string
	Name        string
}

// Activities list of activities
type Activities []ActivityType
