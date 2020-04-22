package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
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

			cli := activitycomm.NewActivitySvcClient(conn)
			resp, err := cli.Create(context.TODO(), &activitycomm.CreateActivityRequest{
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
	cmdCreate.Flags().StringVarP(&code, "code", "c", "", "activity's code")
	cmdCreate.Flags().StringVarP(&desc, "description", "d", "", "activity's description")
	cmdCreate.Flags().StringVarP(&name, "name", "n", "", "activity's name")
}
