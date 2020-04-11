package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytype/comm"
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

			cli := comm.NewActivityTypeSvcClient(conn)
			resp, err := cli.Create(context.TODO(), &comm.CreateActivityTypeRequest{
				Code: code,
				Name: name,
			})
			if err != nil {
				log.Fatalf("error calling create: %v", err)
			}

			log.Printf("created new ActivityType with id: %s", resp.GetId())
		},
	}
)

func init() {
	cmdCreate.PersistentFlags().Int32VarP(&code, "code", "c", 0, "ActivityType's code")
	cmdCreate.PersistentFlags().StringVarP(&desc, "description", "d", "", "ActivityType's description")
	cmdCreate.PersistentFlags().StringVarP(&name, "name", "n", "", "ActivityType's name")
}
