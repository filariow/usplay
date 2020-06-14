package main

import (
	"log"
	"net"

	"github.com/FrancescoIlario/usplay/cmd/services/order/server/data"
	"github.com/FrancescoIlario/usplay/internal/services/order/api"
	"github.com/FrancescoIlario/usplay/internal/services/order/storage"
	"github.com/FrancescoIlario/usplay/pkg/osext"
	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"

	"google.golang.org/grpc"
)

const (
	addressKey     = "US_ADDRESS"
	addressDefault = "localhost:8080"
)

func main() {
	log.Println("Starting Server")

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

	actServer := api.NewOrderServer(repo)
	grpcServer := grpc.NewServer()
	ordergrpc.RegisterOrderSvcServer(grpcServer, actServer)

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
