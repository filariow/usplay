package main

import (
	"github.com/kelseyhightower/envconfig"
)

// EnvPrefix prefix for environmental variables parsed by application
const EnvPrefix = "USPLAY"

type configuration struct {
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

func getConfigurationFromEnv() (s configuration, err error) {
	err = envconfig.Process(EnvPrefix, &s)
	return
}
