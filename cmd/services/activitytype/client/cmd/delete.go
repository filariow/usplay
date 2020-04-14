package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "Deletes an ActivityType",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitytypecomm.NewActivityTypeSvcClient(conn)
			resp, err := cli.Delete(context.TODO(), &activitytypecomm.DeleteActivityTypeRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling delete: %v", err)
			}

			log.Printf(
				"deleted ActivityType:\n\tid: %s\n\tcode: %v\n\tname: %s",
				resp.ActivityType.Id,
				resp.ActivityType.Code,
				resp.ActivityType.Name)
		},
	}
)

func init() {
	cmdDelete.PersistentFlags().StringVarP(&id, "id", "i", "", "ActivityType's id")
}
