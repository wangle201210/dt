package sha

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encode(str, salt string) string {
	// 将盐和原始字符串连接在一起
	saltedString := append([]byte(str), []byte(salt)...)
	// 生成 SHA-256 哈希
	hash := sha256.New()
	hash.Write(saltedString)
	hashBytes := hash.Sum(nil)
	// 将盐和哈希值分别转换为十六进制字符串
	return hex.EncodeToString(hashBytes)
}

func EncodeSafe(str, salt string) string {
	saltedString := append([]byte(str), []byte(salt)...)
	hash := sha256.New()
	hash.Write(saltedString)
	hashBytes := hash.Sum(nil)
	// 二次加密防暴力破解
	hashBytes = append(hashBytes, []byte(salt)...)
	hash2 := sha256.New()
	hash2.Write(hashBytes)
	hashBytes2 := hash2.Sum(nil)
	return hex.EncodeToString(hashBytes2)
}
