package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "Deletes an activity",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitycomm.NewActivitySvcClient(conn)
			resp, err := cli.Delete(context.TODO(), &activitycomm.DeleteActivityRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling delete: %v", err)
			}

			log.Printf(
				"deleted activity:\n\tid: %s\n\tcode: %s\n\tdescription: %s\n\tname: %s",
				resp.Activity.Id,
				resp.Activity.Code,
				resp.Activity.Description,
				resp.Activity.Name)
		},
	}
)

func init() {
	cmdDelete.PersistentFlags().StringVarP(&id, "id", "i", "", "activity's id")
}
