package main

import (
	"fmt"
	"os"
	"strings"

	day1 "github.com/xinxin001/adventofcode2023/day/1"
)

func main() {
	dat, err := os.ReadFile("./day/1/input.txt")
	if err != nil {
		fmt.Println("Error reading file")
	}
	lines := strings.Split(string(dat), "\n")
	ans1 := day1.CalculateSumPart1(lines)
	ans2 := day1.CalculateSumPart2(lines)
	fmt.Printf("Answer to day 1 part 1 is: %v\n", ans1)
	fmt.Printf("Answer to day 1 part 2 is: %v\n", ans2)
}
