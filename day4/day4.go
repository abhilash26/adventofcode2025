package day4

func Part1(lines []string) int {
	if len(lines) == 0 {
		return 0
	}
	return accessiblePaperRollsPart1(lines)
}

func accessiblePaperRollsPart1(lines []string) int {
	const character = '@'
	total_lines := len(lines)
	line_length := len(lines[0])

	accessibleRolls := 0
	rollSum := 0

	for i := range total_lines {
		i_neg := i-1 >= 0
		i_pos := i+1 < total_lines
		for j := range line_length {
			if lines[i][j] != '@' {
				continue
			}
			j_neg := j-1 >= 0
			j_pos := j+1 < line_length

			if i_neg {
				if j_neg && lines[i-1][j-1] == character {
					rollSum++
				}
				if lines[i-1][j] == character {
					rollSum++
				}
				if j_pos && lines[i-1][j+1] == character {
					rollSum++
				}
			}

			if j_neg && lines[i][j-1] == character {
				rollSum++
			}
			if j_pos && lines[i][j+1] == character {
				rollSum++
			}

			if i_pos {
				if j_neg && lines[i+1][j-1] == character {
					rollSum++
				}
				if lines[i+1][j] == character {
					rollSum++
				}
				if j_pos && lines[i+1][j+1] == character {
					rollSum++
				}
			}

			if rollSum < 4 {
				accessibleRolls++
			}
			rollSum = 0
		}

	}

	return accessibleRolls
}

func Part2(lines []string) int {
	if len(lines) == 0 {
		return 0
	}
	totalRemoved := 0
	removed := 0
	byteLines :=
		stringsToBytes(lines)

	for {
		removed, byteLines = accessiblePaperRollsPart2(byteLines)
		totalRemoved += removed
		if removed == 0 {
			break
		}
	}
	return totalRemoved
}

func stringsToBytes(lines []string) [][]byte {
	output := make([][]byte, len(lines))
	for i := range lines {
		output[i] = []byte(lines[i])
	}
	return output
}

func accessiblePaperRollsPart2(lines [][]byte) (int, [][]byte) {
	const character = '@'
	total_lines := len(lines)
	line_length := len(lines[0])

	grid := make([][]byte, len(lines))

	accessibleRolls := 0
	rollSum := 0

	for i := range total_lines {
		grid[i] = []byte(lines[i])
		i_neg := i-1 >= 0
		i_pos := i+1 < total_lines
		for j := range line_length {
			if lines[i][j] != '@' {
				continue
			}
			j_neg := j-1 >= 0
			j_pos := j+1 < line_length

			if i_neg {
				if j_neg && lines[i-1][j-1] == character {
					rollSum++
				}
				if lines[i-1][j] == character {
					rollSum++
				}
				if j_pos && lines[i-1][j+1] == character {
					rollSum++
				}
			}

			if j_neg && lines[i][j-1] == character {
				rollSum++
			}
			if j_pos && lines[i][j+1] == character {
				rollSum++
			}

			if i_pos {
				if j_neg && lines[i+1][j-1] == character {
					rollSum++
				}
				if lines[i+1][j] == character {
					rollSum++
				}
				if j_pos && lines[i+1][j+1] == character {
					rollSum++
				}
			}

			if rollSum < 4 {
				grid[i][j] = 'x'
				accessibleRolls++
			}
			rollSum = 0
		}

	}

	return accessibleRolls, grid
}
