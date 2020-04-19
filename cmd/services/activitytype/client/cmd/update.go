package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "Updates a new ActivityType",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitytypecomm.NewActivityTypeSvcClient(conn)

			if _, err := cli.Update(context.TODO(),
				&activitytypecomm.UpdateActivityTypeRequest{
					Id:          id,
					Code:        code,
					Name:        name,
					Description: desc,
				}); err != nil {
				log.Fatalf("error calling update: %v", err)
			}

			log.Printf("updated ActivityType: %s", id)
		},
	}
)

func init() {
	cmdUpdate.Flags().Int32VarP(&code, "code", "c", 0, "ActivityType's code")
	cmdUpdate.Flags().StringVarP(&desc, "description", "d", "", "ActivityType's description")
	cmdUpdate.Flags().StringVarP(&name, "name", "n", "", "ActivityType's name")
	cmdUpdate.Flags().StringVarP(&id, "id", "i", "", "ActivityType's id")
}
