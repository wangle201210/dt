package rand

import (
	"crypto/rand"
	"math"
	mr "math/rand/v2"
	"strconv"
)

func GenerateRandomString(n int) string {
	const letters = "eddycjyabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	ret := make([]byte, n)
	for i := 0; i < n; {
		b := make([]byte, 1)
		if _, err := rand.Read(b); err != nil {
			panic(err)
		}
		ret[i] = letters[int(b[0])%len(letters)]
		i++
	}
	return string(ret)
}

func GenerateRandomNumber(n int) string {
	if n <= 18 {
		return generateRandomNumber2(n)
	}
	return generateRandomNumber(n)
}

func generateRandomNumber(n int) string {
	if n <= 0 {
		return ""
	}
	digits := make([]byte, n)
	// 第一个不能是0
	digits[0] = byte(mr.IntN(9)+1) + '0'
	for i := 1; i < n; i++ {
		digit := mr.IntN(10)
		digits[i] = byte(digit) + '0'
	}
	return string(digits)
}

func generateRandomNumber2(n int) string {
	if n <= 0 {
		return ""
	}
	minNum := int64(math.Pow10(n - 1))
	maxNum := int64(math.Pow10(n))
	res := mr.Int64N(maxNum-minNum) + minNum
	return strconv.Itoa(int(res))
}
