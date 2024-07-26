package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wangle201210/dt/sdk/rand"
)

var randNum bool
var randLength int

var randCmd = &cobra.Command{
	Use:   "rand",
	Short: "生成随机串",
	Long:  `生成随机串`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case randNum:
			printRes(rand.GenerateRandomNumber(randLength), nil)
		default:
			printRes(rand.GenerateRandomString(randLength), nil)
		}
	},
}

func initRand() {
	randCmd.Flags().BoolVarP(&randNum, "num", "n", false, "随机生成数字")
	randCmd.Flags().IntVarP(&randLength, "len", "l", 1, "生成的长度(默认为1)")

	rootCmd.AddCommand(randCmd)
}
