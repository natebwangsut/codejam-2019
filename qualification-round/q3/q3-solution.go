package main

import (
	"fmt"
	"math"
)

func main() {

	// Scan first line
	var input, maxPrime, numChar, temp int
	fmt.Scanf("%d", &input)

	var line [][]int
	var maxPrimeEachLine []int

	// Then scan subsequences line
	for testCase := 0; testCase < input; testCase++ {

		fmt.Scanf("%d %d", &maxPrime, &numChar)
		maxPrimeEachLine = append(maxPrimeEachLine, maxPrime)

		// For each subsequence element, add it into the line
		tempLine := make([]int, 0)
		for j := 0; j < numChar; j++ {
			fmt.Scanf("%d", &temp)
			tempLine = append(tempLine, temp)
		}
		line = append(line, tempLine)
	}

	for testCase := 0; testCase < input; testCase++ {

		// Check for input
		bufferPrimes := sieveOfEratosthenes(maxPrimeEachLine[testCase] + 1)

		answer := []int{}
		currentPrime, nextPrime := 0, 0
		for _, item := range line[testCase] {
			currentPrime, nextPrime = findNextCoprime(nextPrime, item, bufferPrimes)
			answer = append(answer, currentPrime)
		}
		answer = append(answer, nextPrime)

		if in(0, answer) {
			currentPrime, nextPrime := 0, answer[0]
			answer = []int{}
			for _, item := range line[testCase] {
				currentPrime, nextPrime = findNextCoprime(nextPrime, item, bufferPrimes)
				answer = append(answer, currentPrime)
			}
			answer = append(answer, nextPrime)
			// println("NOO")
		}

		// for i := range answer {
		// 	fmt.Printf("%d ", answer[i])
		// }
		// fmt.Println("")

		// Creating character mapping
		characterSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		sortedAnswer := mergeSort(unique(answer))
		mapping := createPrimeMapping(sortedAnswer, characterSet)

		// for _, i := range sortedAnswer {
		// 	fmt.Printf("%d ", i)
		// }
		// fmt.Println(len(sortedAnswer))

		// Print mapping for each answer
		fmt.Printf("Case #%d: ", testCase+1)
		for _, i := range answer {
			fmt.Printf("%c", mapping[i])
		}
		fmt.Printf("\n")
	}
}

////////////////////////////////////////////////////////////////////////////////

func in(element int, array []int) bool {
	for _, n := range array {
		if element == n {
			return true
		}
	}
	return false
}

func createPrimeMapping(primeMap []int, char string) map[int]rune {
	m := make(map[int]rune)
	counter := 0
	for _, p := range primeMap {
		if counter >= len(char) {
			return nil
		}
		m[p] = rune(char[counter])
		counter++
	}
	return m
}

//
func findNextCoprime(start int, target int, buffer []int) (int, int) {

	if start != 0 {
		coprime := target / start
		if isPrime(coprime) {
			return start, coprime
		}
	}

	for _, prime := range buffer {

		// Skip until start
		// if start == 0 {
		// 	println(prime, start, target)
		// }

		// If find coprime
		if target%prime == 0 {
			coprime := target / prime
			if isPrime(prime) && isPrime(coprime) {
				return prime, coprime
			}
		}
	}
	// Error
	return 0, 0
}

//
func merge(a, b []int) []int {
	size, i, j := len(a)+len(b), 0, 0
	result := make([]int, size)
	for k := 0; k < size; k++ {
		switch true {
		case i == len(a):
			//all the elements of a already been judged
			//assuming a and b both are sorted, this happens because
			//some cases will have not equal a and b, so it might
			// be a possibility that one array got finished earlier.
			result[k] = b[j]
			j++
		case j == len(b):
			//alll the elements of a already been judged
			//assuming a nd b both are sorted
			result[k] = a[i]
			i++
		case a[i] > b[j]:
			result[k] = b[j]
			//increment the index of b array because its present index
			//is already appended to the result array
			j++
		case a[i] < b[j]:
			//increment the index of a array because its present index
			//element is already appended to the result array
			result[k] = a[i]
			i++
		case a[i] == b[j]:
			//in case of equality
			result[k] = b[j]
			j++
		}
	}
	return result
}

// Do merge sort
func mergeSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	middle := int(len(numbers) / 2)
	return merge(mergeSort(numbers[middle:]), mergeSort(numbers[:middle]))
}

// Remove duplicate from the list
func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func sieveOfEratosthenes(value int) []int {
	returnList := []int{}
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := 2; i < value; i++ {
		if f[i] == false {
			// fmt.Printf("%v ", i)
			returnList = append(returnList, i)
		}
	}
	return returnList
}
