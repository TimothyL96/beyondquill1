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
	million100          = 100000000 // 5,761,455 - 55
	million10           = 10000000  // 664,579 - 55
	million1            = 1000000   // 78,498 - 55
	thousand100         = 100000    // 9,592 - 43
	thousand10          = 10000     // 1,229 - 33
	thousand            = 1000      // 168 - 25
	hundred             = 100       // 25 prime circular 13
	circularPrimeFilter = "865420"
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
	// Create a slice with size of half of input for marking numbers
	storeInt = make([]bool, input>>1)

	// Add 2
	primeList = append(primeList, 2)

	// Iterate through only odd numbers till input
	for i := 3; i <= input; i += 2 {
		if storeInt[i>>1] == false {
			// Store current number as prime number
			primeList = append(primeList, i)

			// Iterate multiple of i and mark these numbers as non prime numbers
			for j := i * i; j <= input; j += i << 1 {
				storeInt[j>>1] = true
			}
		}
	}

}

func countCircularPrime() {
	// Create a map to quickly identify which prime number is checked
	circularPrime := make(map[int]struct{})

	// Loop through all prime numbers
	for i := len(primeList) - 1; i >= 0; i-- {
		// Get if current prime number exists in the circularPrime map
		_, ok := circularPrime[primeList[i]]

		// Process only if this prime number is unprocessed (Not in the circularPrime map)
		if !ok {
			// Convert prime number to string for filtering
			inputStr := strconv.Itoa(primeList[i])

			// Skip checking for circular prime number if the number contains any number "024568" and the number not 2 or 5.
			//
			// As it will not be a circular prime number if rotated an even number to the last digit.
			if strings.ContainsAny(inputStr, circularPrimeFilter) && primeList[i] != 2 && primeList[i] != 5 {
				continue
			}

			// Get all possible combination
			nrOfPossibleRotation := len(inputStr)

			// Add current prime number to circularPrime to mark it as visited
			circularPrime[primeList[i]] = struct{}{}

			// Iterate through each possible combination, except the original number itself
			for j, nrOfCombination := 0, nrOfPossibleRotation-1; j < nrOfCombination; j++ {
				// Rotate the prime number
				inputStr = inputStr[1:] + inputStr[:1]

				// Convert the rotated string back to number
				combinationInt, _ := strconv.Atoi(inputStr)

				// Mark rotated number in circular prime as visited
				circularPrime[combinationInt] = struct{}{}

				// If rotated number is same as original number, minus 1 combination
				// Ex: 11 rotate to 11
				if combinationInt == primeList[i] {
					nrOfPossibleRotation--
				}

				// Break if rotated number is not a prime number
				if storeInt[combinationInt>>1] {
					nrOfPossibleRotation = 0
					break
				}
			}

			// Add possible rotation to total number of circular prime numbers
			totalCircularPrime += nrOfPossibleRotation
		}
	}
}
