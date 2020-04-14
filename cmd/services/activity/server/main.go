package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/api"
	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"

	"google.golang.org/grpc"
)

const (
	actTypeHostKey = "ACTTYPE_HOST"
	address        = "localhost:8080"
)

func main() {
	actTypeHost := os.Getenv(actTypeHostKey)
	if actTypeHost == "" {
		log.Fatalf("ActivityType Host address environment variable is empty %s", actTypeHostKey)
	}

	ls, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen at %v: %v", address, err)
	}
	log.Printf("acquired address %v", address)

	store := storage.NewInMemoryStore()

	conn, err := grpc.Dial(actTypeHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to %s: %v", actTypeHost, err)
	}
	actTypeCli := activitytypecomm.NewActivityTypeSvcClient(conn)

	actServer := api.NewActivityServer(store, actTypeCli, 1*time.Second)
	grpcServer := grpc.NewServer()
	activitycomm.RegisterActivitySvcServer(grpcServer, actServer)

	log.Printf("starting server at %v", address)
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
