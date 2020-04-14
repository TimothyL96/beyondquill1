// +build windows

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var storeInt []bool
var primeList []int
var totalCircularPrime int

const (
	million100  = 100000000 // 5,761,455 - 55
	million10   = 10000000  // 664,579 - 55
	million1    = 1000000   // 78,498 - 55
	thousand100 = 100000    // 9,592 - 43
	thousand10  = 10000     // 1,229 - 33
	thousand    = 1000      // 168 - 25
	hundred     = 100       // 25 prime circular 13
	evenFilter  = "865420"
)

func main() {
	// Set the input number size
	const maxNumber = million10

	// Start the timer
	var totalDuration time.Duration
	t := time.Now()

	// Find all prime numbers and store in primeList
	findPrimeNumbers(maxNumber)

	// Find circular prime numbers from the found prime numbers
	countCircularPrime()

	// Stop the timer
	totalDuration = time.Since(t)

	// Print the result
	fmt.Println("Total duration:", totalDuration.String())
	fmt.Println("Max number is: ", maxNumber)
	fmt.Println("Number of prime number is: ", len(primeList))
	fmt.Println("Number of circular prime number is ", totalCircularPrime)
}

func findPrimeNumbers(input int) {
	storeInt = make([]bool, input>>1)

	// Add 2
	primeList = append(primeList, 2)

	for i := 3; i <= input; i += 2 {
		if storeInt[i>>1] == false {
			primeList = append(primeList, i)

			for j := i * i; j <= input; j += i << 1 {
				storeInt[j>>1] = true
			}
		}
	}

}

func countCircularPrime() {
	// Create a map to quickly identify which prime number is checked
	circular := make(map[int]struct{}, len(primeList))

	// Loop through all prime numbers
	for i := len(primeList) - 1; i >= 0; i-- {
		_, ok := circular[primeList[i]]

		// Process only if this prime number is unprocessed (Not in the circular map)
		if !ok {
			inputStr := itoa(primeList[i])
			if strings.ContainsAny(inputStr, evenFilter) && primeList[i] != 2 && primeList[i] != 5 {
				continue
			}

			nrOfPossibleRotation := len(inputStr)
			circular[primeList[i]] = struct{}{}

			for j, nrOfCombination := 0, nrOfPossibleRotation-1; j < nrOfCombination; j++ {
				// Create nrOfCombination
				inputStr = inputStr[1:] + inputStr[:1]
				combinationInt := atoi(inputStr)
				circular[combinationInt] = struct{}{}

				if combinationInt == primeList[i] {
					nrOfPossibleRotation--
				}

				// Check if its prime number
				if storeInt[combinationInt>>1] {
					nrOfPossibleRotation = 0
					break
				}
			}

			totalCircularPrime += nrOfPossibleRotation
		}
	}
}

func atoi(s string) int {
	// Fast path for small integers that fit int type.
	var n int
	for _, ch := range []byte(s) {
		ch -= '0'
		n = n*10 + int(ch)
	}

	return n
}

func itoa(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
