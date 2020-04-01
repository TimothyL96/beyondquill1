package main

import (
	"testing"
)

func BenchmarkFindNrOfCircularPrimeNumber10m(b *testing.B) {
	benchmarkFindNrOfCircularPrimeNumber(10000000, b)
}

// func BenchmarkFindNrOfCircularPrimeNumber1000(b *testing.B) {
// 	benchmarkFindNrOfCircularPrimeNumber(1000, b)
// }

func benchmarkFindNrOfCircularPrimeNumber(j int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		findNrOfCircularPrimeNumber(j)
	}
}
