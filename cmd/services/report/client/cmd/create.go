package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdCreate = &cobra.Command{
		Use:   "generate",
		Short: "Creates a new report",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := reportcomm.NewReportSvcClient(conn)
			resp, err := cli.Create(context.TODO(), &reportcomm.CreateReportRequest{Name: name})
			if err != nil {
				log.Fatalf("error calling generate: %v", err)
			}

			log.Printf("generated new report with id: %s", resp.GetId())
		},
	}

	code string
	desc string
	name string
)

func init() {
	cmdCreate.Flags().StringVarP(&code, "code", "c", "", "report's code")
	cmdCreate.Flags().StringVarP(&desc, "description", "d", "", "report's description")
	cmdCreate.Flags().StringVarP(&name, "name", "n", "", "report's name")
}
