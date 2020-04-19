package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
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

			cli := activitytypecomm.NewActivityTypeSvcClient(conn)
			resp, err := cli.List(context.TODO(), &activitytypecomm.ListActivityTypesRequest{
				FilterIds: listIds,
			})
			if err != nil {
				log.Fatalf("error calling list: %v", err)
			}

			for _, v := range resp.GetActivityTypes() {
				log.Printf(
					"list ActivityType:\n\tid: %s\n\tcode: %v\n\tname: %s",
					v.Id,
					v.Code,
					v.Name)
			}
		},
	}

	listIds = []string{}
)

func init() {
	cmdList.Flags().StringArrayVarP(&listIds, "ids", "i", []string{}, "filter IDs")
}
