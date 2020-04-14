// +build windows

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var storeInt []bool
var prime = make(map[uint32]struct{})
var totalCircularPrime int

const (
	million100  = 100000000 // 5,761,455
	million10   = 10000000  // 664,579
	million1    = 1000000   // 78,498
	thousand100 = 100000    // 9,592
	thousand10  = 10000     // 1,229
	thousand    = 1000      // 168
	hundred     = 100       // 25 prime
	evenFilter  = "865420"
)

func main() {
	const maxNumber = million10
	var totalDuration time.Duration

	t := time.Now()
	findPrimeNumbers(maxNumber)
	countCircularPrime()
	totalDuration = time.Since(t)

	fmt.Println("Total duration:", totalDuration)
	fmt.Println("Max number is: ", maxNumber)
	fmt.Println("Number of prime number is: ", len(prime))
	fmt.Println("Number of circular prime number is ", totalCircularPrime)
}

func findPrimeNumbers(input uint32) {
	storeInt = make([]bool, input>>1)

	// Add 2
	storeInt[0] = true
	prime[2] = struct{}{}

	for i := uint32(3); i <= input; i += 2 {
		if storeInt[i>>1] == false {
			prime[i] = struct{}{}

			for j := i * i; j <= input; j += i << 1 {
				storeInt[j>>1] = true
			}
		}
	}
}

func countCircularPrime() {
	circular := make(map[uint32]struct{})

	for input := range prime {
		_, ok := circular[input]

		if !ok {
			inputStr := itoa(input)
			if strings.ContainsAny(inputStr, evenFilter) && input != 2 && input != 5 {
				continue
			}

			nrOfCircularPrime := len(inputStr)
			circular[input] = struct{}{}

			for i, nrOfCombination := 0, nrOfCircularPrime-1; i < nrOfCombination; i++ {
				// Create nrOfCombination
				inputStr = inputStr[1:] + inputStr[:1]
				combinationInt := atoi(inputStr)
				circular[combinationInt] = struct{}{}

				if combinationInt == input {
					nrOfCircularPrime--
				}

				// Check if its prime number
				if _, ok := prime[combinationInt]; !ok {
					nrOfCircularPrime = 0
					break
				}
			}

			totalCircularPrime += nrOfCircularPrime
		}
	}
}

func atoi(s string) uint32 {
	// Fast path for small integers that fit int type.
	var n uint32
	for _, ch := range []byte(s) {
		ch -= '0'
		n = n*10 + uint32(ch)
	}

	return n
}

func itoa(i uint32) string {
	return strconv.FormatInt(int64(i), 10)
}
