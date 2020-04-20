package storage

// Order order model for storage
type Order struct {
	ID          string `bson:"_id"`
	Code        string `bson:"code"`
	Description string `bson:"description"`
	Name        string `bson:"name"`
}

// Orders list of orders
type Orders []Order
