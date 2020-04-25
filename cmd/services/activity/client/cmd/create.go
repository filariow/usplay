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
			interval, err := parseInterval(from, to)
			if err != nil {
				log.Fatal("error parsing interval: %v", err)
			}

			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitycomm.NewActivitySvcClient(conn)
			resp, err := cli.Create(context.TODO(),
				&activitycomm.CreateActivityRequest{
					ActTypeID: actid,
					OrderID:   ordid,
					Period:    interval,
				})
			if err != nil {
				log.Fatalf("error calling create: %v", err)
			}

			log.Printf("created new activity with id: %s", resp.GetId())
		},
	}

	actid string
	ordid string
)

func init() {
	cmdCreate.Flags().StringVarP(&actid, "activity-type", "a", "", "activityType's ID")
	cmdCreate.MarkFlagRequired("activity-type")
	cmdCreate.Flags().StringVarP(&ordid, "order", "o", "", "Order ID")
	cmdCreate.MarkFlagRequired("order")
	cmdCreate.Flags().StringVarP(&from, "from", "F", "", "Interval's From")
	cmdCreate.MarkFlagRequired("from")
	cmdCreate.Flags().StringVarP(&to, "to", "T", "", "Interval's To")
	cmdCreate.MarkFlagRequired("to")
}
