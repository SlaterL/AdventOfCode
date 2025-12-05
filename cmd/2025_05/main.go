package main

import (
	"advent-of-code/internal/parse"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	part := flag.String("part", "both", "Which part to run: 1, 2, or both")
	inputFile := flag.String("input", "input.txt", "What test file should be used")
	flag.Parse()
	inputPath := filepath.Join("cmd", "2025_05", *inputFile)

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
	ranges := [][]int{}
	ingredients := []int{}
	emptyLineSeen := false
	for _, line := range lines {
		if line == "" {
			emptyLineSeen = true
			continue
		}
		if !emptyLineSeen {
			r := strings.Split(line, "-")
			start, errEnd := strconv.Atoi(r[0])
			if errEnd != nil {
				panic(errEnd)
			}
			end, errEnd := strconv.Atoi(r[1])
			if errEnd != nil {
				panic(errEnd)
			}
			ranges = append(ranges, []int{start, end})
		} else {
			i, errI := strconv.Atoi(line)
			if errI != nil {
				panic(errI)
			}
			ingredients = append(ingredients, i)
		}
	}

	sum := 0
	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if ingredient >= r[0] && ingredient <= r[1] {
				sum++
				break
			}
		}
	}

	return sum
}

func part2(lines []string) any {
	ranges := [][]int{}
	for _, line := range lines {
		if line == "" {
			break
		}
		r := strings.Split(line, "-")
		start, errEnd := strconv.Atoi(r[0])
		if errEnd != nil {
			panic(errEnd)
		}
		end, errEnd := strconv.Atoi(r[1])
		if errEnd != nil {
			panic(errEnd)
		}
		ranges = append(ranges, []int{start, end})
	}
	sum := 0

	for _, r := range consolidateRanges(ranges) {
		sum += r[1] - r[0] + 1
	}

	return sum
}

func consolidateRanges(ranges [][]int) [][]int {
	out := [][]int{}
	for {
		consolidateCount := 0
		out = [][]int{}

		for _, r := range ranges {
			consolidated := false
			for _, rOut := range out {
				startIn := r[0] >= rOut[0] && r[0] <= rOut[1]
				endIn := r[1] >= rOut[0] && r[1] <= rOut[1]
				if startIn && !endIn {
					rOut[1] = r[1]
					consolidated = true
				} else if endIn && !startIn {
					rOut[0] = r[0]
					consolidated = true
				} else if r[0] <= rOut[0] && r[1] >= rOut[1] {
					rOut[0] = r[0]
					rOut[1] = r[1]
					consolidated = true
				} else if startIn && endIn {
					consolidated = true
					continue
				}
			}
			if !consolidated {
				out = append(out, []int{r[0], r[1]})
			} else {
				consolidateCount++
			}
		}
		if consolidateCount == 0 {
			break
		} else {
			ranges = out
		}
	}

	return out
}
