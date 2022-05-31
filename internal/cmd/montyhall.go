package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version string

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "montyhall",
		Short: "This program simulates the Monty Hall problem and calculates the percentage of correct answer",
	}

	rootCmd.AddCommand(NewPlayCmd())

	return rootCmd
}

func HandleCmdErr(err error) {
	os.Exit(1)
}
