package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/FrancescoIlario/usplay/pkg/services/reportcomm"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdCreate = &cobra.Command{
		Use:   "create",
		Short: "Creates a new report",
		Run: func(cmd *cobra.Command, args []string) {
			fromTimestamp, err := parseDateTime(from)
			if err != nil {
				log.Fatalf(`"from" field is invalid: %v`, err)
			}
			toTimestamp, err := parseDateTime(to)
			if err != nil {
				log.Fatalf(`"to" field is invalid: %v`, err)
			}

			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := reportcomm.NewReportSvcClient(conn)
			resp, err := cli.Create(context.TODO(), &reportcomm.CreateReportRequest{
				Name: name,
				Period: &reportcomm.Interval{
					From: fromTimestamp,
					To:   toTimestamp,
				},
			})
			if err != nil {
				log.Fatalf("error calling generate: %v", err)
			}

			log.Printf("generated new report with id: %s", resp.GetId())
		},
	}

	name string
	from string
	to   string
)

const (
	format      = "dd/MM/yyyy-hh:mm:ss"
	parseFormat = "02/01/2006-15:04:05"
)

func init() {
	cmdCreate.Flags().StringVarP(&name, "name", "n", "", "report's name")
	cmdCreate.MarkFlagRequired("name")
	cmdCreate.Flags().StringVarP(&from, "from", "F", "", "Interval's From")
	cmdCreate.MarkFlagRequired("from")
	cmdCreate.Flags().StringVarP(&to, "to", "T", "", "Interval's To")
	cmdCreate.MarkFlagRequired("to")
}

func parseDateTime(value string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(parseFormat, value)
	if err != nil {
		return nil, err
	}
	ts, err := ptypes.TimestampProto(t)
	return ts, err
}

func parseInterval(from, to string) (*reportcomm.Interval, error) {
	fromTimestamp, err := parseDateTime(from)
	if err != nil {
		return nil, fmt.Errorf(`"from" field is invalid: %v`, err)
	}
	toTimestamp, err := parseDateTime(to)
	if err != nil {
		return nil, fmt.Errorf(`"to" field is invalid: %v`, err)
	}
	return &reportcomm.Interval{
		From: fromTimestamp,
		To:   toTimestamp,
	}, nil
}
