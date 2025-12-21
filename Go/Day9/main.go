package main

import (
	"aoc25/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

func parseCoords(data string) []Coord {
	coords := []Coord{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		coords = append(coords, Coord{x, y})
	}
	return coords
}

func createPairCoords(coords []Coord) [][2]Coord {
	pairs := [][2]Coord{}
	for i := 0; i < len(coords); i++ {
		for j := i + 1; j < len(coords); j++ {
			pairs = append(pairs, [2]Coord{coords[i], coords[j]})
		}
	}
	return pairs
}

func calcArea(c1, c2 Coord) int {
	width := math.Abs(float64(c1.x-c2.x)) + 1
	height := math.Abs(float64(c1.y-c2.y)) + 1
	return int(width * height)
}

func main() {
	contents := utils.OpenAndReadFile("input.txt")
	coords := parseCoords(contents)
	pairs := createPairCoords(coords)
	maxArea := 0
	for _, pair := range pairs {
		area := calcArea(pair[0], pair[1])
		fmt.Println("Calculating pair,", pair, "Area is:", area)
		if area > maxArea {
			maxArea = area
		}
	}
	println("Max area:", maxArea)
}
