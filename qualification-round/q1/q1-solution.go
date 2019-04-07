package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	// Scan first line
	var input, temp int64
	fmt.Scanln(&input)

	array := []int64{}
	for i := int64(0); i < input; i++ {
		fmt.Scanln(&temp)
		array = append(array, temp)
	}

	// Check for input
	for i := range array {
		answer := eliminateFour(array[i])
		fmt.Printf("Case #%d: %d %d\n", i+1, array[i]-answer, answer)
	}

}

func eliminateFour(number int64) int64 {
	four := strconv.FormatInt(number, 10)
	fourLength := len(four)
	answer := int64(0)

	for i, c := range four {
		if string(c) == "4" {
			answer += int64(math.Pow(float64(10), float64(fourLength-i-1)))
		}
	}
	return answer
}
