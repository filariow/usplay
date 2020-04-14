package storage

import (
	"time"

	"github.com/google/uuid"
)

// Activity activity model for storage
type Activity struct {
	ID             uuid.UUID
	ActivityTypeID uuid.UUID
	Description    string
	Code           string
	Name           string
	CreationTime   time.Time
}

// Activities list of activities
type Activities []Activity
