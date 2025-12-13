package main

import (
	"aoc25/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func parseInput(content string) ([][]int, []string) {
	lines := strings.Split(content, "\n")
	var result [][]int
	var operations []string

	for i, line := range lines {
		var row []int
		toParseLine := strings.Fields(line)

		if i == len(lines)-1 {
			operations = append(operations, toParseLine...)
			break
		}

		for _, value := range toParseLine {
			num, _ := strconv.Atoi(value)
			row = append(row, num)
		}
		result = append(result, row)
	}
	return result, operations
}

func performOperation(matrix [][]int, operations []string) int {
	result := 0

	// Iterate through each column
	for col := 0; col < len(operations); col++ {
		op := operations[col]
		columnProduct := 1
		columnSum := 0

		// Process each row in the current column
		for row := 0; row < len(matrix); row++ {
			if op == "*" {
				columnProduct *= matrix[row][col]
			} else if op == "+" {
				columnSum += matrix[row][col]
			}
		}

		// Add the column result to the total
		if op == "*" {
			result += columnProduct
		} else if op == "+" {
			result += columnSum
		}
	}

	return result
}

// Part 2 challange

func parseSpecialInput(content string) ([]int, []string) {
	var operations []string
	var operations_buffer []int
	var number_buffer []string
	lines := strings.Split(content, "\n")
	operations = strings.Fields(lines[len(lines)-1])
	slices.Reverse(operations)
	for j := len(lines[0]) - 1; j >= 0; j-- {
		if lines[0][j] == ' ' && lines[1][j] == ' ' && lines[2][j] == ' ' && lines[3][j] == ' ' {
			operations_buffer = append(operations_buffer, -1)
			continue
		}
		for i := 0; i < len(lines)-1; i++ {
			if lines[i][j] == ' ' {
				continue
			}
			number_buffer = append(number_buffer, string(lines[i][j]))
		}
		number, err := strconv.Atoi(strings.Join(number_buffer, ""))
		utils.CheckErr(err)
		operations_buffer = append(operations_buffer, number)
		number_buffer = []string{}
	}
	return operations_buffer, operations
}

func performTotalCalculation(operations_buffer []int, operations []string) int {
	result := 0
	var final_buffer []int
	var calculation int
	saved_i := 0
	for _, op := range operations {
		if op == "+" {
			calculation = 0
			for i := saved_i; i < len(operations_buffer); i++ {
				if operations_buffer[i] == -1 {
					saved_i = i + 1
					break
				}
				calculation += operations_buffer[i]
			}
		} else if op == "*" {
			calculation = 1
			for i := saved_i; i < len(operations_buffer); i++ {
				if operations_buffer[i] == -1 {
					saved_i = i + 1
					break
				}
				calculation *= operations_buffer[i]
			}
		}
		final_buffer = append(final_buffer, calculation)
	}
	for _, val := range final_buffer {
		result += val
	}
	return result
}

func main() {
	content := utils.OpenAndReadFile("input.txt")
	parsedContent, operations := parseInput(content)
	fmt.Println(parsedContent)
	fmt.Println(operations)
	fmt.Println(performOperation(parsedContent, operations))
	operations_buffer, operations := parseSpecialInput(content)
	fmt.Println("Part 2:", performTotalCalculation(operations_buffer, operations))
}
