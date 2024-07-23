package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version of dt",
	Long:  `All tools have a version, this is dt's"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dt-0.0.1-beta")
	},
}

func initVersion() {
	rootCmd.AddCommand(versionCmd)
}
