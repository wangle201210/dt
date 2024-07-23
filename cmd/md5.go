package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wangle201210/dt/sdk/md5"
)

var md5Src string
var md5Salt string
var md5Length int

var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "把数据进行md5加密",
	Long:  `对数据进行md5加密`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(md5Src) == 0 && len(args) > 0 {
			md5Src = args[0]
		}
		printRes(md5.Md5(md5Src, md5Salt, md5Length), nil)
	},
}

func initMd5() {
	md5Cmd.Flags().StringVarP(&md5Src, "src", "s", "", "需要加密的字符串")
	md5Cmd.Flags().StringVarP(&md5Salt, "salt", "t", "", "盐(可以为空)")
	md5Cmd.Flags().IntVarP(&md5Length, "len", "l", 32, "返回的数据长度,默认32位")

	rootCmd.AddCommand(md5Cmd)
}
