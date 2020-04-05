package cmd

import (
	"context"
	"log"
	"usplay/us-activity/gen"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdList = &cobra.Command{
		Use:   "list",
		Short: "Lists all the activities",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := gen.NewActivitySvcClient(conn)
			resp, err := cli.List(context.TODO(), &gen.ListActivitiesRequest{})
			if err != nil {
				log.Fatalf("error calling list: %v", err)
			}

			for _, v := range resp.GetActivities() {
				log.Printf(
					"list activity:\n\tid: %s\n\tcode: %s\n\tdescription: %s\n\tname: %s",
					v.Id,
					v.Code,
					v.Description,
					v.Name)
			}
		},
	}
)

func init() {
	cmdList.PersistentFlags().StringVarP(&id, "id", "i", "", "activity's id")
}
