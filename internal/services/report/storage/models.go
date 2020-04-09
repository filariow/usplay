package storage

import (
	"time"

	"github.com/google/uuid"
)

// Report report model for storage
type Report struct {
	ID          uuid.UUID
	Description string
	Code        string
	Name        string
	Activities  Activities
}

// Reports list of reports
type Reports []Report

// Activity an user activity
type Activity struct {
	ID      uuid.UUID
	OrderID uuid.UUID
	Type    ActivityType
}

// Activities a list of activities
type Activities []Activity

// ActivityType the activity type
type ActivityType struct {
	ID   uuid.UUID
	Code int
	Name string
}

// Timespan an interval between two dates
type Timespan struct {
	From time.Time
	To   time.Time
}
