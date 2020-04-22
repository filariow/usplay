package storage

import (
	"time"
)

// Activity activity model for storage
type Activity struct {
	ID             string    `bson:"_id"`
	ActivityTypeID string    `bson:"activitytype_id"`
	OrderID        string    `bson:"order_id"`
	Description    string    `bson:"description"`
	Code           string    `bson:"code"`
	Name           string    `bson:"name"`
	CreationTime   time.Time `bson:"creation_time"`
	Interval       Interval  `bson:"interval"`
}

//Interval interval data structure
type Interval struct {
	From time.Time `bson:"from"`
	To   time.Time `bson:"to"`
}

// Activities list of activities
type Activities []Activity
