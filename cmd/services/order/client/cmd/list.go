package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/ordergrpc"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdList = &cobra.Command{
		Use:   "list",
		Short: "Lists all the orders",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := ordergrpc.NewOrderSvcClient(conn)
			resp, err := cli.List(context.TODO(), &ordergrpc.ListOrdersRequest{})
			if err != nil {
				log.Fatalf("error calling list: %v", err)
			}

			for _, v := range resp.GetOrders() {
				log.Printf(
					"list order:\n\tid: %s\n\tcode: %s\n\tdescription: %s\n\tname: %s",
					v.Id,
					v.Code,
					v.Description,
					v.Name)
			}
		},
	}
)

func init() {
	cmdList.Flags().StringVarP(&id, "id", "i", "", "order's id")
}
