package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wangle201210/dt/sdk/json"
	"strings"
)

var jsonTomlStr string
var jsonTomlFile string
var jsonTomlPrint bool
var jsonStructStr string
var jsonStructFile string
var jsonStructPrint bool

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json 相关操作",
	Long:  `json 相关操作`,
	Run: func(cmd *cobra.Command, args []string) {
		notify("json")
	},
}

var jsonTomlCmd = &cobra.Command{
	Use:   "toml",
	Short: "把数据进行json加密",
	Long:  `对数据进行json加密`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if strings.HasSuffix(args[0], ".toml") && len(jsonTomlFile) == 0 {
				jsonTomlFile = args[0]
			} else if len(jsonTomlStr) == 0 {
				jsonTomlStr = args[0]
			}
		}
		switch {
		case jsonTomlStr != "":
			if jsonTomlPrint {
				printRes(json.TomlToJson(jsonTomlStr))
				return
			}
			printRes(json.TomlToJsonFile(jsonTomlStr))
		case jsonTomlFile != "":
			if jsonTomlPrint {
				printRes(json.TomlFileToJson(jsonTomlFile))
				return
			}
			printRes(json.TomlFileToJsonFile(jsonTomlFile))
		default:
			notify("json toml")
		}
	},
}

var jsonStructCmd = &cobra.Command{
	Use:   "struct",
	Short: "把json转 golang struct",
	Long:  `把json转 golang struct`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(jsonStructStr) == 0 {
			jsonStructStr = args[0]
		}
		if len(args) > 0 {
			if strings.HasSuffix(args[0], ".json") && len(jsonStructFile) == 0 {
				jsonStructFile = args[0]
			} else if len(jsonStructStr) == 0 {
				jsonStructStr = args[0]
			}
		}
		switch {
		case jsonStructStr != "":
			if jsonStructPrint {
				printRes(json.ToStructByContent(jsonStructStr))
				return
			}
			printRes(json.ToStructByContentToFile(jsonStructStr))
		case jsonStructFile != "":
			if jsonStructPrint {
				printRes(json.ToStructByFile(jsonStructFile))
				return
			}
			printRes(json.ToStructByFileToFile(jsonStructFile))
		default:
			notify("json struct")
		}
	},
}

func initJson() {
	jsonTomlCmd.Flags().StringVarP(&jsonTomlStr, "toml", "t", "", "需要转为json的toml内容")
	jsonTomlCmd.Flags().StringVarP(&jsonTomlFile, "tomlFile", "f", "", "需要转为json的toml文件地址")
	jsonTomlCmd.Flags().BoolVarP(&jsonTomlPrint, "print", "p", false, "结果在控制台直接输出")

	jsonStructCmd.Flags().StringVarP(&jsonStructStr, "data", "d", "", "json文件内容")
	jsonStructCmd.Flags().StringVarP(&jsonStructFile, "file", "f", "", "json文件路径")
	jsonStructCmd.Flags().BoolVarP(&jsonStructPrint, "print", "p", false, "结果在控制台直接输出")

	jsonCmd.AddCommand(jsonStructCmd)
	jsonCmd.AddCommand(jsonTomlCmd)
	rootCmd.AddCommand(jsonCmd)
}
