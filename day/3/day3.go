package day3

import (
	"strconv"
	"unicode"
)

/*
1. Read digits from left->right, while reading, scan for adjacent to validate its inclusion in the sum
2. Form the number from the digit and add it to the sum
EDGE CASE: Numbers ending at the end of line, not processing buffer if that happens and number is lost
SOLUTION: Add a period at the end of each line, it extends the line, no number finishes at EOL and they are always processed
*/
func CalculateEngineSum(lines []string) int {
	var totalSum int
	ROW, COL := len(lines), len(lines[0])
	for i, line := range lines {
		var buffer int
		var isAdjacent bool
		line = line + "."
		for j, runeSymbol := range line {
			if digit, err := strconv.Atoi(string(runeSymbol)); err == nil {
				buffer = buffer*10 + digit
				if !isAdjacent {
					for _, tuple := range [8][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}} {
						r, c := i+tuple[0], j+tuple[1]
						if r >= 0 && r < ROW && c >= 0 && c < COL {
							if !unicode.IsDigit(rune(lines[r][c])) && string(lines[r][c]) != "." {
								isAdjacent = true
								break
							}
						}
					}
				}
			} else {
				if isAdjacent {
					totalSum += buffer
				}
				buffer = 0
				isAdjacent = false
			}
		}
	}
	return totalSum
}
