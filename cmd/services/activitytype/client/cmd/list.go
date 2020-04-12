package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytype/comm"

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

			cli := comm.NewActivityTypeSvcClient(conn)
			resp, err := cli.List(context.TODO(), &comm.ListActivityTypesRequest{
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

	listIds = make([]string, 0)
)

func init() {
	cmdList.PersistentFlags().StringArrayVarP(&listIds, "i", "ids", nil, "filter IDs")
}
