package main

import (
	"aoc25/utils"
	"math"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Edge struct {
	x1, y1, x2, y2 int
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

func solveDay1(pairs [][2]Coord) int {
	maxArea := 0
	for _, pair := range pairs {
		area := calcArea(pair[0], pair[1])
		//fmt.Println("Calculating pair,", pair, "Area is:", area)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

func solveDay2(coords []Coord) int {
	if len(coords) < 2 {
		return 0
	}

	var result int = 0

	// Create edges between consecutive coordinates
	edges := []Edge{}
	for i := 0; i < len(coords)-1; i++ {
		edges = append(edges, Edge{
			x1: coords[i].x,
			y1: coords[i].y,
			x2: coords[i+1].x,
			y2: coords[i+1].y,
		})
	}

	// Close the polygon (connect last to first)
	edges = append(edges, Edge{
		x1: coords[len(coords)-1].x,
		y1: coords[len(coords)-1].y,
		x2: coords[0].x,
		y2: coords[0].y,
	})

	// Function to check if a rectangle intersects with any edge
	intersections := func(minX, minY, maxX, maxY int) bool {
		for _, edge := range edges {
			iMinX := min(edge.x1, edge.x2)
			iMaxX := max(edge.x1, edge.x2)
			iMinY := min(edge.y1, edge.y2)
			iMaxY := max(edge.y1, edge.y2)
			if minX < iMaxX && maxX > iMinX && minY < iMaxY && maxY > iMinY {
				return true
			}
		}
		return false
	}

	// Check all pairs of coordinates for maximum non-intersecting rectangle
	for i := 0; i < len(coords)-1; i++ {
		for j := i + 1; j < len(coords); j++ {
			fromCoord := coords[i]
			toCoord := coords[j]
			minX := min(fromCoord.x, toCoord.x)
			maxX := max(fromCoord.x, toCoord.x)
			minY := min(fromCoord.y, toCoord.y)
			maxY := max(fromCoord.y, toCoord.y)

			if !intersections(minX, minY, maxX, maxY) {
				area := calcArea(fromCoord, toCoord)
				if area > result {
					result = area
				}
			}
		}
	}

	return result
}

func main() {
	contents := utils.OpenAndReadFile("input.txt")
	coords := parseCoords(contents)

	// Part 1
	pairs := createPairCoords(coords)
	maxArea := solveDay1(pairs)
	println("Part 1 - Max area:", maxArea)

	// Part 2
	maxNonIntersectingArea := solveDay2(coords)
	println("Part 2 - Max non-intersecting area:", maxNonIntersectingArea)
}
