package storage

import "github.com/google/uuid"

// ActivityType ActivityType model for storage
type ActivityType struct {
	ID          uuid.UUID `bson:"_id"`
	Code        int32     `bson:"code"`
	Description string    `bson:"description"`
	Name        string    `bson:"name"`
}

// Activities list of activities
type Activities []ActivityType
