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
	inputPath := filepath.Join("cmd", "2022_01", "input.txt")
	part := flag.String("part", "both", "Which part to run: 1, 2, or both")
	flag.Parse()

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
	// TODO: implement part 1
	most := 0
	cur := 0
	for _, line := range lines {
		if line == "" {
			most = max(cur, most)
			cur = 0
			continue
		}
		item, errParse := strconv.Atoi(line)
		if errParse != nil {
			panic(errParse)
		}
		cur += item
	}
	return max(cur, most)
}

func part2(lines []string) any {
	// TODO: implement part 2
	most := 0
	cur := 0
	for _, line := range lines {
		if line == "" {
			most = max(cur, most)
			cur = 0
			continue
		}
		item, errParse := strconv.Atoi(line)
		if errParse != nil {
			panic(errParse)
		}
		cur += item
	}
	return 0
}
