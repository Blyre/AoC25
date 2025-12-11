package main

import (
	"aoc25/utils"
	"fmt"
	"strconv"
	"strings"
)

func extractRangesAndValues(content string) ([][2]int, []int) {
	// Iterate over each line of the content, the ranges are in a min-max format and the values are seperated by a blank line
	var ranges [][2]int
	var values []int

	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			min, _ := strconv.Atoi(parts[0])
			max, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, [2]int{min, max})
		} else if line != "" {
			val, _ := strconv.Atoi(line)
			values = append(values, val)
		}
	}
	return ranges, values
}

func optimizeRanges(ranges [][2]int) [][2]int {
	// Sort and merge overlapping ranges
	//fmt.Println("Previous range:", ranges)
	if len(ranges) == 0 {
		return ranges
	}

	// First, sort the ranges by their starting values
	for i := 0; i < len(ranges)-1; i++ {
		for j := 0; j < len(ranges)-i-1; j++ {
			if ranges[j][0] > ranges[j+1][0] {
				ranges[j], ranges[j+1] = ranges[j+1], ranges[j]
			}
		}
	}
	merged := [][2]int{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := merged[len(merged)-1]
		current := ranges[i]
		if current[0] <= last[1] {
			// Overlapping ranges, merge them
			if current[1] > last[1] {
				merged[len(merged)-1][1] = current[1]
			}
		} else {
			merged = append(merged, current)
		}
	}
	//fmt.Println("Optimized range:", merged)
	return merged
}

func checkValueInRanges(value int, ranges [][2]int) bool {
	for _, r := range ranges {
		if value >= r[0] && value <= r[1] {
			return true
		}
	}
	return false
}

func countValidValuesInRanges(ranges [][2]int) int {
	count := 0
	for _, r := range ranges {
		count += r[1] - r[0] + 1
	}
	return count
}

func main() {
	content := utils.OpenAndReadFile("input.txt")

	// Precess content and ex
	ranges, values := extractRangesAndValues(content)
	//fmt.Println("Ranges:", ranges)
	//fmt.Println("Values:", values)

	optimized_ranges := optimizeRanges(ranges)

	count := 0
	for _, val := range values {
		if checkValueInRanges(val, optimized_ranges) {
			count++
		}
	}
	fmt.Println("Total valid values:", count)
	fmt.Println("Total valid range count:", countValidValuesInRanges(optimized_ranges))
}
