package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wangle201210/dt/sdk/url"
)

var urlEncodeStr string
var urlDecodeStr string

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "对url进行编解码操作",
	Long:  `对url进行编解码操作`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case urlEncodeStr != "":
			printRes(url.Encode(urlEncodeStr), nil)
		case urlDecodeStr != "":
			printRes(url.Decode(urlDecodeStr))
		default:
			fmt.Println("Use dt url -h or --help for help.")
		}
	},
}

func initUrl() {
	urlCmd.Flags().StringVarP(&urlDecodeStr, "decode", "d", "", "需要解码的url")
	urlCmd.Flags().StringVarP(&urlEncodeStr, "encode", "e", "", "需要编码的url")

	rootCmd.AddCommand(urlCmd)
}
