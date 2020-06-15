package conf

import (
	"github.com/kelseyhightower/envconfig"
)

// EnvPrefix prefix for environmental variables parsed by application
const EnvPrefix = "USPLAY"

// Configuration Bookmaster service configuration
type Configuration struct {
	OrderHost                   string `required:"true" split_words:"true"`
	Address                     string `default:"localhost:8080"`
	ActivityMongoConnstr        string `required:"true" split_words:"true"`
	ActivityMongoDatabase       string `required:"true" split_words:"true"`
	ActivityMongoCollection     string `required:"true" split_words:"true"`
	ActivityMongoUsername       string `required:"true" split_words:"true"`
	ActivityMongoPassword       string `required:"true" split_words:"true"`
	ActivitytypeMongoConnstr    string `required:"true" split_words:"true"`
	ActivitytypeMongoDatabase   string `required:"true" split_words:"true"`
	ActivitytypeMongoCollection string `required:"true" split_words:"true"`
	ActivitytypeMongoUsername   string `required:"true" split_words:"true"`
	ActivitytypeMongoPassword   string `required:"true" split_words:"true"`
}

// GetConfigurationFromEnv ...
func GetConfigurationFromEnv() (s Configuration, err error) {
	err = envconfig.Process(EnvPrefix, &s)
	return
}
