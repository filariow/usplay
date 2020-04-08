package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/order/comm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdRead = &cobra.Command{
		Use:   "read",
		Short: "Reads an order",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := comm.NewOrderSvcClient(conn)
			resp, err := cli.Read(context.TODO(), &comm.ReadOrderRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling read: %v", err)
			}

			log.Printf(
				"read order:\n\tid: %s\n\tcode: %s\n\tdescription: %s\n\tname: %s",
				resp.Order.Id,
				resp.Order.Code,
				resp.Order.Description,
				resp.Order.Name)
		},
	}

	id string
)

func init() {
	cmdRead.PersistentFlags().StringVarP(&id, "id", "i", "", "order's id")
}
