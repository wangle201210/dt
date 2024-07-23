package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "dt",
	Short: "data-transfer 转换各种格式的数据",
	Long: `使用dt，把数据从一种格式转化为另一种数据格式
比如把string转为md5、date和时间戳互转`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use dt -h or --help for help.")
	},
}

func initAll() {
	initVersion()
	initMd5()
	initTime()
	initBase64()
}

func Execute() {
	initAll()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printRes(data any, err error) {
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("%+v\n", data)
}
