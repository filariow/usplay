package cmd

import (
	"context"
	"fmt"
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
	cmdRead.Flags().StringVarP(&id, "id", "i", "", "report's id")
}

func printReport(report *reportcomm.Report) {
	fmt.Printf(`Report %s - %s
From: %s
To: %s
`, report.GetId(), report.GetName(), report.GetPeriod().GetFrom(),
		report.GetPeriod().GetTo())

	if activities := report.GetActivities(); len(activities) > 0 {
		fmt.Printf("Printing Activities (%v)\n", len(activities))
		for _, act := range report.GetActivities() {
			fmt.Printf(`%s | %v | %s | %s | %s |`,
				act.Id,
				act.GetActType().GetCode(),
				act.GetPeriod().GetFrom(),
				act.GetPeriod().GetTo(),
				act.GetOrder().GetName(),
			)
		}
	}
}
