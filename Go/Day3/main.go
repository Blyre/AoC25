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

func openAndReadFile() string {
	filename := "input.txt"
	data, err := os.ReadFile(filename)
	checkErr(err)
	content := string(data)
	return content
}

func getMaxes(line string) int {
	val0, err := strconv.Atoi(string(line[0]))
	checkErr(err)
	val1, err := strconv.Atoi(string(line[1]))
	max_joltage := val0*10 + val1

	for i := 0; i < len(line); i++ {
		current_val, err := strconv.Atoi(string(line[i]))
		checkErr(err)
		for j := i + 1; j < len(line); j++ {
			next_val, err := strconv.Atoi(string(line[j]))
			checkErr(err)
			new_joltage := current_val*10 + next_val
			if new_joltage > max_joltage {
				max_joltage = new_joltage
			}
		}
	}
	return max_joltage
}

func main() {
	fmt.Println("Starting...")

	content := openAndReadFile()
	// Split content string by new lines resulting in a slice of strings
	lines := strings.Split(content, "\n")

	sum := 0

	for _, line := range lines {
		joltage := getMaxes(line)
		sum += joltage
		fmt.Println("In line:", line, "joltage:", joltage, "partial sum:", sum)
	}

	fmt.Println("Final sum of joltages:", sum)
}
