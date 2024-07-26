package json

import (
	"fmt"
	"github.com/wangle201210/dt/sdk/json/convert"
	"os"
	"time"
)

func ToStructByFile(file string) (res string, err error) {
	tomlData, err := os.ReadFile(file)
	if err != nil {
		return
	}
	return convert.Generate(string(tomlData), "TestStruct", "main", []string{"json"}, false, true)
}

func ToStructByContent(content string) (res string, err error) {
	return convert.Generate(content, "TestStruct", "main", []string{"json"}, false, true)
}

func ToStructByFileToFile(file string) (res string, err error) {
	data, err := ToStructByFile(file)
	return structToFile(data)
}

func ToStructByContentToFile(content string) (res string, err error) {
	data, err := ToStructByContent(content)
	return structToFile(data)
}

func structToFile(content string) (filename string, err error) {
	// 将解码的数据写入临时文件
	tempFile, err := os.Create(fmt.Sprintf("temp_go_%d.go", time.Now().Unix()))
	if err != nil {
		return
	}
	defer tempFile.Close()
	if _, err = tempFile.Write([]byte(content)); err != nil {
		return
	}
	filename = tempFile.Name()
	return
}
