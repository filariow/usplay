package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStore struct {
	Client        *mongo.Client
	Collection    *mongo.Collection
	Configuration MongoConfiguration
}

// MongoConfiguration configuration for Mongo storage
type MongoConfiguration struct {
	ConnectionString string
	Database         string
	Collection       string
}

// NewRepository New Mongo database instance
func NewRepository(conf MongoConfiguration) (storage.Repository, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(conf.ConnectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	collection := client.
		Database(conf.Database).
		Collection(conf.Collection)

	// build database
	db := mongoStore{
		Client:        client,
		Collection:    collection,
		Configuration: conf,
	}
	return &db, nil
}
