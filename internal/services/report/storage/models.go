package storage

import (
	"time"
)

// Report report model for storage
type Report struct {
	ID          string
	Description string
	Code        string
	Name        string
	Activities  Activities
}

// Reports list of reports
type Reports []Report

// Activity an user activity
type Activity struct {
	ID     string
	Period Interval
	Order  Order
	Type   ActivityType
}

// Activities a list of activities
type Activities []Activity

// ActivityType the activity type
type ActivityType struct {
	ID   string
	Code int
	Name string
}

// Order order
type Order struct {
	ID   string
	Name string
}

// Interval an interval between two dates
type Interval struct {
	From time.Time
	To   time.Time
}
