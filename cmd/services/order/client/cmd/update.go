package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
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

			cli := ordergrpc.NewOrderSvcClient(conn)
			resp, err := cli.Update(context.TODO(), &ordergrpc.UpdateOrderRequest{
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
	cmdUpdate.Flags().StringVarP(&code, "code", "c", "", "order's code")
	cmdUpdate.Flags().StringVarP(&desc, "description", "d", "", "order's description")
	cmdUpdate.Flags().StringVarP(&name, "name", "n", "", "order's name")
	cmdUpdate.Flags().StringVarP(&id, "id", "i", "", "order's id")
}
