package day1

import (
	"strconv"
	"strings"
)

func CalculateSumPart1(lines []string) int {
	var total int
	for _, line := range lines {
		var digits []string
		for _, runeVal := range line {
			if _, err := strconv.Atoi(string(runeVal)); err == nil {
				digits = append(digits, string(runeVal))
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(string(line[i])); err == nil {
				digits = append(digits, string(line[i]))
				break
			}
		}
		number := strings.Join(digits, "")
		if val, err := strconv.Atoi(number); err == nil {
			total += val
		}
	}
	return total
}

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func CalculateSumPart2(lines []string) int {
	var total int
	for _, line := range lines {
		var digits []string
	Forward:
		for lineIdx := 0; lineIdx < len(line); lineIdx++ {
			if _, err := strconv.Atoi(string(line[lineIdx])); err == nil {
				digits = append(digits, string(line[lineIdx]))
				break Forward
			}
			for offset := 5; offset >= 3; offset-- {
				if lineIdx+offset <= len(line) {
					if val, ok := numbers[line[lineIdx:lineIdx+offset]]; ok {
						digits = append(digits, val)
						break Forward
					}
				}
			}
		}
	Reverse:
		for lineIdx := len(line) - 1; lineIdx >= 0; lineIdx-- {
			if _, err := strconv.Atoi(string(line[lineIdx])); err == nil {
				digits = append(digits, string(line[lineIdx]))
				break Reverse
			}
			for offset := -4; offset <= -2; offset++ {
				if lineIdx+offset >= 0 {
					if val, ok := numbers[line[lineIdx+offset:lineIdx+1]]; ok {
						digits = append(digits, val)
						break Reverse
					}
				}
			}
		}
		number := strings.Join(digits, "")
		if val, err := strconv.Atoi(number); err == nil {
			total += val
		}
	}
	return total
}
