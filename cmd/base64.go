package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wangle201210/dt/sdk/base64"
)

var base64Img string
var base64ImgPath string
var base64EncodeStr string
var base64DecodeStr string

var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "base64 处理",
	Long:  `base64 处理`,
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case base64ImgPath != "":
			printRes(base64.ImgToBase64(base64ImgPath))
		case base64Img != "":
			printRes(base64.StrToImg(base64Img))
		case base64EncodeStr != "":
			printRes(base64.Encode(base64EncodeStr), nil)
		case base64DecodeStr != "":
			printRes(base64.Decode(base64DecodeStr))
		default:
			fmt.Println("Use dt base64 -h or --help for help.")
		}
	},
}

func initBase64() {
	base64Cmd.Flags().StringVarP(&base64Img, "imgData", "i", "", "用来转图片的base64数据")
	base64Cmd.Flags().StringVarP(&base64ImgPath, "img path", "p", "", "用来转base64的图片地址")
	base64Cmd.Flags().StringVarP(&base64EncodeStr, "encodeStr", "e", "", "用来转base64的字符串")
	base64Cmd.Flags().StringVarP(&base64DecodeStr, "decodeStr", "d", "", "用来解码成字符串的base64")
	rootCmd.AddCommand(base64Cmd)
}
