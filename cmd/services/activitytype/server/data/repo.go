package data

import (
	"fmt"

	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage/inmemstore"
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage/mongostore"
)

//BuildStore builds the repository for the given configuration
func BuildStore(conf *Configuration) (storage.Repository, error) {
	if conf.Store == nil {
		return nil, fmt.Errorf("Invalid repository configuration found")
	}

	switch c := conf.Store.(type) {
	case *mongostore.Configuration:
		return mongostore.NewRepository(c)
	case *inmemstore.Configuration:
		return inmemstore.New(), nil
	default:
		return nil, fmt.Errorf(`Not recognized configuration type "%T"`, c)
	}
}
