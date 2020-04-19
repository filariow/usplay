package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdCreate = &cobra.Command{
		Use:   "create",
		Short: "Creates a new ActivityType",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitytypecomm.NewActivityTypeSvcClient(conn)
			resp, err := cli.Create(context.TODO(), &activitytypecomm.CreateActivityTypeRequest{
				Code:        code,
				Name:        name,
				Description: desc,
			})
			if err != nil {
				log.Fatalf("error calling create: %v", err)
			}

			log.Printf("created new ActivityType with id: %s", resp.GetId())
		},
	}
)

func init() {
	cmdCreate.Flags().Int32VarP(&code, "code", "c", 0, "ActivityType's code")
	cmdCreate.Flags().StringVarP(&desc, "description", "d", "", "ActivityType's description")
	cmdCreate.Flags().StringVarP(&name, "name", "n", "", "ActivityType's name")
}
