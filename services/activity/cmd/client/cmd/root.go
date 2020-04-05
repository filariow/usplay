package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "usactivity_cli",
		Short: "Client for the uSPlay project's Activity service",
		Long:  `Command Line gRPC Client for the uSPlay project's Activity microservice `,
	}

	target string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "localhost:8080", "the address of the target server")
	rootCmd.AddCommand(cmdCreate)
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
