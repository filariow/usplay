package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/report/comm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdList = &cobra.Command{
		Use:   "list",
		Short: "Lists all the reports",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := comm.NewReportSvcClient(conn)
			resp, err := cli.List(context.TODO(), &comm.ListReportsRequest{})
			if err != nil {
				log.Fatalf("error calling list: %v", err)
			}

			for _, v := range resp.GetReports() {
				log.Printf(
					"list report:\n\tid: %s\n\tname: %s",
					v.Id,
					v.Name)
			}
		},
	}
)

func init() {
	cmdList.PersistentFlags().StringVarP(&id, "id", "i", "", "report's id")
}
