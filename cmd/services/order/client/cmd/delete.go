package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/order/comm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "Deletes an order",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := comm.NewOrderSvcClient(conn)
			resp, err := cli.Delete(context.TODO(), &comm.DeleteOrderRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling delete: %v", err)
			}

			log.Printf(
				"deleted order:\n\tid: %s\n\tcode: %s\n\tdescription: %s\n\tname: %s",
				resp.Order.Id,
				resp.Order.Code,
				resp.Order.Description,
				resp.Order.Name)
		},
	}
)

func init() {
	cmdDelete.PersistentFlags().StringVarP(&id, "id", "i", "", "order's id")
}
