package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdRead = &cobra.Command{
		Use:   "read",
		Short: "Reads an report",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := reportcomm.NewReportSvcClient(conn)
			resp, err := cli.Read(context.TODO(), &reportcomm.ReadReportRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling read: %v", err)
			}

			log.Printf(
				"read report:\n\tid: %s\n\tname: %s",
				resp.Report.Id,
				resp.Report.Name)
		},
	}

	id string
)

func init() {
	cmdRead.PersistentFlags().StringVarP(&id, "id", "i", "", "report's id")
}
