package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func openAndReadFile() string {
	filename := "input.txt"
	data, err := os.ReadFile(filename)
	checkErr(err)
	content := string(data)
	return content
}

func getJoltage(line string, numBanks int) int64 {
	result := int64(0)
	start := 0

	for i := numBanks; i > 0; i-- {
		max := byte('0')
		end := len(line) - i

		for j := start; j <= end; j++ {
			if line[j] > max {
				max = line[j]
				start = j + 1
			}
		}

		result += int64(max-'0') * int64(math.Pow(10.0, float64(i-1)))
	}

	return result
}

func main() {
	fmt.Println("Starting...")

	content := openAndReadFile()
	// Split content string by new lines resulting in a slice of strings
	lines := strings.Split(content, "\n")

	sum := int64(0)

	for _, line := range lines {
		joltage := getJoltage(line, 12)
		sum += joltage
		fmt.Println("In line:", line, "joltage:", joltage, "partial sum:", sum)
	}

	fmt.Println("Final sum of joltages:", sum)
}
