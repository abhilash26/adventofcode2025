package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"adventofcode2025/day1"
)

func ReadFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	day := flag.Int("day", 1, "Day number 1-12")
	part := flag.Int("part", 1, "Part number 1-2")
	input := flag.String("input", "input.txt", "Input file")
	flag.Parse()

	filePath := filepath.Join(fmt.Sprintf("day%d", *day), *input)
	fmt.Println("Input file: ", filePath)
	lines := ReadFile(filePath)

	switch *day {
	case 1:
		switch *part {
		case 1:
			fmt.Println("Day1, Part1")
			out := day1.Part1(lines)
			fmt.Println("Part 1: ", out)
		case 2:
			fmt.Println("Day1, Part2")
			out := day1.Part2(lines)
			fmt.Println("Part 2: ", out)
		default:
			fmt.Println("Invalid part")
		}
	default:
		fmt.Println("Invalid day")
	}
}
