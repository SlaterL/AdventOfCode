package main

import (
	"advent-of-code/internal/parse"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
)

var (
	asciiInt = 48
)

func main() {
	part := flag.String("part", "both", "Which part to run: 1, 2, or both")
	inputFile := flag.String("input", "input.txt", "What test file should be used")
	flag.Parse()
	inputPath := filepath.Join("cmd", "2025_03", *inputFile)

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
	for _, line := range lines {
		maximum := 0
		indexOfMax := 0
		for i, char := range line {
			if i == len(line)-1 {
				break
			}
			val := int(char) - asciiInt
			if val > maximum {
				maximum = val
				indexOfMax = i
			}
		}
		maximum2 := 0
		for i := indexOfMax + 1; i < len(line); i++ {
			val := int(line[i]) - asciiInt
			if val > maximum2 {
				maximum2 = val
				indexOfMax = i
			}
		}
		jolt, errJolt := strconv.Atoi(strconv.Itoa(maximum) + strconv.Itoa(maximum2))
		if errJolt != nil {
			panic(errJolt)
		}
		sum += jolt
	}
	return sum
}

func part2(lines []string) any {
	joltParts := 12
	sum := 0
	for _, line := range lines {
		lastIndex := -1
		output := ""
		for i := 1; i <= joltParts; i++ {
			maximum := 0
			maxIndex := 0
			for ii := lastIndex + 1; ii < len(line)-(joltParts-i); ii++ {
				val := int(line[ii]) - asciiInt
				if val > maximum {
					maximum = val
					maxIndex = ii
				}
			}
			output += strconv.Itoa(maximum)
			lastIndex = maxIndex
		}
		jolt, errJolt := strconv.Atoi(output)
		if errJolt != nil {
			panic(errJolt)
		}
		sum += jolt
	}

	return sum
}
