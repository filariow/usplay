package main

import (
	"log"
	"net"
	"os"

	"github.com/FrancescoIlario/usplay/internal/services/report/api"
	"github.com/FrancescoIlario/usplay/internal/services/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/osext"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"google.golang.org/grpc"
)

const (
	addressDefault = "localhost:8080"
	addressKey     = "US_ADDRESS"

	//ActivityTargetKey target env key where is stored the Activity Host
	ActivityTargetKey = "US_ACTIVITY_HOST"
)

func main() {
	log.Println("Starting Server")

	address := osext.GetEnvOrDefault(addressKey, addressDefault)

	ls, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen at %v: %v", address, err)
	}
	log.Printf("acquired address %v", address)

	activityHost := os.Getenv(ActivityTargetKey)
	if activityHost == "" {
		log.Fatalf("env variable %s not set", ActivityTargetKey)
	}

	conn, err := grpc.Dial(activityHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to %s: %v", activityHost, err)
	}
	activityCli := activitycomm.NewActivitySvcClient(conn)
	store := storage.NewInMemoryStore()
	actServer := api.NewReportServer(activityCli, store)
	grpcServer := grpc.NewServer()
	reportcomm.RegisterReportSvcServer(grpcServer, actServer)

	log.Printf("starting server at %v", address)
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
