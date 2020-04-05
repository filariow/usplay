package cmd

import (
	"context"
	"log"
	"usplay/us-activity/gen"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdCreate = &cobra.Command{
		Use:   "create",
		Short: "Creates a new activity",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := gen.NewActivitySvcClient(conn)
			resp, err := cli.Create(context.TODO(), &gen.CreateActivityRequest{
				Code:        code,
				Description: desc,
				Name:        name,
			})
			if err != nil {
				log.Fatalf("error calling create: %v", err)
			}

			log.Printf("created new activity with id: %s", resp.GetId())
		},
	}

	code string
	desc string
	name string
)

func init() {
	cmdCreate.PersistentFlags().StringVarP(&code, "code", "c", "", "activity's code")
	cmdCreate.PersistentFlags().StringVarP(&desc, "description", "d", "", "activity's description")
	cmdCreate.PersistentFlags().StringVarP(&name, "name", "n", "", "activity's name")
}
