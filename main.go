package main

import (
	"fmt"
	"os"
	"strings"

	day1 "github.com/xinxin001/adventofcode2023/day/1"
	day2 "github.com/xinxin001/adventofcode2023/day/2"
)

func main() {
	dat, err := os.ReadFile("./day/2/input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
	}
	lines := strings.Split(string(dat), "\n")
	ans1 := day2.ValidateGames(lines)
	ans2 := day2.CalculatePower(lines)
	fmt.Printf("Answer to day 2 part 1 is: %v\n", ans1)
	fmt.Printf("Answer to day 2 part 2 is: %v\n", ans2)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func runDay1() {
	dat, err := os.ReadFile("./day/1/input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
	}
	lines := strings.Split(string(dat), "\n")
	ans1 := day1.CalculateSumPart1(lines)
	ans2 := day1.CalculateSumPart2(lines)
	fmt.Printf("Answer to day 1 part 1 is: %v\n", ans1)
	fmt.Printf("Answer to day 1 part 2 is: %v\n", ans2)
}
