package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdRead = &cobra.Command{
		Use:   "read",
		Short: "Reads an activity",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitycomm.NewActivitySvcClient(conn)
			resp, err := cli.Read(context.TODO(), &activitycomm.ReadActivityRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling read: %v", err)
			}

			v := resp.GetActivity()
			log.Printf(`read activity:
	id: %s
	activity type: %s
	order: %s
	period: 
		from: %s
		to: %s`,
				v.Id, v.ActType.Id, v.Order.Id,
				v.Period.From, v.Period.To)
		},
	}

	id string
)

func init() {
	cmdRead.Flags().StringVarP(&id, "id", "i", "", "activity's id")
}
