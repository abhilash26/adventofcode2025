package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"adventofcode2025/day1"
	"adventofcode2025/day2"
	"adventofcode2025/day3"
)

type PartKey struct {
	Day   int
	Part  int
	Input string
}

var parts = map[PartKey]func([]string) int{
	{Day: 1, Part: 1, Input: "test.txt"}: day1.Part1,
	{Day: 1, Part: 1, Input: "real.txt"}: day1.Part1,
	{Day: 1, Part: 2, Input: "test.txt"}: day1.Part2,
	{Day: 1, Part: 2, Input: "real.txt"}: day1.Part2,

	{Day: 2, Part: 1, Input: "test.txt"}: day2.Part1,
	{Day: 2, Part: 1, Input: "real.txt"}: day2.Part1,
	{Day: 2, Part: 2, Input: "test.txt"}: day2.Part2,
	{Day: 2, Part: 2, Input: "real.txt"}: day2.Part2,

	{Day: 3, Part: 1, Input: "test.txt"}: day3.Part1,
	{Day: 3, Part: 1, Input: "real.txt"}: day3.Part1,
	{Day: 3, Part: 2, Input: "test.txt"}: day3.Part2,
	{Day: 3, Part: 2, Input: "real.txt"}: day3.Part2,
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func main() {
	day := flag.Int("day", 1, "Day number")
	part := flag.Int("part", 1, "Part number")
	input := flag.String("input", "real.txt", "Input file")
	flag.Parse()

	key := PartKey{
		Day:   *day,
		Part:  *part,
		Input: *input,
	}

	partFunc, ok := parts[key]
	if !ok {
		fmt.Println("Invalid day, part, or input")
		return
	}

	filePath := filepath.Join(
		fmt.Sprintf("day%d", *day),
		*input,
	)

	lines := readFile(filePath)
	result := partFunc(lines)

	fmt.Printf("Day %d, Part %d, Input %s: %d\n", *day, *part, *input, result)
}
