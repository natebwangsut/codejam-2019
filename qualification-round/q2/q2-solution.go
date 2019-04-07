package main

import (
	"fmt"
)

func main() {

	// Scan first line
	var input int
	var gridTemp int
	var temp string
	fmt.Scanln(&input)

	grid := []int{}
	array := []string{}
	for i := 0; i < input; i++ {
		fmt.Scanln(&temp)
		grid = append(grid, gridTemp)
		fmt.Scanln(&temp)
		array = append(array, temp)
	}

	// Check for input
	for i := range array {
		answer := mirrorLydiaSolution(array[i])
		fmt.Printf("Case #%d: %s\n", i+1, answer)
	}
}

func mirrorLydiaSolution(lydia string) string {
	solution := ""
	buffer := ""
	moveToMake := ""
	for i, c := range lydia {
		if moveToMake != mirrorMove(c) {
			solution += buffer
			buffer = ""
		}
		moveToMake = mirrorMove(c)
		buffer += moveToMake
		if i == len(lydia)-1 {
			solution += buffer
		}
	}
	return solution
}

func mirrorMove(move rune) string {
	if move == 'S' {
		return "E"
	} else if move == 'E' {
		return "S"
	} else {
		return " "
	}
}
