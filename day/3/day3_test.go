package day3

import (
	"strings"
	"testing"
)

var engine string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestPart1(t *testing.T) {
	lines := strings.Split(engine, "\n")
	res := CalculateEngineSumPart1(lines)
	want := 4361
	if want != res {
		t.Fatalf("Wanted: %v, got: %v", want, res)
	}
}

func TestPart2(t *testing.T) {
	lines := strings.Split(engine, "\n")
	res := CalculateEngineSumPart2(lines)
	want := 467835
	if want != res {
		t.Fatalf("Wanted: %v, got: %v", want, res)
	}
}
