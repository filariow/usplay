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
					Code: code,
					Name: name}); err != nil {
				log.Fatalf("error calling update: %v", err)
			}

			log.Printf("updated ActivityType: %s", id)
		},
	}
)

func init() {
	cmdUpdate.PersistentFlags().Int32VarP(&code, "code", "c", 0, "ActivityType's code")
	cmdUpdate.PersistentFlags().StringVarP(&desc, "description", "d", "", "ActivityType's description")
	cmdUpdate.PersistentFlags().StringVarP(&name, "name", "n", "", "ActivityType's name")
	cmdUpdate.PersistentFlags().StringVarP(&id, "id", "i", "", "ActivityType's id")
}
