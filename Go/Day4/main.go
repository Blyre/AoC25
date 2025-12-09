package main

import (
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

func getMarkedPositions(content [][]byte, radius int, threshold int) (int, [][]byte) {
	// Convert strings to byte slices for mutability
	modifiedGrid := make([][]byte, len(content))
	for i, line := range content {
		modifiedGrid[i] = make([]byte, len(line))
		copy(modifiedGrid[i], line)
	}

	validCount := 0
	foundCount := 0
	testCount := 0
	// Iterate through each character in the 2D grid
	for i, line := range content {
		for j, char := range line {
			// When we find a '@', check its surrounding area
			if char == '@' {
				testCount++
				foundCount = 0
				// Check surrounding area within the given radius
				for k := i - radius; k <= i+radius; k++ {
					for l := j - radius; l <= j+radius; l++ {
						// Ensure we don't go out of bounds
						if k >= 0 && k < len(content) && l >= 0 && l < len(content[k]) && (k != i || l != j) {
							// Check if the surrounding character is also '@'
							if content[k][l] == '@' {
								// Increment found count if another '@' is found
								foundCount++
							}
						}
					}
				}
				if foundCount < threshold {
					modifiedGrid[i][j] = 'X'
					validCount++
				}
			}
		}
	}
	print("The modified modifiedGrid is:\n")
	for _, line := range modifiedGrid {
		println(string(line))
	}
	println("Tested positions:", testCount)
	return validCount, modifiedGrid
}

func main() {
	content := openAndReadFile()
	lines := strings.Split(content, "\n")

	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	total := 0
	lastTotal := -1
	removed := int(0)
	for lastTotal != total {
		lastGrid := grid
		lastTotal = total
		removed, grid = getMarkedPositions(lastGrid, 1, 4)
		total += removed
	}
	println("Total marked positions:", total)
}
