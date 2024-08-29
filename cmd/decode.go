package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wangle201210/dt/sdk/decode"
)

var decodeKey string
var decodeData string

var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "解密一些数据",
	Long:  `解密一些数据`,
	Run: func(cmd *cobra.Command, args []string) {
		printRes(decode.DoDecrypt(decodeKey, decodeData))
	},
}

func initDecode() {
	decodeCmd.Flags().StringVarP(&decodeKey, "key", "k", "", "解密使用的key")
	decodeCmd.Flags().StringVarP(&decodeData, "data", "d", "", "需要解密的数据")
	rootCmd.AddCommand(decodeCmd)
}
