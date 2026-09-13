package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func sumInvalidIDsPart1(left string, right string) int {
	invalidIDsSum := 0

	leftInt, err := strconv.Atoi(strings.TrimSpace(left))
	if err != nil {
		fmt.Println("Left Range can't be converted to Integer, incorrect input")
		return 0
	}
	rightInt, err := strconv.Atoi(strings.TrimSpace(right))
	if err != nil {
		fmt.Println("Right Range can't be converted to Integer, incorrect input")
		return 0
	}

	fmt.Println(left, "-", right)

	if len(left) == len(right) && len(left)%2 != 0 {
		return 0
	}

	for i := leftInt; i <= rightInt; i++ {
		n := strconv.Itoa(i)
		if len(n)%2 != 0 {
			continue
		}

		mid := len(n) / 2
		if n[:mid] == n[mid:] {
			fmt.Println("Number:", i, "First Half:", n[:mid], "secondHalf:", n[mid:])
			invalidIDsSum += i
		}
	}

	return invalidIDsSum
}

func Part1(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	sum := 0

	for productIDRange := range strings.SplitSeq(lines[0], ",") {
		left, right, found := strings.Cut(productIDRange, "-")
		if !found {
			fmt.Println("Product ID range not found correctly, ending program")
			return 0
		}

		sum += sumInvalidIDsPart1(left, right)
	}

	return sum
}

func Part2(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	sum := 0

	for productIDRange := range strings.SplitSeq(lines[0], ",") {
		left, right, found := strings.Cut(productIDRange, "-")
		if !found {
			fmt.Println("Product ID range not found correctly, ending program")
			return 0
		}

		sum += sumInvalidIDsPart2(left, right)
	}

	return sum
}

func sumInvalidIDsPart2(left string, right string) int {
	invalidIDsSum := 0

	leftInt, err := strconv.Atoi(strings.TrimSpace(left))
	if err != nil {
		fmt.Println("Left Range can't be converted to Integer, incorrect input")
		return 0
	}
	rightInt, err := strconv.Atoi(strings.TrimSpace(right))
	if err != nil {
		fmt.Println("Right Range can't be converted to Integer, incorrect input")
		return 0
	}

	fmt.Println(left, "-", right)

	for i := leftInt; i <= rightInt; i++ {
		n := strconv.Itoa(i)
		if isInValidIDPart2(n) {
			invalidIDsSum += i
		}
	}

	return invalidIDsSum
}

func isInValidIDPart2(n string) bool {
	length := len(n)
	for size := 1; size <= length/2; size++ {
		if length%size != 0 {
			continue
		}
		isInvalid := true
		pattern := n[:size]
		for i := size; i < length; i += size {
			if n[i:i+size] != pattern {
				isInvalid = false
				break
			}
		}
		if isInvalid {
			fmt.Println("Number: ", n)
			return true
		}
	}
	return false
}
