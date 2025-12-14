package main

import (
	"aoc25/utils"
)

type Coord struct {
	row int
	col int
}

var beams []Coord
var seen map[Coord]bool
var memo map[Coord]int

func findStartingPosition(matrix [][]rune) Coord {
	for i, row := range matrix {
		for j, val := range row {
			if val == 'S' {
				return Coord{i, j}
			}
		}
	}
	return Coord{-1, -1}
}

func addBeam(row, col int) {
	coord := Coord{row, col}
	if !seen[coord] {
		seen[coord] = true
		beams = append(beams, coord)
	}
}

func countBeamSplits(matrix [][]rune) int {
	startingPosition := findStartingPosition(matrix)
	beams = append(beams, Coord{startingPosition.row + 1, startingPosition.col})
	seen = make(map[Coord]bool)
	splits := 0
	for len(beams) > 0 {
		currentBeam := beams[0]
		beams = beams[1:]
		if string(matrix[currentBeam.row][currentBeam.col]) == "." {
			if currentBeam.row == len(matrix)-1 {
				continue
			}
			addBeam(currentBeam.row+1, currentBeam.col)
		} else if string(matrix[currentBeam.row][currentBeam.col]) == "^" {
			splits++
			addBeam(currentBeam.row, currentBeam.col-1)
			addBeam(currentBeam.row, currentBeam.col+1)
		}
	}
	return splits
}

func solve(matrix [][]rune, coord Coord) int {
	// Check if result is already memoized
	if val, exists := memo[coord]; exists {
		return val
	}

	// Compute result
	var result int
	if coord.row >= len(matrix) {
		result = 1
	} else if string(matrix[coord.row][coord.col]) == "." || string(matrix[coord.row][coord.col]) == "S" {
		result = solve(matrix, Coord{coord.row + 1, coord.col})
	} else if string(matrix[coord.row][coord.col]) == "^" {
		result = solve(matrix, Coord{coord.row, coord.col - 1}) + solve(matrix, Coord{coord.row, coord.col + 1})
	} else {
		result = 0
	}

	// Store in memo and return
	memo[coord] = result
	return result
}

func main() {
	content := utils.OpenAndReadFile("input.txt")
	matrix := utils.StringtToMatrix(content)
	splits := countBeamSplits(matrix)
	println("Part 1 - Number of beam splits:", splits)

	memo = make(map[Coord]int)
	startingPosition := findStartingPosition(matrix)
	result := solve(matrix, startingPosition)
	println("Part 2 - Result:", result)
}
