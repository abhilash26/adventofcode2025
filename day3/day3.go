package day3

func Part1(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	sum := 0
	joltage := 0

	for _, line := range lines {
		joltage = getMaxJoltagePart1(line)
		println("Joltage for line ", line, " is ", joltage)
		sum = sum + joltage
	}

	return sum
}

func getBankInt(bank string, bank_int []int, seen []bool, index int) (int, error) {
	if !seen[index] {
		value := int(bank[index] - '0')
		bank_int[index] = value
		seen[index] = true
	}
	return bank_int[index], nil
}

func getMaxJoltagePart1(bank string) int {
	first_battery := 0
	second_battery := 0

	length := len(bank)

	bank_int := make([]int, length)
	seen := make([]bool, length)

	for i := 0; i < length-1; i++ {
		value_i, err := getBankInt(bank, bank_int, seen, i)
		if err != nil {
			println("Could not convert ", bank[i], " to number")
			return 0
		}

		if value_i > first_battery {
			first_battery = value_i
			// Reset second_battery
			second_battery = 0
		} else {
			continue
		}

		for j := i + 1; j < length; j++ {
			value_j, err := getBankInt(bank, bank_int, seen, j)
			if err != nil {
				println("Could not convert ", bank[j], " to number")
				return 0
			}
			if value_j > second_battery {
				second_battery = value_j
			}
		}
	}

	return first_battery*10 + second_battery
}

func Part2(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	sum := 0

	for _, line := range lines {
		joltage := getMaxJoltagePart2(line, 12)
		println("Joltage for line", line, "is", joltage)
		sum += joltage
	}

	return sum
}

func getMaxJoltagePart2(line string, totalDigits int) int {
	length := len(line)

	if length < totalDigits {
		println("Invalid requirement")
		return 0
	}

	joltage := 0
	minIndex := 0

	for digitsLeft := totalDigits; digitsLeft > 0; digitsLeft-- {
		maxIndex := length - digitsLeft

		index, digit := getLargestDigit(minIndex, maxIndex, line)

		joltage = joltage*10 + digit
		minIndex = index + 1
	}

	return joltage
}

func getLargestDigit(minIndex, maxIndex int, line string) (index, value int) {
	index = minIndex
	value = -1

	for i := minIndex; i <= maxIndex; i++ {
		digit := int(line[i] - '0')

		if digit > value {
			value = digit
			index = i
		}
	}

	return index, value
}
