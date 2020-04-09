package main

import (
	"log"
	"net"

	"github.com/FrancescoIlario/usplay/internal/report/api"
	"github.com/FrancescoIlario/usplay/pkg/services/report/comm"
	"google.golang.org/grpc"
)

const address = "localhost:8080"

func main() {
	log.Println("Hello world!")

	ls, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen at %v: %v", address, err)
	}
	log.Printf("acquired address %v", address)

	actServer := api.NewReportServer()
	grpcServer := grpc.NewServer()
	comm.RegisterReportSvcServer(grpcServer, actServer)

	log.Printf("starting server at %v", address)
	if err := grpcServer.Serve(ls); err != nil {
		log.Fatalf("Error serving: %v", err)
	}
}
