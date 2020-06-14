package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/datext"
	"github.com/FrancescoIlario/usplay/pkg/services/reportgrpc"
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

			cli := reportgrpc.NewReportSvcClient(conn)
			resp, err := cli.Read(context.TODO(), &reportgrpc.ReadReportRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling read: %v", err)
			}

			printReport(resp.GetReport())
		},
	}

	id string
)

func init() {
	cmdRead.Flags().StringVarP(&id, "id", "i", "", "report's id")
}

func printReport(report *reportgrpc.Report) {
	if report == nil {
		fmt.Println("Empty report returned")
		return
	}

	from, to := datext.ExtractDatesStr(report.GetPeriod())

	fmt.Printf(`Report %s - %s
From: %s
To: %s
`, report.GetId(), report.GetName(), from, to)

	if activities := report.GetActivities(); len(activities) > 0 {
		fmt.Printf("Printing Activities (%v)\n", len(activities))

		for _, act := range report.GetActivities() {
			from, to := datext.ExtractDatesStr(act.GetPeriod())

			fmt.Printf("%s | %v | %s | %s | %s |\n",
				act.Id,
				act.GetActType().GetCode(),
				from,
				to,
				act.GetOrder().GetName(),
			)
		}
	}
}
