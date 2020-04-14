package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdDelete = &cobra.Command{
		Use:   "delete",
		Short: "Deletes an report",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := reportcomm.NewReportSvcClient(conn)
			resp, err := cli.Delete(context.TODO(), &reportcomm.DeleteReportRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling delete: %v", err)
			}

			log.Printf(
				"deleted report:\n\tid: %s\n\tname: %s",
				resp.Report.Id,
				resp.Report.Name)
		},
	}
)

func init() {
	cmdDelete.PersistentFlags().StringVarP(&id, "id", "i", "", "report's id")
}
