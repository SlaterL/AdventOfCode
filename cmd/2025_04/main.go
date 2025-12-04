package main

import (
	"advent-of-code/internal/parse"
	"flag"
	"fmt"
	"log"
	"path/filepath"
)

var (
	adj = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
)

func main() {
	part := flag.String("part", "both", "Which part to run: 1, 2, or both")
	inputFile := flag.String("input", "input.txt", "What test file should be used")
	flag.Parse()
	inputPath := filepath.Join("cmd", "2025_04", *inputFile)

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

func part1(lines []string) any {
	sum := 0
	rowMax := len(lines[0]) - 1
	colMax := len(lines) - 1
	for row, line := range lines {
		for col, char := range line {
			if char != '@' {
				continue
			}
			surrounding := 0

			for _, deltas := range adj {
				dRow := row + deltas[0]
				dCol := col + deltas[1]
				if dRow < 0 || dRow > rowMax || dCol < 0 || dCol > colMax {
					continue
				}
				if lines[dRow][dCol] == '@' {
					surrounding++
				}
			}
			if surrounding <= 3 {
				sum++
			}
		}
	}
	return sum
}

func part2(lines []string) any {
	sum := 0
	rowMax := len(lines[0]) - 1
	colMax := len(lines) - 1
	for {
		removed := 0
		for row, line := range lines {
			for col, char := range line {
				if char != '@' {
					continue
				}
				surrounding := 0

				for _, deltas := range adj {
					dRow := row + deltas[0]
					dCol := col + deltas[1]
					if dRow < 0 || dRow > rowMax || dCol < 0 || dCol > colMax {
						continue
					}
					if lines[dRow][dCol] == '@' {
						surrounding++
					}
				}
				if surrounding <= 3 {
					sum++
					removed++
					lines[row] = lines[row][:col] + "." + lines[row][col+1:]
				}
			}
		}
		if removed == 0 {
			break
		}
	}
	return sum
}
