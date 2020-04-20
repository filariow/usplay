package mongostore

import (
	"context"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoStore struct {
	Client        *mongo.Client
	Collection    *mongo.Collection
	Configuration Configuration
}

// Configuration configuration for Mongo storage
type Configuration struct {
	ConnectionString string
	Database         string
	Collection       string
	Username         string
	Password         string
}

// New New Mongo database instance
func New(conf *Configuration) (storage.Repository, error) {
	// Set client options
	clientOptions := options.Client().
		ApplyURI(conf.ConnectionString)

	if conf.Username != "" {
		clientOptions = clientOptions.
			SetAuth(options.Credential{
				Username: conf.Username,
				Password: conf.Password,
			})
	}

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
	stconf := *conf
	stconf.Password = ""

	db := mongoStore{
		Client:        client,
		Collection:    collection,
		Configuration: stconf,
	}
	return &db, nil
}
