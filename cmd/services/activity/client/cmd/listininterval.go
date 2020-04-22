package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/FrancescoIlario/usplay/pkg/services/activitycomm"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdListInInterval = &cobra.Command{
		Use:   "list-interval",
		Short: "Lists all the activities",
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

			cli := activitycomm.NewActivitySvcClient(conn)
			resp, err := cli.ListInInterval(context.TODO(),
				&activitycomm.ListInIntervalActivitiesRequest{
					From: fromTimestamp,
					To:   toTimestamp,
				})
			if err != nil {
				log.Fatalf("error calling list: %v", err)
			}

			for _, v := range resp.GetActivities() {
				log.Printf(`list activity:
	tid: %s
	code: %s
	description: %s
	name: %s`, v.Id, v.Code, v.Description, v.Name)
			}
		},
	}

	from string
	to   string
)

const (
	format      = "dd/MM/yyyy-hh:mm:ss"
	parseFormat = "02/01/2006-15:04:05"
)

func init() {
	cmdListInInterval.Flags().StringVarP(&to, "to", "T", "", fmt.Sprintf(`"to" time of activity in format %s`, format))
	cmdListInInterval.Flags().StringVarP(&from, "from", "F", "", fmt.Sprintf(`"from" time of activity in format %s`, format))
}

func parseDateTime(value string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(parseFormat, value)
	if err != nil {
		return nil, err
	}
	ts, err := ptypes.TimestampProto(t)
	return ts, err
}
