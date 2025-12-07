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
				repeat := true
				for j := 1; j < len(numStr); j++ {
					if numStr[j] != numStr[0] {
						repeat = false
						break
					}
				}
				if repeat {
					fmt.Println("Repeat pattern found:", i)
				}
			}
		}
	}
}
