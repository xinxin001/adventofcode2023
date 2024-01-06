package day4

import (
	"strconv"
	"strings"
)

func processLine(line string) (winningNumbers map[int]struct{}, numbers map[int]struct{}) {
	split := func(c rune) bool { return c == ':' || c == '|' }
	fields := strings.FieldsFunc(line, split)
	winningNumbers = map[int]struct{}{}
	for _, n := range strings.Fields(fields[1]) {
		val, _ := strconv.Atoi(n)
		winningNumbers[int(val)] = struct{}{}
	}
	numbers = map[int]struct{}{}
	for _, n := range strings.Fields(fields[2]) {
		val, _ := strconv.Atoi(n)
		numbers[val] = struct{}{}
	}
	return
}

func CalculatePointsPart1(lines []string) int {
	var totalPoints int
	for _, line := range lines {
		var points int
		winningNumbers, numbers := processLine(line)
		for n := range numbers {
			if _, ok := winningNumbers[n]; ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		totalPoints += points
	}
	return totalPoints
}

func CalculateCardCount(lines []string) int {
	var cardCount int
	cardCopies := map[int]int{}
	for idx := range lines {
		cardCopies[idx+1] = 1
	}
	for idx, line := range lines {
		var points int
		winningNumbers, numbers := processLine(line)
		for n := range numbers {
			if _, ok := winningNumbers[n]; ok {
				points += 1
			}
		}
		var currentCardCopies int = cardCopies[idx+1]
		for i := (idx + 2); i < (idx + 2 + points); i++ {
			// fmt.Printf("At card %v, winning %v copies of card number %v\n", idx+1, currentCardCopies, i)
			cardCopies[i] += currentCardCopies
		}
		// fmt.Println("-------")
	}
	for _, val := range cardCopies {
		cardCount += val
	}
	return cardCount
}
