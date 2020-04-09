package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/order/comm"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdCreate = &cobra.Command{
		Use:   "create",
		Short: "Creates a new order",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := comm.NewOrderSvcClient(conn)
			resp, err := cli.Create(context.TODO(), &comm.CreateOrderRequest{
				Code:        code,
				Description: desc,
				Name:        name,
			})
			if err != nil {
				log.Fatalf("error calling create: %v", err)
			}

			log.Printf("created new order with id: %s", resp.GetId())
		},
	}

	code string
	desc string
	name string
)

func init() {
	cmdCreate.PersistentFlags().StringVarP(&code, "code", "c", "", "order's code")
	cmdCreate.PersistentFlags().StringVarP(&desc, "description", "d", "", "order's description")
	cmdCreate.PersistentFlags().StringVarP(&name, "name", "n", "", "order's name")
}
