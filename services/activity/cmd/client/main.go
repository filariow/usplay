package main

import (
	"context"
	"log"
	"usplay/us-activity/cmd/client/cmd"
	"usplay/us-activity/gen"

	"google.golang.org/grpc"
)

const target = "localhost:8080"

func main() {
	cmd.Execute()
}

func callCreate() {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect to %s: %v", target, err)
	}
	defer conn.Close()

	cli := gen.NewActivitySvcClient(conn)
	resp, err := cli.Create(context.TODO(), &gen.CreateActivityRequest{
		Code:        "code",
		Description: "description",
		Name:        "name",
	})
	if err != nil {
		log.Fatalf("error calling create: %v", err)
	}

	log.Printf("created new activity with id: %s", resp.GetId())
}
