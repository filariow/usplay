package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/FrancescoIlario/usplay/cmd/services/activity/server/data"
	"github.com/FrancescoIlario/usplay/internal/services/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/osext"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"

	"google.golang.org/grpc"
)

const (
	actTypeHostKey = "US_ACTTYPE_HOST"
	orderHostKey   = "US_ORDER_HOST"
	addressKey     = "US_ADDRESS"
	addressDefault = "localhost:8080"
)

func main() {
	log.Println("Starting Server")

	actTypeHost := os.Getenv(actTypeHostKey)
	if actTypeHost == "" {
		log.Fatalf("ActivityType Host address environment variable is empty %s", actTypeHostKey)
	}
	orderHost := os.Getenv(orderHostKey)
	if orderHost == "" {
		log.Fatalf("Order Host address environment variable is empty %s", orderHostKey)
	}

	address := osext.GetEnvOrDefault(addressKey, addressDefault)

	ls, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen at %v: %v", address, err)
	}
	log.Printf("acquired address %v", address)

	connActType, err := grpc.Dial(actTypeHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to %s: %v", actTypeHost, err)
	}
	actTypeCli := activitytypecomm.NewActivityTypeSvcClient(connActType)

	connOrders, err := grpc.Dial(orderHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to %s: %v", orderHost, err)
	}
	orderCli := ordercomm.NewOrderSvcClient(connOrders)

	repo, err := setUpRepository()
	if err != nil {
		log.Fatalf("error setting up repository: %v", err)
	}

	actServer := api.NewActivityServer(repo, actTypeCli, orderCli, 1*time.Second)
	grpcServer := grpc.NewServer()
	activitycomm.RegisterActivitySvcServer(grpcServer, actServer)

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
