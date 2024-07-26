package rand

import "testing"

var n = 2

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateRandomNumber(n)
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateRandomNumber2(n)
	}
}

func Benchmark3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomNumber(n)
	}
}
