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
	inputPath := filepath.Join("cmd", "2025_02", *inputFile)

	lines, err := parse.ParseCSV(inputPath)
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
	for _, line := range lines {
		s := strings.Split(line, "-")
		if len(s) != 2 {
			panic(line)
		}
		start, errStart := strconv.Atoi(s[0])
		if errStart != nil {
			panic(errStart)
		}
		end, errEnd := strconv.Atoi(s[1])
		if errEnd != nil {
			panic(errEnd)
		}
		for i := start; i <= end; i++ {
			cur := strconv.Itoa(i)
			if len(cur)%2 != 0 {
				continue
			}
			firstHalf := cur[0 : len(cur)/2]
			secondHalf := cur[len(cur)/2:]
			if firstHalf == secondHalf {
				sum += i
			}
		}
	}
	return sum
}

func part2(lines []string) any {
	sum := 0
	for _, line := range lines {
		s := strings.Split(line, "-")
		if len(s) != 2 {
			panic(line)
		}
		start, errStart := strconv.Atoi(s[0])
		if errStart != nil {
			panic(errStart)
		}
		end, errEnd := strconv.Atoi(s[1])
		if errEnd != nil {
			panic(errEnd)
		}
		for i := start; i <= end; i++ {
			cur := strconv.Itoa(i)
			if hasPattern(cur) {
				sum += i
			}
		}
	}
	return sum
}

func hasPattern(s string) bool {
	for i := 1; i < (len(s)/2)+1; i++ {
		if len(s)%i != 0 {
			continue
		}
		position := 0
		pattern := false
		for position+i+i <= len(s) {
			pattern = true
			cur := s[position : position+i]
			next := s[position+i : position+i+i]
			position += i
			if cur != next {
				pattern = false
				break
			}
		}
		if pattern == true {
			return true
		}
	}
	return false
}
