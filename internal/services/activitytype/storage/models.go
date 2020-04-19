package storage

// ActivityType ActivityType model for storage
type ActivityType struct {
	ID          string `bson:"_id"`
	Code        int32  `bson:"code"`
	Description string `bson:"description"`
	Name        string `bson:"name"`
}

// ActivityTypes list of activities
type ActivityTypes []ActivityType
