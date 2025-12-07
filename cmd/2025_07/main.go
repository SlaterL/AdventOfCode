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
	inputPath := filepath.Join("cmd", "2025_07", *inputFile)

	linesPt1, err := parse.ParseLines(inputPath)
	if err != nil {
		log.Fatalf("failed to parse input file: %v", err)
	}
	linesPt2, err := parse.ParseLines(inputPath)
	if err != nil {
		log.Fatalf("failed to parse input file: %v", err)
	}

	switch *part {
	case "1":
		fmt.Println("Part 1:", part1(linesPt1))
	case "2":
		fmt.Println("Part 2:", part2(linesPt2))
	case "both":
		fmt.Println("Part 1:", part1(linesPt1))
		fmt.Println("Part 2:", part2(linesPt2))
	default:
		log.Fatalf("unknown part: %s (expected 1, 2, or both)", *part)
	}
}

func part1(lines []string) any {
	splits := 0
	lines[0] = strings.ReplaceAll(lines[0], "S", "|")
	for lineIndex, line := range lines {
		if lineIndex == len(lines)-1 {
			break
		}
		nextLine := []rune(lines[lineIndex+1])
		for charIndex, char := range line {
			if char == '|' && lines[lineIndex+1][charIndex] == '^' {
				splits++
				nextLine[charIndex+1] = '|'
				nextLine[charIndex-1] = '|'
			} else if char == '|' {
				nextLine[charIndex] = '|'
			}
		}
		lines[lineIndex+1] = string(nextLine)
	}
	return splits
}

func part2(lines []string) any {
	startingIndex := strings.Index(lines[0], "S")
	splitterCache := map[string]int{}
	return followParticle(lines, 0, startingIndex, splitterCache)
}

func followParticle(lines []string, curLine, particleIndex int, cache map[string]int) int {
	if curLine == len(lines)-1 {
		return 1
	}
	nextLine := []rune(lines[curLine+1])
	if nextLine[particleIndex] == '^' {
		key := strconv.Itoa(curLine+1) + "|" + strconv.Itoa(particleIndex)
		if val := cache[key]; val != 0 {
			return val
		}
		val := followParticle(lines, curLine+1, particleIndex+1, cache) + followParticle(lines, curLine+1, particleIndex-1, cache)
		cache[key] = val
		return val
	}

	return followParticle(lines, curLine+1, particleIndex, cache)
}
