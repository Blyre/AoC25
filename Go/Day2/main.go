package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Part 1
/* func main() {
	fmt.Println("Starting...")

	// Read file
	filename := "input.txt"
	data, err := os.ReadFile(filename)
	checkErr(err)
	content := string(data)

	// Split content by comma
	lines := strings.Split(content, ",")
	fmt.Println("Split result:")
	fmt.Println(lines)
	sum := 0
	// Go through each line and split the string by "-"
	for _, line := range lines {
		parts := strings.Split(line, "-")
		fmt.Println("Parts:")
		fmt.Println(parts)
		// Convert the first position and last of the parts into integers and run a loop from the first to the last
		// Print each number in the loop
		start, err := strconv.Atoi(parts[0])
		checkErr(err)
		end, err := strconv.Atoi(parts[1])
		checkErr(err)
		for i := start; i <= end; i++ {
			// If the number is a repeat pattern like 11, 22, 1010, print it
			numStr := strconv.Itoa(i)
			if len(numStr) >= 2 && len(numStr)%2 == 0 {
				//repeat := true
				// split string in half and compare both halves
				half := len(numStr) / 2
				firstHalf := numStr[:half]
				secondHalf := numStr[half:]
				if firstHalf == secondHalf {
					fmt.Println("Repeat pattern found:", i)
					sum += i
				}
			}
		}
	}
	fmt.Println("Sum of repeat patterns:", sum)
} */

func findMinCommonMultipleLength(numStr string) []int {
	result := []int{}
	length := len(numStr)
	for i := 1; i <= length/2; i++ {
		if length%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func checkIfRepeatPattern(numStr string, length int) bool {
	pattern := numStr[:length]
	times := len(numStr) / length
	for i := 1; i < times; i++ {
		if numStr[i*length:(i+1)*length] != pattern {
			return false
		}
	}
	return true
}

// Part 2
func main() {
	fmt.Println("Starting...")

	// Read file
	filename := "input.txt"
	data, err := os.ReadFile(filename)
	checkErr(err)
	content := string(data)

	// Split content by comma
	lines := strings.Split(content, ",")
	fmt.Println("Split result:")
	fmt.Println(lines)
	sum := 0
	// Go through each line and split the string by "-"
	for _, line := range lines {
		parts := strings.Split(line, "-")
		fmt.Println("Parts:")
		fmt.Println(parts)
		// Convert the first position and last of the parts into integers and run a loop from the first to the last
		// Print each number in the loop
		start, err := strconv.Atoi(parts[0])
		checkErr(err)
		end, err := strconv.Atoi(parts[1])
		checkErr(err)
		for i := start; i <= end; i++ {
			numStr := strconv.Itoa(i)
			// Find minimum multiple common number of len(numStr)
			minCommonLengths := findMinCommonMultipleLength(numStr)
			for _, length := range minCommonLengths {
				if checkIfRepeatPattern(numStr, length) {
					fmt.Println("Repeat pattern found:", i)
					sum += i
					break
				}
			}
		}
	}
	fmt.Println("Sum of repeat patterns:", sum)
}
