package data

import (
	"os"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage/inmemstore"
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage/mongostore"
	"github.com/FrancescoIlario/usplay/pkg/errs/conferrs"
)

const (
	storageTypeEnvKey = "US_STORAGE_TYPE"
	mongoConnString   = "US_MONGO_CONNSTR"
	mongoDatabase     = "US_MONGO_DATABASE"
	mongoCollection   = "US_MONGO_COLLECTION"
	mongoUsername     = "US_MONGO_USERNAME"
	mongoPassword     = "US_MONGO_PASSWORD"
)

//storeType type of storage
type storeType string

const (
	//InMemory In memory StoreType
	InMemory storeType = "InMemory"
	//Mongo MongoDB StoreType
	Mongo storeType = "Mongo"
)

//Configuration Application configuration
type Configuration struct {
	Store interface{}
}

//ParseFromEnvs builds App configuration from environment variables
func ParseFromEnvs() (*Configuration, error) {
	store, err := parseStorageFromEnvs()
	if err != nil {
		return nil, err
	}

	return &Configuration{
		Store: store,
	}, nil
}

func parseStorageFromEnvs() (interface{}, error) {
	storageTypeStr := os.Getenv(storageTypeEnvKey)
	if storageTypeStr == "" {
		return nil, conferrs.NewEnvKeyNotFoundError(storageTypeEnvKey)
	}

	storeType := storeType(storageTypeStr)
	switch storeType {
	case InMemory:
		return &inmemstore.Configuration{}, nil
	case Mongo:
		return parseMongoStorageConfFromEnvs()
	default:
		return nil, conferrs.NewEnvKeyInvalidError(storageTypeEnvKey, storageTypeStr)
	}
}

func parseMongoStorageConfFromEnvs() (*mongostore.Configuration, error) {
	connStr := os.Getenv(mongoConnString)
	database := os.Getenv(mongoDatabase)
	collection := os.Getenv(mongoCollection)
	username := os.Getenv(mongoUsername)
	password := os.Getenv(mongoPassword)

	return &mongostore.Configuration{
		ConnectionString: connStr,
		Database:         database,
		Collection:       collection,
		Username:         username,
		Password:         password,
	}, nil
}
