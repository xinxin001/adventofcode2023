package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	day1 "github.com/xinxin001/adventofcode2023/day/1"
	day2 "github.com/xinxin001/adventofcode2023/day/2"
	day3 "github.com/xinxin001/adventofcode2023/day/3"
	day4 "github.com/xinxin001/adventofcode2023/day/4"
	day5 "github.com/xinxin001/adventofcode2023/day/5"
	day6 "github.com/xinxin001/adventofcode2023/day/6"
)

func main() {
	runDay6()
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func extractPageLinks() {
	url := "https://en.wikipedia.org/wiki/Wikipedia"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching the URL:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}
	html := string(body)
	fmt.Println(html)
	linkRegex := regexp.MustCompile(`href="(https?://[^"]+)"`)
	links := linkRegex.FindAllStringSubmatch(html, -1)
	for _, match := range links {
		fmt.Println(match[1]) // The second element in the match slice contains the link
	}
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

//lint:ignore U1000 Ignore unused function temporarily for debugging
func runDay4() {
	file, _ := loadFile("./day/4/input.txt")
	lines := strings.Split(file, "\n")
	ans1 := day4.CalculatePointsPart1(lines)
	ans2 := day4.CalculateCardCount(lines)
	fmt.Printf("Answer to day 4 part 1 is: %v\n", ans1)
	fmt.Printf("Answer to day 4 part 2 is: %v\n", ans2)
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func runDay5() {
	file, _ := loadFile("./day/5/input.txt")
	lines := strings.Split(file, "\n\n")
	ans1, ans2 := day5.FindLowestLocationNumber(lines)
	fmt.Printf("Answer to day 5 part 1 is: %v\n", ans1)
	fmt.Printf("Answer to day 5 part 2 is: %v\n", ans2)
}

func runDay6() {
	file, _ := loadFile("./day/6/input.txt")
	lines := strings.Split(file, "\n")
	ans1 := day6.CalculateWinningMargin(lines)
	ans2 := day6.CalculateWinningMargin2(lines)
	fmt.Printf("Answer to day 6 part 1 is: %v\n", ans1)
	fmt.Printf("Answer to day 6 part 2 is: %v\n", ans2)
}
