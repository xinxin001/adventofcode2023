package day1

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	lines := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	res := CalculateSumPart1(lines)
	want := 142
	if res != want {
		t.Fatalf("Wrong answer, wanted %v, got %v", want, res)
	}
}

func TestPart2(t *testing.T) {
	dat := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	lines := strings.Split(dat, "\n")
	res := CalculateSumPart2(lines)
	want := 281
	if res != want {
		t.Fatalf("Wrong answer, wanted %v, got %v", want, res)
	}
}
