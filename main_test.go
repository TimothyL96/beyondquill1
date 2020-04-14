package main

import (
	"testing"
)

func BenchmarkFindPrimeNumbers10m(b *testing.B) {
	benchmarkFindPrimeNumbers(10000000, b)
}

// func BenchmarkFindNrOfCircularPrimeNumber1000(b *testing.B) {
// 	benchmarkFindNrOfCircularPrimeNumber(1000, b)
// }

func benchmarkFindPrimeNumbers(j int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		findPrimeNumbers(j)
	}
}
