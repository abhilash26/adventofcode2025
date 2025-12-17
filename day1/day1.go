package day1

import (
	"strconv"
)


func getDialRotation(current int, line string) int {
	rotationType := line[0]
	rotationValue, err := strconv.Atoi(line[1:])
	if err != nil {
		return 0
	}
	switch rotationType {
	case 'L':
		if current - rotationValue < 0 {	
			return 100 - (rotationValue - current)
		} else {
			return current - rotationValue
		}
	case 'R':
		if current + rotationValue > 99 {
			return current + rotationValue - 100
		} else {
			return current + rotationValue
		}
	}
	return 0
}


func Part1(lines []string) int {
	countZero := 0
	current := 50
	for _, line := range lines {
		current = getDialRotation(current, line)
		if current == 0 {
			countZero++
		}
	}
	return countZero
}

func Part2(lines []string) int {
	return 0
}
