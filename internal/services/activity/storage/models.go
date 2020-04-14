package storage

import "github.com/google/uuid"

// Activity activity model for storage
type Activity struct {
	ID             uuid.UUID
	ActivityTypeID uuid.UUID
	Description    string
	Code           string
	Name           string
}

// Activities list of activities
type Activities []Activity
