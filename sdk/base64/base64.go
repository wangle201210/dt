package base64

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"
)

func Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Decode(str string) (res string, err error) {
	decode, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return
	}
	res = string(decode)
	return
}

func StrToImg(base64Image string) (res string, err error) {
	// 去除可能存在的前缀（如 data:image/png;base64,）
	if strings.HasPrefix(base64Image, "data:image/") {
		commaIndex := strings.Index(base64Image, ",")
		if commaIndex != -1 {
			base64Image = base64Image[commaIndex+1:]
		}
	}

	// 解码 Base64 字符串
	imgData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return
	}

	// 将解码的数据写入临时文件
	tempFile, err := os.Create(fmt.Sprintf("temp_image_%d.png", time.Now().Unix()))
	if err != nil {
		return
	}
	defer tempFile.Close()
	if _, err = tempFile.Write(imgData); err != nil {
		return
	}
	res = tempFile.Name()
	return
}

func ImgToBase64(filePath string) (res string, err error) {
	imageData, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	// 将图片数据编码为 Base64
	res = base64.StdEncoding.EncodeToString(imageData)
	return
}
