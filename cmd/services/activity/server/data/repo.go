package data

import (
	"fmt"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage/inmemstore"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage/mongostore"
)

//BuildStore builds the repository for the given configuration
func BuildStore(conf *Configuration) (storage.Repository, error) {
	if conf.Store == nil {
		return nil, fmt.Errorf("Invalid repository configuration found")
	}

	switch c := conf.Store.(type) {
	case *mongostore.Configuration:
		return mongostore.New(c)
	case *inmemstore.Configuration:
		return inmemstore.New(), nil
	default:
		return nil, fmt.Errorf(`Not recognized configuration type "%T"`, c)
	}
}
