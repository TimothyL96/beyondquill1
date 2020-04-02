package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var storeInt []bool
var prime = make(map[int]struct{})
var totalCircularPrime int

const (
	million100  = 100000000 // 5,761,455
	million10   = 10000000  // 664,579
	million1    = 1000000   // 78,498
	thousand100 = 100000    // 9,592
	thousand10  = 10000     // 1,229
	thousand    = 1000      // 168
	hundred     = 100       // 25 prime
	evenFilter  = "02468"
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

func findPrimeNumbers(input int) {
	storeInt = make([]bool, input>>1)

	// Add 2
	storeInt[0] = true
	prime[2] = struct{}{}

	for i := 3; i <= input; i += 2 {
		if storeInt[i>>1] == false {
			prime[i] = struct{}{}

			for j := i * i; j <= input; j += i << 1 {
				storeInt[j>>1] = true
			}
		}
	}
}

func countCircularPrime() {
	circular := make(map[int]struct{})

	for input := range prime {
		_, ok := circular[input]

		if !ok {
			inputStr := strconv.Itoa(input)
			if strings.ContainsAny(inputStr, evenFilter) && input != 2 {
				continue
			}

			nrOfCircularPrime := len(inputStr)
			circular[input] = struct{}{}

			for i, nrOfCombination := 0, nrOfCircularPrime-1; i < nrOfCombination; i++ {
				// Create nrOfCombination
				inputStr = inputStr[1:] + inputStr[:1]
				combinationInt, _ := strconv.Atoi(inputStr)
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

// func test() {
// 	v := 1000000
// 	start1 := time.Now()
// 	num1 := countPrimes(v)
// 	elapsed1 := time.Since(start1)
// 	log.Printf("Test %d: %v \t  took %s", v, num1, elapsed1)
//
// 	w := 10000000
// 	start2 := time.Now()
// 	num2 := countPrimes(w)
// 	elapsed2 := time.Since(start2)
// 	log.Printf("Test %d: %v \t  took %s", w, num2, elapsed2)
//
// 	x := 100000000
// 	start3 := time.Now()
// 	num3 := countPrimes(x)
// 	elapsed3 := time.Since(start3)
// 	log.Printf("Test %d: %v \t  took %s", x, num3, elapsed3)
// }
//
// func isCircularPrime(primenr int) bool {
// 	// Skip checking for single digit since they are circular prime
// 	if primenr > 10 {
// 		digitcount := 0
// 		// number := primenr;
// 		rotatednumber := primenr
//
// 		// Get the total digits of the number
// 		for number := primenr; number != 0; number = number / 10 {
// 			remainder := number % 10
// 			// Immediately quit and return false if any of the number is not 1,3,7 or 9
// 			// Since the last digit of the number must be any of these
// 			// Not fulfilling this condition means the number is not circular prime
// 			if remainder != 1 && remainder != 3 && remainder != 7 && remainder != 9 {
// 				return false
// 			}
// 			digitcount++
// 		}
//
// 		// Check if all the rotated numbers are prime number
// 		for i := 0; i < digitcount; i++ {
// 			rotatednumber = RotateNumber(rotatednumber)
//
// 			if !IsPrimeNumber(rotatednumber) {
// 				return false
// 			}
// 		}
// 	}
//
// 	return true
// }
//
// func countPrimes(limit int) int {
// 	// Return if less than 1
// 	if limit <= 1 {
// 		return 0
// 	}
//
// 	// Get the sqrt of the limit
// 	sqrtLimit := int(math.Sqrt(float64(limit)))
//
// 	// Create array
// 	numbers := make([]bool, limit)
//
// 	// Set 1 to prime
// 	numbers[0] = true
// 	numPrimes := 0
//
// 	// Count the number of olds
// 	if limit%2 == 0 {
// 		numPrimes = limit / 2
// 	} else {
// 		numPrimes = (limit + 1) / 2
// 	}
//
// 	// Loop through odd numbers
// 	for i := 3; i <= sqrtLimit; i += 2 {
// 		if !numbers[i] {
// 			for j := i * i; j < limit; j += i * 2 {
// 				if !numbers[j] {
// 					numbers[j] = true
// 					numPrimes -= 1
// 				}
// 			}
// 		}
// 	}
//
// 	return numPrimes
// }
//
// func RotateNumber(nr int) int {
// 	result := nr
// 	// No rotate if the number is in single digit
// 	if nr > 10 {
// 		divisor := 10
// 		remainingnr := nr % divisor
// 		number := nr / divisor
//
// 		multiplier := divisor
// 		tempnumber := number
//
// 		// While loop to find the total digits the number have
// 		for tempnumber/divisor >= 1 {
// 			multiplier = multiplier * divisor
// 			tempnumber = tempnumber / divisor
// 		}
// 		result = number + (remainingnr * multiplier)
// 	}
//
// 	return result
// }
//
// func IsPrimeNumber(nr int) bool {
// 	// Method 1: Custom check, square root the number and see if there is any value
// 	// before square root is the factor of the number, because square root * square root
// 	// equal to the number, any number larger than the square root cannot be the factor of the number
// 	if nr <= 1 {
// 		return false
// 	}
// 	if nr == 2 || nr == 3 || nr == 5 {
// 		return true
// 	}
//
// 	remainder := nr % 10
//
// 	if remainder == 1 || remainder == 3 || remainder == 7 || remainder == 9 {
// 		for i := 3; i <= int(math.Floor(math.Sqrt(float64(nr)))); i += 2 {
// 			if nr%i == 0 {
// 				return false
// 			}
// 		}
// 	} else {
// 		return false
// 	}
//
// 	return true
// }
