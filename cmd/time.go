package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wangle201210/dt/sdk/time"
	"strconv"
)

var timeStr string
var timestamp int64

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式转换",
	Long:  `时间格式转换，如把date转为时间戳`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use dt time -h or --help for help.")
	},
}

func initTime() {
	dateCmd.Flags().StringVarP(&timeStr, "date", "d", "", "日期：格式2006-01-02 15:04:05,如果为空则返回当前时间戳")
	timestampCmd.Flags().Int64VarP(&timestamp, "timestamp", "t", 0, "把时间戳转为2006-01-02 15:04:05格式,如果是-1则返回当前时间")

	timeCmd.AddCommand(dateCmd)
	timeCmd.AddCommand(timestampCmd)

	rootCmd.AddCommand(timeCmd)
}

var dateCmd = &cobra.Command{
	Use:   "date",
	Short: "日期转时间戳",
	Long:  `日期转时间戳`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(timeStr) == 0 && len(args) > 0 {
			timeStr = args[0]
		}
		res, err := time.ToTimestamp(timeStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	},
}

var timestampCmd = &cobra.Command{
	Use:   "ts",
	Short: "时间戳转日期",
	Long:  `时间戳转日期`,
	Run: func(cmd *cobra.Command, args []string) {
		if timestamp == 0 && len(args) > 0 {
			ts, err := strconv.Atoi(args[0])
			if err == nil {
				timestamp = int64(ts)
			}
		}
		res, err := time.ToDate(timestamp)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res)
	},
}
