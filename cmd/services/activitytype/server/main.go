package main

import (
	"log"
	"net"

	"github.com/FrancescoIlario/usplay/cmd/services/activitytype/server/data"
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/api"
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/osext"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"

	"google.golang.org/grpc"
)

const (
	addressKey     = "US_ADDRESS"
	addressDefault = "localhost:8080"
)

func main() {
	log.Println("Starting server")

	address := osext.GetEnvOrDefault(addressKey, addressDefault)
	ls, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen at %v: %v", address, err)
	}
	log.Printf("acquired address %v", address)

	repo, err := setUpRepository()
	if err != nil {
		log.Fatalf("error setting up repository: %v", err)
	}

	actServer := api.NewActivityTypeServer(repo)
	grpcServer := grpc.NewServer()
	activitytypecomm.RegisterActivityTypeSvcServer(grpcServer, actServer)

	log.Printf("starting server at %v", address)
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}

func setUpRepository() (storage.Repository, error) {
	configuration, err := data.ParseFromEnvs()
	if err != nil {
		return nil, err
	}

	return data.BuildStore(configuration)
}
