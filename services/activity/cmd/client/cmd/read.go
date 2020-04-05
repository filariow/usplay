package cmd

import (
	"context"
	"log"
	"usplay/us-activity/gen"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdRead = &cobra.Command{
		Use:   "read",
		Short: "Reads an activity",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := gen.NewActivitySvcClient(conn)
			resp, err := cli.Read(context.TODO(), &gen.ReadActivityRequest{
				Id: id,
			})
			if err != nil {
				log.Fatalf("error calling read: %v", err)
			}

			log.Printf(
				"read activity:\n\tid: %s\n\tcode: %s\n\tdescription: %s\n\tname: %s",
				resp.Activity.Id,
				resp.Activity.Code,
				resp.Activity.Description,
				resp.Activity.Name)
		},
	}

	id string
)

func init() {
	cmdRead.PersistentFlags().StringVarP(&id, "id", "i", "", "activity's id")
}
