package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Tpresent",
	Short: "Tpresent is a terminal alternative to golang's Present",
	Long: `Tpresent render slides in terminal and is compatible with 
golang's Present slide files`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func main() {
	execute()
}

func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
