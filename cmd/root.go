package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is root command
var rootCmd = &cobra.Command{
	Use:   "circle",
	Short: "uu-circle20のためのコマンド",
	Long:  `uu-circle20のためのコマンド`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("aaaa")
	},
}

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
