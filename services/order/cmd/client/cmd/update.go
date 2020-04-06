package cmd

import (
	"context"
	"log"
	"usplay/us-order/gen"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdUpdate = &cobra.Command{
		Use:   "update",
		Short: "Updates a new order",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := gen.NewOrderSvcClient(conn)
			resp, err := cli.Update(context.TODO(), &gen.UpdateOrderRequest{
				Code:        code,
				Description: desc,
				Name:        name,
			})
			if err != nil {
				log.Fatalf("error calling update: %v", err)
			}

			log.Printf("updated order: %s", resp.Order.Id)
		},
	}
)

func init() {
	cmdUpdate.PersistentFlags().StringVarP(&code, "code", "c", "", "order's code")
	cmdUpdate.PersistentFlags().StringVarP(&desc, "description", "d", "", "order's description")
	cmdUpdate.PersistentFlags().StringVarP(&name, "name", "n", "", "order's name")
	cmdUpdate.PersistentFlags().StringVarP(&id, "id", "i", "", "order's id")
}
