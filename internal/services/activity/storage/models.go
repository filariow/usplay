package storage

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// Activity activity model for storage
type Activity struct {
	ID             string   `bson:"_id"`
	ActivityTypeID string   `bson:"activitytype_id"`
	OrderID        string   `bson:"order_id"`
	Period         Interval `bson:"interval"`
}

//Interval interval data structure
type Interval struct {
	From time.Time `bson:"from"`
	To   time.Time `bson:"to"`
}

//NewIntervalProto creates the interval from protos
func NewIntervalProto(from, to *timestamp.Timestamp) (*Interval, error) {
	_f, err := ptypes.Timestamp(from)
	if err != nil {
		return nil, err
	}

	_t, err := ptypes.Timestamp(to)
	if err != nil {
		return nil, err
	}

	return &Interval{From: _f, To: _t}, nil
}

// Activities list of activities
type Activities []Activity
