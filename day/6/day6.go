package day6

import (
	"strconv"
	"strings"
)

func CalculateWinningMargin(lines []string) int {
	var marginProduct int = 1
	var raceLength []int
	var distanceRecord []int
	for _, field := range strings.Fields(lines[0])[1:] {
		val, _ := strconv.Atoi(field)
		raceLength = append(raceLength, val)
	}
	for _, field := range strings.Fields(lines[1])[1:] {
		val, _ := strconv.Atoi(field)
		distanceRecord = append(distanceRecord, val)
	}
	for i := 0; i < len(raceLength); i++ {
		var winningMoves int
		for t := 0; t <= raceLength[i]; t++ {
			// t represents time holding the button down, also represents the speed
			timeLeft := raceLength[i] - t
			distance := timeLeft * t
			if distance > distanceRecord[i] {
				winningMoves += 1
			}
		}
		marginProduct *= winningMoves
	}
	return marginProduct
}

func CalculateWinningMargin2(lines []string) int {
	raceLength, _ := strconv.Atoi(strings.Join(strings.Fields(lines[0])[1:], ""))
	distanceRecord, _ := strconv.Atoi(strings.Join(strings.Fields(lines[1])[1:], ""))
	var winningMoves int
	for t := 0; t <= raceLength; t++ {
		// t represents time holding the button down, also represents the speed
		timeLeft := raceLength - t
		distance := timeLeft * t
		if distance > distanceRecord {
			winningMoves += 1
		}
	}
	return winningMoves
}
