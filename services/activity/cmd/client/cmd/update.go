package cmd

import (
	"context"
	"log"
	"usplay/us-activity/gen"

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

			cli := gen.NewActivitySvcClient(conn)
			resp, err := cli.Update(context.TODO(), &gen.UpdateActivityRequest{
				Code:        code,
				Description: desc,
				Name:        name,
			})
			if err != nil {
				log.Fatalf("error calling update: %v", err)
			}

			log.Printf("updated activity: %s", resp.Activity.Id)
		},
	}
)

func init() {
	cmdUpdate.PersistentFlags().StringVarP(&code, "code", "c", "", "activity's code")
	cmdUpdate.PersistentFlags().StringVarP(&desc, "description", "d", "", "activity's description")
	cmdUpdate.PersistentFlags().StringVarP(&name, "name", "n", "", "activity's name")
	cmdUpdate.PersistentFlags().StringVarP(&id, "id", "i", "", "activity's id")
}
