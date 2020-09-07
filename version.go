package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	version = "0.0.1"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printVersion() {
	fmt.Println("Tpresent version ", version)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Tpresent",
	Long:  "Print the version number of Tpresent",
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}
