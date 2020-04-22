package main

import (
	"log"
	"net"
	"os"

	"github.com/FrancescoIlario/usplay/internal/services/report/api"
	"github.com/FrancescoIlario/usplay/internal/services/report/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"

	//ActivityTargetKey target env key where is stored the Activity Host
	ActivityTargetKey = "US_ACTIVITY_TARGET"
)

func main() {
	log.Println("Hello world!")

	ls, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen at %v: %v", address, err)
	}
	log.Printf("acquired address %v", address)

	actTypeHost := os.Getenv(ActivityTargetKey)
	if actTypeHost == "" {
		log.Fatalf("env variable %s not set", ActivityTargetKey)
	}

	conn, err := grpc.Dial(actTypeHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to %s: %v", actTypeHost, err)
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
