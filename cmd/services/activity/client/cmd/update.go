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
			interval, err := parseInterval(from, to)
			if err != nil {

			}
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitycomm.NewActivitySvcClient(conn)
			if _, err = cli.Update(context.TODO(), &activitycomm.UpdateActivityRequest{
				ActTypeID: actid,
				OrderID:   ordid,
				Period:    interval,
			}); err != nil {
				log.Fatalf("error calling update: %v", err)
			}

			log.Printf("updated activity %s", id)
		},
	}
)

func init() {
	cmdUpdate.Flags().StringVarP(&id, "id", "i", "", "activity's id")
	cmdUpdate.MarkFlagRequired("id")
	cmdUpdate.Flags().StringVarP(&actid, "activity-type", "a", "", "activityType's ID")
	cmdUpdate.MarkFlagRequired("activity-type")
	cmdUpdate.Flags().StringVarP(&ordid, "order", "o", "", "Order ID")
	cmdUpdate.MarkFlagRequired("order")
	cmdUpdate.Flags().StringVarP(&from, "from", "F", "", "Interval's From")
	cmdUpdate.MarkFlagRequired("from")
	cmdUpdate.Flags().StringVarP(&to, "to", "T", "", "Interval's To")
	cmdUpdate.MarkFlagRequired("to")
}
