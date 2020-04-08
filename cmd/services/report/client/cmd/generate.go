package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/report/comm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdGenerate = &cobra.Command{
		Use:   "generate",
		Short: "Generates a new report",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := comm.NewReportSvcClient(conn)
			resp, err := cli.Generate(context.TODO(), &comm.GenerateReportRequest{Name: name})
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
	cmdGenerate.PersistentFlags().StringVarP(&code, "code", "c", "", "report's code")
	cmdGenerate.PersistentFlags().StringVarP(&desc, "description", "d", "", "report's description")
	cmdGenerate.PersistentFlags().StringVarP(&name, "name", "n", "", "report's name")
}
