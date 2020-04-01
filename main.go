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
	million10   = 10000000
	million1    = 1000000
	thousand100 = 100000
	thousand10  = 10000
	thousand    = 1000
	hundred     = 100
)

func main() {
	const maxNumber = million10
	fmt.Println(time.Now().Format(time.Kitchen))
	nrOfCircularPrimer := findNrOfCircularPrimeNumber(maxNumber)
	fmt.Println(time.Now().Format(time.Kitchen))

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
			nrOfCircularPrime += countCircularPrime(i)
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
	inputStr := strconv.Itoa(input)
	if len(inputStr) > 1 {
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
