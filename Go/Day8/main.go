package main

import (
	"aoc25/utils"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Breaker struct {
	x, y, z int
}

type Circuit struct {
	breakers []Breaker
}

type PairWithDistance struct {
	i, j int
	dist float32
}

func parseBreakers(data string) []Breaker {
	breakers := []Breaker{}
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		z, _ := strconv.Atoi(parts[2])
		breakers = append(breakers, Breaker{x, y, z})
	}
	return breakers
}

func distance(b1, b2 Breaker) float32 {
	dx := float32(b1.x - b2.x)
	dy := float32(b1.y - b2.y)
	dz := float32(b1.z - b2.z)
	return float32(math.Sqrt(float64(dx*dx + dy*dy + dz*dz)))
}

func findShortestPairsDistance(breakers []Breaker) []PairWithDistance {
	pairs := []PairWithDistance{}
	for i := 0; i < len(breakers); i++ {
		for j := i + 1; j < len(breakers); j++ {
			distance := distance(breakers[i], breakers[j])
			pairs = append(pairs, PairWithDistance{i, j, distance})
		}
	}

	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].dist < pairs[b].dist
	})

	return pairs
}

func findBreakerInCircuitMap(breaker Breaker, circuits map[int]Circuit) (int, bool) {
	for circuitID, circuit := range circuits {
		for _, b := range circuit.breakers {
			if b == breaker {
				return circuitID, true
			}
		}
	}
	return -1, false
}

/* func addBreakerToCircuit(breaker Breaker, circuits []Circuit, circuitIndex int) []Circuit {

} */

func connectCircuits(pairs []PairWithDistance, breakers []Breaker, size int) map[int]Circuit {
	connected := make([]bool, len(breakers))
	circuits := make(map[int]Circuit)
	fmt.Println("Connecting circuits...", len(connected))
	// Initialize first circuit
	circuits[0] = Circuit{[]Breaker{breakers[pairs[0].i], breakers[pairs[0].j]}}
	connected[pairs[0].i] = true
	connected[pairs[0].j] = true
	nextCircuitID := 1
	i := 1
	for {
		// Check loop termination conditions
		if size != -1 && i >= size {
			break
		}
		// Check if all breakers are connected and in one circuit
		allConnected := true
		for _, isConnected := range connected {
			if !isConnected {
				allConnected = false
				break
			}
		}
		breakerAIndex := pairs[i].i
		breakerBIndex := pairs[i].j
		breakerA := breakers[breakerAIndex]
		breakerB := breakers[breakerBIndex]
		if size == -1 && len(circuits) == 1 && allConnected {
			fmt.Println("size:", len(circuits), "breakers:", len(breakers), "all connected")
			fmt.Println("The last breakers used were:", breakerA, breakerB)
			fmt.Println("Part 2 result:", breakerA.x*breakerB.x)
			break
		}

		if connected[breakerAIndex] && connected[breakerBIndex] {
			// Merge the two circuits
			circuitAID, _ := findBreakerInCircuitMap(breakerA, circuits)
			circuitBID, _ := findBreakerInCircuitMap(breakerB, circuits)
			if circuitAID != circuitBID {
				// Merge circuitB into circuitA
				circuitA := circuits[circuitAID]
				circuitA.breakers = append(circuitA.breakers, circuits[circuitBID].breakers...)
				circuits[circuitAID] = circuitA
				// Remove circuitB
				delete(circuits, circuitBID)
				//fmt.Println("All breakers connected into a single circuit.")
				//fmt.Println("The last breakers where:", breakerA, breakerB)

			}
		} else if connected[breakerAIndex] {
			circuitID, _ := findBreakerInCircuitMap(breakerA, circuits)
			circuit := circuits[circuitID]
			circuit.breakers = append(circuit.breakers, breakerB)
			circuits[circuitID] = circuit
			connected[breakerBIndex] = true
		} else if connected[breakerBIndex] {
			circuitID, _ := findBreakerInCircuitMap(breakerB, circuits)
			circuit := circuits[circuitID]
			circuit.breakers = append(circuit.breakers, breakerA)
			circuits[circuitID] = circuit
			connected[breakerAIndex] = true
		} else {
			// Create a new circuit
			circuits[nextCircuitID] = Circuit{[]Breaker{breakerA, breakerB}}
			connected[breakerAIndex] = true
			connected[breakerBIndex] = true
			nextCircuitID++
		}
		i++
	}
	return circuits
}

func printCircuits(circuits map[int]Circuit) {
	// Create a slice of circuit IDs and sort by breaker count
	type circuitInfo struct {
		id    int
		count int
	}
	infos := []circuitInfo{}
	for id, circuit := range circuits {
		infos = append(infos, circuitInfo{id, len(circuit.breakers)})
	}
	sort.Slice(infos, func(i, j int) bool {
		return infos[i].count > infos[j].count
	})

	// Print circuits in sorted order
	for _, info := range infos {
		circuit := circuits[info.id]
		fmt.Printf("Circuit %d (%d breakers): ", info.id, info.count)
		for _, breaker := range circuit.breakers {
			fmt.Printf("(%d,%d,%d) ", breaker.x, breaker.y, breaker.z)
		}
		fmt.Println()
	}
}

func finalCalculation(circuits map[int]Circuit) int {
	sizes := []int{}
	for _, circuit := range circuits {
		sizes = append(sizes, len(circuit.breakers))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	if len(sizes) < 3 {
		return 0
	}

	return sizes[0] * sizes[1] * sizes[2]
}

func main() {
	//circuits := []Circuit{}
	contents := utils.OpenAndReadFile("input.txt")
	breakers := parseBreakers(contents)
	//fmt.Println("Breakers are:", breakers)
	pairs := findShortestPairsDistance(breakers)
	circuits := connectCircuits(pairs, breakers, -1)
	printCircuits(circuits)
	result := finalCalculation(circuits)
	fmt.Println("Final result is:", result)

}
