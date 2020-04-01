package main

import (
	"fmt"
	"strconv"
	"time"
)

var allPrimeNumbers = make(map[int]struct{})
var allPrimeNumbersInList = make([]int, 0)

// var visitedNumbers = make(map[int]struct{})

const (
	million100  = 100000000 // 664,579
	million10   = 10000000  // 664,579
	million1    = 1000000   // 78,498
	thousand100 = 100000    // 9,592
	thousand10  = 10000     // 1,229
	thousand    = 1000      // 1,68
	hundred     = 100       // 25 prime
)

func main() {
	const maxNumber = million100
	fmt.Println(time.Now().Format(time.StampNano))
	nrOfCircularPrimer := findNrOfCircularPrimeNumber1(maxNumber)
	fmt.Println(time.Now().Format(time.StampNano))

	fmt.Println("Number of prime number is: ", len(allPrimeNumbers))
	fmt.Println("Number of circular prime number for max number:", maxNumber, " is ", nrOfCircularPrimer)
}

func findNrOfCircularPrimeNumber(input int) int {
	nrOfCircularPrime := 0

	// Add 2 to prime number
	allPrimeNumbers[2] = struct{}{}
	allPrimeNumbersInList = append(allPrimeNumbersInList, 2)
	nrOfCircularPrime++

	for i := 3; i <= input; i = i + 2 {
		if _, ok := allPrimeNumbers[i]; !ok && isPrimeNumber(i) {
			allPrimeNumbers[i] = struct{}{}
			allPrimeNumbersInList = append(allPrimeNumbersInList, i)
			// nrOfCircularPrime += countCircularPrime(i)
		}
	}

	return nrOfCircularPrime
}

func countCircularPrime(input int) int {
	inputStr := strconv.Itoa(input)
	nrOfCircularPrime := len(inputStr)

	// Test all nrOfCombination
	combinationStr := ""

	for i, nrOfCombination := 0, nrOfCircularPrime-1; i < nrOfCombination; i++ {
		// Create nrOfCombination
		combinationStr = inputStr[1:] + inputStr[:1]
		combinationInt, _ := strconv.Atoi(combinationStr)

		if combinationInt == input {
			nrOfCircularPrime--
		}

		// Check if its prime number
		if _, ok := allPrimeNumbers[i]; !ok && combinationInt != input && !isPrimeNumber(combinationInt) {
			return 0
		}

		if _, ok := allPrimeNumbers[i]; !ok {
			allPrimeNumbers[combinationInt] = struct{}{}
			allPrimeNumbersInList = append(allPrimeNumbersInList, combinationInt)
		}
	}

	return nrOfCircularPrime
}

func isPrimeNumber(input int) bool {
	if input >= 10 {
		inputStr := strconv.Itoa(input)
		inputStrLastChar := inputStr[len(inputStr)-1:]

		if inputStrLastChar != "1" &&
			inputStrLastChar != "3" &&
			inputStrLastChar != "7" &&
			inputStrLastChar != "9" {
			return false
		}
	}

	for _, k := range allPrimeNumbersInList {
		if input%k == 0 {
			return false
		}
	}

	return true
}

func findNrOfCircularPrimeNumber1(input int) int {
	nrOfCircularPrime := 0
	storeInt := make([]int, input+1)
	allPrimeNumbers[2] = struct{}{}
	for i := 3; i <= input; i += 2 {
		if storeInt[i] == 0 {
			allPrimeNumbers[i] = struct{}{}

			for j := i * i; j <= input; j += i {
				storeInt[j] = j
			}
		}
	}

	return nrOfCircularPrime
}

func countCircularPrime1(input int) int {
	inputStr := strconv.Itoa(input)
	nrOfCircularPrime := len(inputStr)

	// Test all nrOfCombination
	combinationStr := ""

	for i, nrOfCombination := 0, nrOfCircularPrime-1; i < nrOfCombination; i++ {
		// Create nrOfCombination
		combinationStr = inputStr[1:] + inputStr[:1]
		combinationInt, _ := strconv.Atoi(combinationStr)

		if combinationInt == input {
			nrOfCircularPrime--
		}

		// Check if its prime number
		if _, ok := allPrimeNumbers[i]; !ok && combinationInt != input && !isPrimeNumber(combinationInt) {
			return 0
		}

		if _, ok := allPrimeNumbers[i]; !ok {
			allPrimeNumbers[combinationInt] = struct{}{}
			allPrimeNumbersInList = append(allPrimeNumbersInList, combinationInt)
		}
	}

	return nrOfCircularPrime
}
