package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const defTarget = "localhost:8080"

var (
	rootCmd = &cobra.Command{
		Use:   "usactivitytype_cli",
		Short: "Client for the uSPlay project's activitytype service",
		Long:  `Command Line gRPC Client for the uSPlay project's activitytype microservice `,
	}

	target string

	code int32
	desc string
	name string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", defTarget, "the address of the target server")
	rootCmd.AddCommand(cmdCreate)
	rootCmd.AddCommand(cmdExist)
	rootCmd.AddCommand(cmdRead)
	rootCmd.AddCommand(cmdUpdate)
	rootCmd.AddCommand(cmdDelete)
	rootCmd.AddCommand(cmdList)
}

// Execute executes the root cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
