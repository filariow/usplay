package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "usreport_cli",
		Short: "Client for the uSPlay project's Report service",
		Long:  `Command Line gRPC Client for the uSPlay project's Report microservice `,
	}

	target string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "localhost:8080", "the address of the target server")
	rootCmd.AddCommand(cmdCreate)
	rootCmd.AddCommand(cmdRead)
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
