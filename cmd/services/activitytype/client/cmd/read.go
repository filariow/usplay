package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdRead = &cobra.Command{
		Use:   "read",
		Short: "Reads an ActivityType",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitytypecomm.NewActivityTypeSvcClient(conn)
			resp, err := cli.Read(context.TODO(), &activitytypecomm.ReadActivityTypeRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling read: %v", err)
			}

			log.Printf(
				"read ActivityType:\n\tid: %s\n\tcode: %v\n\tname: %s",
				resp.ActivityType.Id,
				resp.ActivityType.Code,
				resp.ActivityType.Name)
		},
	}

	id string
)

func init() {
	cmdRead.PersistentFlags().StringVarP(&id, "id", "i", "", "ActivityType's id")
}
