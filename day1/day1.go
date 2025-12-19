package day1

import (
	"strconv"
)

func Part1(lines []string) int {
	countZero := 0
	current := 50

	for _, line := range lines {
		rotationDirection := line[0]
		rotationVal, _ := strconv.Atoi(line[1:])
		switch rotationDirection {
		case 'L':
			current -= rotationVal
		case 'R':
			current += rotationVal
		}
		current %= 100
		if current == 0 {
			countZero++
		}
	}
	return countZero
}

func Part2(lines []string) int {
	countZero := 0
	current := 50
	distancetoZero := 0

	for _, line := range lines {
		rotationDirection := line[0]
		rotationVal, _ := strconv.Atoi(line[1:])


		if rotationDirection == 'L' {
			distancetoZero = 100 - current
			current += rotationVal
		} else {
			distancetoZero = current
			current -= rotationVal
		}

		// Update countZero
		if distancetoZero == 0 {
			countZero += rotationVal / 100
		} else if rotationVal >= distancetoZero {
			countZero += 1 + (rotationVal-distancetoZero)/100
		}

		// Update current
		current %= 100
		if current < 0 {
			current += 100
		}

	}
	return countZero
}
