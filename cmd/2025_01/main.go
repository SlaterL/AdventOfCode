package main

import (
	"advent-of-code/internal/parse"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
)

func main() {
	part := flag.String("part", "both", "Which part to run: 1, 2, or both")
	inputFile := flag.String("input", "input.txt", "What test file should be used")
	flag.Parse()
	inputPath := filepath.Join("cmd", "2025_01", *inputFile)

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
	count := 0
	curVal := 50
	for _, line := range lines {
		dist, errDist := strconv.Atoi(line[1:])
		if errDist != nil {
			panic(errDist)
		}
		if line[0] == 'L' {
			curVal -= dist
			for curVal < 0 {
				curVal += 100
			}
		} else {
			curVal += dist
			for curVal >= 100 {
				curVal -= 100
			}
		}
		if curVal == 0 {
			count++
		}
	}

	return count
}

func part2(lines []string) any {
	count := 0
	curVal := 50
	for _, line := range lines {
		dist, errDist := strconv.Atoi(line[1:])
		if errDist != nil {
			panic(errDist)
		}
		if line[0] == 'L' {
			for range dist {
				curVal--
				if curVal < 0 {
					curVal += 100
				}
				if curVal == 0 {
					count++
				}
			}
		} else {
			for range dist {
				curVal++
				if curVal >= 100 {
					curVal -= 100
				}
				if curVal == 0 {
					count++
				}
			}
		}

	}

	return count
}
