package json

import (
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"time"
)

func TomlToJson(tomlData string) (jsonData string, err error) {
	// 解析TOML数据
	var data interface{}
	if _, err = toml.Decode(tomlData, &data); err != nil {
		return
	}

	// 将Go数据结构转换为JSON
	jsonByte, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}
	jsonData = string(jsonByte)
	return
}

func TomlToJsonFile(tomlData string) (filename string, err error) {
	data, err := TomlToJson(tomlData)
	if err != nil {
		return
	}
	// 将解码的数据写入临时文件
	tempFile, err := os.Create(fmt.Sprintf("temp_json_%d.json", time.Now().Unix()))
	if err != nil {
		return
	}
	defer tempFile.Close()
	if _, err = tempFile.Write([]byte(data)); err != nil {
		return
	}
	filename = tempFile.Name()
	return
}

func TomlFileToJsonFile(fromFilename string) (filename string, err error) {
	tomlData, err := os.ReadFile(fromFilename)
	if err != nil {
		return
	}
	return TomlToJsonFile(string(tomlData))
}

func TomlFileToJson(fromFilename string) (res string, err error) {
	tomlData, err := os.ReadFile(fromFilename)
	if err != nil {
		return
	}
	return TomlToJson(string(tomlData))
}
