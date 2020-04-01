package main

import (
	"fmt"
	"strconv"
	"time"
)

var storeInt []bool
var prime = make(map[int]struct{})
var totalCircularPrime int

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
	const maxNumber = million10
	fmt.Println(time.Now().Format(time.StampNano))
	findNrOfCircularPrimeNumber1(maxNumber)
	countCircularPrime1()
	fmt.Println(time.Now().Format(time.StampNano))

	fmt.Println("Max number is: ", maxNumber)
	fmt.Println("Number of prime number is: ", len(prime))
	fmt.Println("Number of circular prime number is ", totalCircularPrime)
}

func findNrOfCircularPrimeNumber1(input int) {
	storeInt = make([]bool, input+1)

	// Add 2
	storeInt[2] = true
	prime[2] = struct{}{}

	for i := 3; i <= input; i += 2 {
		if storeInt[i] == false {
			prime[i] = struct{}{}

			for j := i * i; j <= input; j += i {
				storeInt[j] = true
			}
		}
	}
}

func countCircularPrime1() {
	for input := range prime {
		inputStr := strconv.Itoa(input)
		nrOfCircularPrime := len(inputStr)

		for i, nrOfCombination := 0, nrOfCircularPrime-1; i < nrOfCombination; i++ {
			// Create nrOfCombination
			inputStr = inputStr[1:] + inputStr[:1]
			combinationInt, _ := strconv.Atoi(inputStr)

			if combinationInt == input {
				nrOfCircularPrime--
			}

			// Check if its prime number
			if _, ok := prime[combinationInt]; !ok {
				nrOfCircularPrime = 0
				break
			}
		}

		if nrOfCircularPrime != 0 {
			totalCircularPrime++
		}
	}
}
