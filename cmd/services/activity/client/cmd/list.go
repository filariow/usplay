package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
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

			cli := activitycomm.NewActivitySvcClient(conn)
			resp, err := cli.List(context.TODO(),
				&activitycomm.ListActivitiesRequest{FilterIds: ids})
			if err != nil {
				log.Fatalf("error calling list: %v", err)
			}

			for _, v := range resp.GetActivities() {
				order, activitytype := "", ""
				if v.Order != nil {
					order = fmt.Sprintf(`
		id: %s
		code: %s
		description: %s`,
						v.Order.Id, v.Order.Code, v.Order.Description)
				}
				if v.ActType != nil {
					activitytype = fmt.Sprintf(`
		id: %s
		code: %v
		description: %s`,
						v.ActType.Id, v.ActType.Code, v.ActType.Description)
				}

				log.Printf(`list activity:
	id: %s
	activity type: %s
	order: %s
	period: 
		from: %s
		to: %s`,
					v.Id, activitytype, order,
					v.Period.From, v.Period.To)
			}
		},
	}

	ids []string
)

func init() {
	cmdList.Flags().StringArrayVarP(&ids, "ids", "i", nil, "activities id")
}
