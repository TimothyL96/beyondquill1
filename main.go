package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	// Store all odd numbers for marking
	storeInt []bool

	// Store only prime numbers
	primeList []int

	// Store result of total circular prime numbers
	totalCircularPrimeNumbers int
)

const (
	million100          = 100000000 // 5,761,455 - 55
	million10           = 10000000  // 664,579 - 55
	million1            = 1000000   // 78,498 - 55
	thousand100         = 100000    // 9,592 - 43
	thousand10          = 10000     // 1,229 - 33
	thousand            = 1000      // 168 - 25
	hundred             = 100       // 25 prime - 13 circular
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

	// Print the results
	fmt.Println("Total duration:", totalDuration.String())
	fmt.Println("Max number is: ", maxNumber)
	fmt.Println("Number of prime number is: ", len(primeList))
	fmt.Println("Number of circular prime number is ", totalCircularPrimeNumbers)
}

func findPrimeNumbers(input int) {
	// Create a slice with size of half of input for marking numbers
	storeInt = make([]bool, input>>1)

	// Add 2
	primeList = append(primeList, 2)

	// Iterate through only odd numbers till value of input reached
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

// *****
// Re-use storeInt to mark visited circular prime numbers
// *****
func countCircularPrime() {
	// Loop through all prime numbers
	for i := len(primeList) - 1; i >= 0; i-- {

		// Skip if number is marked in storeInt or the number is not 2.
		// Second condition checking is because 2>>1 will get 1 instead of 0
		if storeInt[primeList[i]>>1] && primeList[i] != 2 {
			continue
		}

		// Convert prime number to string for filtering
		inputStr := strconv.Itoa(primeList[i])

		// Skip checking for circular prime number if the number contains any number "024568" and the number is not 2 or 5.
		// As it will not be a circular prime number if rotated the even number to the last digit.
		if strings.ContainsAny(inputStr, circularPrimeFilter) &&
			primeList[i] != 2 && primeList[i] != 5 {
			continue
		}

		// Get all possible combinations
		nrOfPossibleRotation := len(inputStr)

		// Iterate through each possible combination, except the original number itself
		for j, nrOfCombination := 0, nrOfPossibleRotation-1; j < nrOfCombination; j++ {
			// Rotate the prime number
			inputStr = inputStr[1:] + inputStr[:1]

			// Convert the rotated string back to number
			combinationInt, _ := strconv.Atoi(inputStr)

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

			// Mark rotated number in storeInt as visited
			storeInt[combinationInt>>1] = true
		}

		// Mark current prime number in storeInt as visited
		storeInt[primeList[i]>>1] = true

		// Add total possible rotation to total number of circular prime numbers
		totalCircularPrimeNumbers += nrOfPossibleRotation
	}
}
