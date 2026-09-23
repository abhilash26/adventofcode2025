package day5

import (
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	min int
	max int
}

func Part1(lines []string) int {
	ranges, idsStart := parseRanges(lines)
	ranges = mergeRanges(ranges)

	freshIngredients := 0

	for _, line := range lines[idsStart:] {
		id, err := strconv.Atoi(line)
		if err != nil {
			continue
		}

		for _, r := range ranges {
			if id >= r.min && id <= r.max {
				freshIngredients++
				break
			}
		}
	}

	return freshIngredients
}

func Part2(lines []string) int {
	ranges, _ := parseRanges(lines)
	ranges = mergeRanges(ranges)

	freshIngredientIDs := 0

	for _, r := range ranges {
		freshIngredientIDs += r.max - r.min + 1
	}

	return freshIngredientIDs
}

func parseRanges(lines []string) ([]Range, int) {
	var ranges []Range

	for i, line := range lines {
		if line == "" {
			return ranges, i + 1
		}

		left, right, found := strings.Cut(line, "-")
		if !found {
			continue
		}

		min, err1 := strconv.Atoi(left)
		max, err2 := strconv.Atoi(right)
		if err1 != nil || err2 != nil {
			continue
		}

		ranges = append(ranges, Range{
			min: min,
			max: max,
		})
	}

	return ranges, len(lines)
}

func mergeRanges(ranges []Range) []Range {
	if len(ranges) < 2 {
		return ranges
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})

	merged := make([]Range, 0, len(ranges))
	merged = append(merged, ranges[0])

	for _, current := range ranges[1:] {
		last := &merged[len(merged)-1]

		if current.min <= last.max {
			if current.max > last.max {
				last.max = current.max
			}
			continue
		}

		merged = append(merged, current)
	}

	return merged
}
