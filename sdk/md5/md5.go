package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str, salt string, length int) string {
	h := md5.New()
	h.Write([]byte(str + salt))
	res := hex.EncodeToString(h.Sum(nil))
	if length > 0 && len(res) > length {
		return res[:length]
	}
	return res
}
