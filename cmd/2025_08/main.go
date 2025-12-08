package main

import (
	"advent-of-code/internal/parse"
	"flag"
	"fmt"
	"log"
	"math"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func main() {
	part := flag.String("part", "both", "Which part to run: 1, 2, or both")
	inputFile := flag.String("input", "input.txt", "What test file should be used")
	flag.Parse()
	inputPath := filepath.Join("cmd", "2025_08", *inputFile)

	lines, err := parse.ParseLines(inputPath)
	if err != nil {
		log.Fatalf("failed to parse input file: %v", err)
	}

	switch *part {
	case "1":
		fmt.Println("Part 1:", part1(lines))
	case "2":
		fmt.Println("Part 2:", part2(lines))
	case "both":
		fmt.Println("Part 1:", part1(lines))
		fmt.Println("Part 2:", part2(lines))
	default:
		log.Fatalf("unknown part: %s (expected 1, 2, or both)", *part)
	}
}

type point struct {
	x float64
	y float64
	z float64
}

type pair struct {
	dist float64
	a    *point
	b    *point
}

func calcDist(p1, p2 *point) float64 {
	xDist := math.Abs(p2.x - p1.x)
	yDist := math.Abs(p2.y - p1.y)
	zDist := math.Abs(p2.z - p1.z)

	hyp := math.Sqrt((xDist * xDist) + (yDist * yDist) + (zDist * zDist))
	return hyp
}

func part1(lines []string) any {
	pairs := []pair{}
	points := []*point{}
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, errX := strconv.Atoi(coords[0])
		if errX != nil {
			panic(errX)
		}
		y, errY := strconv.Atoi(coords[1])
		if errY != nil {
			panic(errY)
		}
		z, errZ := strconv.Atoi(coords[2])
		if errZ != nil {
			panic(errZ)
		}
		points = append(points, &point{x: float64(x), y: float64(y), z: float64(z)})
	}

	seen := []*point{}
	for _, pointA := range points {
		for _, pointB := range points {
			if pointA == pointB || slices.Contains(seen, pointB) {
				continue
			}
			pairs = append(pairs, pair{dist: calcDist(pointA, pointB), a: pointA, b: pointB})
		}
		seen = append(seen, pointA)
	}

	slices.SortFunc(pairs, func(a, b pair) int {
		if a.dist < b.dist {
			return -1
		} else if a.dist > b.dist {
			return 1
		}
		return 0
	})

	circuits := [][]*point{}
	paired := 0
	for _, pair := range pairs {
		if paired == 1000 {
			break
		}
		aCircuit := -1
		bCircuit := -1
		for i, circuit := range circuits {
			if slices.Contains(circuit, pair.a) {
				aCircuit = i
			}
			if slices.Contains(circuit, pair.b) {
				bCircuit = i
			}
		}

		if aCircuit == -1 && bCircuit == -1 { // case a and b unconnected
			circuits = append(circuits, []*point{pair.a, pair.b})
		} else if aCircuit >= 0 && bCircuit == -1 { // case a connected, b unconnected
			circuits[aCircuit] = append(circuits[aCircuit], pair.b)
		} else if bCircuit >= 0 && aCircuit == -1 { // case b connected, a unconnected
			circuits[bCircuit] = append(circuits[bCircuit], pair.a)
		} else if aCircuit != bCircuit { // case a and b connected in different circuits
			circuits[aCircuit] = append(circuits[aCircuit], circuits[bCircuit]...)
			circuits = append(circuits[:bCircuit], circuits[bCircuit+1:]...)
		}
		paired++
	}

	slices.SortFunc(circuits, func(a, b []*point) int {
		if len(a) > len(b) {
			return -1
		} else if len(b) > len(a) {
			return 1
		}
		return 0
	})
	return len(circuits[0]) * len(circuits[1]) * len(circuits[2])
}

func part2(lines []string) any {
	pairs := []pair{}
	points := []*point{}
	for _, line := range lines {
		coords := strings.Split(line, ",")
		x, errX := strconv.Atoi(coords[0])
		if errX != nil {
			panic(errX)
		}
		y, errY := strconv.Atoi(coords[1])
		if errY != nil {
			panic(errY)
		}
		z, errZ := strconv.Atoi(coords[2])
		if errZ != nil {
			panic(errZ)
		}
		points = append(points, &point{x: float64(x), y: float64(y), z: float64(z)})
	}

	seen := []*point{}
	for _, pointA := range points {
		for _, pointB := range points {
			if pointA == pointB || slices.Contains(seen, pointB) {
				continue
			}
			pairs = append(pairs, pair{dist: calcDist(pointA, pointB), a: pointA, b: pointB})
		}
		seen = append(seen, pointA)
	}

	slices.SortFunc(pairs, func(a, b pair) int {
		if a.dist < b.dist {
			return -1
		} else if a.dist > b.dist {
			return 1
		}
		return 0
	})

	circuits := [][]*point{}
	for _, pair := range pairs {
		aCircuit := -1
		bCircuit := -1
		for i, circuit := range circuits {
			if slices.Contains(circuit, pair.a) {
				aCircuit = i
			}
			if slices.Contains(circuit, pair.b) {
				bCircuit = i
			}
		}

		if aCircuit == -1 && bCircuit == -1 { // case a and b unconnected
			circuits = append(circuits, []*point{pair.a, pair.b})
		} else if aCircuit >= 0 && bCircuit == -1 { // case a connected, b unconnected
			circuits[aCircuit] = append(circuits[aCircuit], pair.b)
		} else if bCircuit >= 0 && aCircuit == -1 { // case b connected, a unconnected
			circuits[bCircuit] = append(circuits[bCircuit], pair.a)
		} else if aCircuit != bCircuit { // case a and b connected in different circuits
			circuits[aCircuit] = append(circuits[aCircuit], circuits[bCircuit]...)
			circuits = append(circuits[:bCircuit], circuits[bCircuit+1:]...)
		}
		if len(circuits) == 1 && len(circuits[0]) == len(points) {
			return pair.a.x * pair.b.x
		}
	}

	return 0
}
