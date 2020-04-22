package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "Updates a new activity",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitycomm.NewActivitySvcClient(conn)
			if _, err = cli.Update(context.TODO(), &activitycomm.UpdateActivityRequest{
				Code:        code,
				Description: desc,
				Name:        name,
			}); err != nil {
				log.Fatalf("error calling update: %v", err)
			}

			log.Printf("updated activity %s", id)
		},
	}
)

func init() {
	cmdUpdate.Flags().StringVarP(&code, "code", "c", "", "activity's code")
	cmdUpdate.Flags().StringVarP(&desc, "description", "d", "", "activity's description")
	cmdUpdate.Flags().StringVarP(&name, "name", "n", "", "activity's name")
	cmdUpdate.Flags().StringVarP(&id, "id", "i", "", "activity's id")
}
