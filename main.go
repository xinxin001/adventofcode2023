package main

import (
	"fmt"
	"os"
	"strings"

	day1 "github.com/xinxin001/adventofcode2023/day/1"
	day2 "github.com/xinxin001/adventofcode2023/day/2"
	day3 "github.com/xinxin001/adventofcode2023/day/3"
)

func main() {
	runDay3()
}

func loadFile(path string) (string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		return "", err
	}
	return string(dat), nil
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func runDay1() {
	file, _ := loadFile("./day/1/input.txt")
	lines := strings.Split(file, "\n")
	ans1 := day1.CalculateSumPart1(lines)
	ans2 := day1.CalculateSumPart2(lines)
	fmt.Printf("Answer to day 1 part 1 is: %v\n", ans1)
	fmt.Printf("Answer to day 1 part 2 is: %v\n", ans2)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func runDay2() {
	file, _ := loadFile("./day/2/input.txt")
	lines := strings.Split(string(file), "\n")
	ans1 := day2.ValidateGames(lines)
	ans2 := day2.CalculatePower(lines)
	fmt.Printf("Answer to day 2 part 1 is: %v\n", ans1)
	fmt.Printf("Answer to day 2 part 2 is: %v\n", ans2)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func runDay3() {
	file, _ := loadFile("./day/3/input.txt")
	lines := strings.Split(string(file), "\n")
	ans1 := day3.CalculateEngineSumPart1(lines)
	ans2 := day3.CalculateEngineSumPart2(lines)
	fmt.Printf("Answer to day 3 part 1 is: %v\n", ans1)
	fmt.Printf("Answer to day 3 part 2 is: %v\n", ans2)
}
