package main

import (
	"context"
	"log"
	"net"

	"github.com/FrancescoIlario/usplay/cmd/services/activitytype/server/data"
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/api"
	"github.com/FrancescoIlario/usplay/internal/services/activitytype/storage"
	"github.com/FrancescoIlario/usplay/pkg/osext"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	errHandler := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		log.Printf("method %q request: %v", info.FullMethod, req)
		resp, err := handler(ctx, req)
		if err != nil {
			log.Printf("method %q failed: %s", info.FullMethod, err)
		} else {
			log.Printf("method %q successed: ", info.FullMethod)
		}

		return resp, err
	}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(errHandler))

	reflection.Register(grpcServer)
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
