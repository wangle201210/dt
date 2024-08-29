package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wangle201210/dt/sdk/sha"
)

var shaSrc string
var shaSalt string
var shaSafe bool

var shaCmd = &cobra.Command{
	Use:   "sha",
	Short: "把数据进行sha加密",
	Long:  `对数据进行sha加密`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(shaSrc) == 0 && len(args) > 0 {
			shaSrc = args[0]
		}
		if shaSafe {
			printRes(sha.EncodeSafe(shaSrc, shaSalt), nil)
			return
		}
		printRes(sha.Encode(shaSrc, shaSalt), nil)
	},
}

func initSha() {
	shaCmd.Flags().StringVarP(&shaSrc, "src", "s", "", "需要加密的字符串")
	shaCmd.Flags().StringVarP(&shaSalt, "salt", "t", "", "盐(可以为空)")
	shaCmd.Flags().BoolVarP(&shaSafe, "safe", "f", false, "安全模式sha，防暴力破解(sha(src+salt)+salt)")

	rootCmd.AddCommand(shaCmd)
}
