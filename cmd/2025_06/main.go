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
	inputPath := filepath.Join("cmd", "2025_06", *inputFile)

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
	symbols := strings.Fields(lines[len(lines)-1])
	lines = lines[:len(lines)-1]
	problems := [][]int{}
	for _, line := range lines {
		s := strings.Fields(line)
		curRow := []int{}
		for _, numS := range s {
			num, errNum := strconv.Atoi(numS)
			if errNum != nil {
				panic(errNum)
			}
			curRow = append(curRow, num)
		}
		problems = append(problems, curRow)
	}

	sum := 0

	for colI := range len(problems[0]) {
		curSum := 0
		curSym := string(symbols[colI])
		for rowI := range len(problems) {
			val := problems[rowI][colI]
			if curSym == "+" {
				// fmt.Printf("%d + %d\n", curSum, val)
				curSum += val
			} else {
				// fmt.Printf("%d * %d\n", curSum, val)
				if curSum == 0 {
					curSum = val
				} else {
					curSum *= val
				}
			}
		}
		sum += curSum
	}

	return sum
}

func part2(lines []string) any {
	symbols := strings.Fields(lines[len(lines)-1])
	lines = lines[:len(lines)-1]

	curRow := []int{}
	sum := 0
	symI := 0
	for colI := range len(lines[0]) {
		s := ""
		for rowI := range len(lines) {
			row := lines[rowI]
			if char := string(row[colI]); char != " " {
				s += char
			}
		}
		if s != "" {
			num, errNum := strconv.Atoi(s)
			if errNum != nil {
				panic(errNum)
			}
			curRow = append(curRow, num)
		}
		if s == "" || colI == len(lines[0])-1 {
			curSym := symbols[symI]
			curSum := 0
			for _, val := range curRow {
				if curSym == "+" {
					// fmt.Printf("%d + %d\n", curSum, val)
					curSum += val
				} else {
					// fmt.Printf("%d * %d\n", curSum, val)
					if curSum == 0 {
						curSum = val
					} else {
						curSum *= val
					}
				}
			}
			symI++
			sum += curSum
			curRow = []int{}
		}
	}

	return sum
}

func part2Separated(lines []string) any {
	symbols := strings.Fields(lines[len(lines)-1])
	lines = lines[:len(lines)-1]
	problems := [][]int{}

	curRow := []int{}
	for colI := range len(lines[0]) {
		s := ""
		for rowI := range len(lines) {
			row := lines[rowI]
			if char := string(row[colI]); char != " " {
				s += char
			}
		}
		if s == "" {
			problems = append(problems, curRow)
			curRow = []int{}
		} else {
			num, errNum := strconv.Atoi(s)
			if errNum != nil {
				panic(errNum)
			}
			curRow = append(curRow, num)
		}
	}
	problems = append(problems, curRow)

	sum := 0
	for i, problem := range problems {
		curSym := symbols[i]
		curSum := 0
		for _, val := range problem {
			if curSym == "+" {
				// fmt.Printf("%d + %d\n", curSum, val)
				curSum += val
			} else {
				// fmt.Printf("%d * %d\n", curSum, val)
				if curSum == 0 {
					curSum = val
				} else {
					curSum *= val
				}
			}
		}
		sum += curSum
	}

	return sum
}
