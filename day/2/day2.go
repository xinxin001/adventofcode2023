package day2

import (
	"strconv"
	"strings"
)

func ValidateGames(lines []string) int {
	var idSum int
	for idx, line := range lines {
		game := strings.Split(line, ": ")
		rounds := strings.Split(game[1], "; ")
		if roundsValid(rounds) {
			idSum += idx + 1
		}
	}
	return idSum
}

func roundsValid(rounds []string) bool {
	for _, round := range rounds {
		cubes := strings.Split(round, ", ")
		colorCount := make(map[string]int)
		for _, cube := range cubes {
			cubecount := strings.Split(cube, " ")
			count, _ := strconv.Atoi(cubecount[0])
			cubeColor := cubecount[1]
			colorCount[cubeColor] += count
		}
		if colorCount["red"] > 12 || colorCount["green"] > 13 || colorCount["blue"] > 14 {
			return false
		}
	}
	return true
}

func CalculatePower(lines []string) int {
	var powerSum int
	for _, line := range lines {
		game := strings.Split(line, ": ")
		rounds := strings.Split(game[1], "; ")
		powerSum += calculateGamePower(rounds)
	}
	return powerSum
}

func calculateGamePower(rounds []string) int {
	gameCount := make(map[string]int)
	for _, round := range rounds {
		cubes := strings.Split(round, ", ")
		roundCount := make(map[string]int)
		for _, cube := range cubes {
			cubecount := strings.Split(cube, " ")
			count, _ := strconv.Atoi(cubecount[0])
			cubeColor := cubecount[1]
			roundCount[cubeColor] += count
		}
		gameCount["red"] = max(gameCount["red"], roundCount["red"])
		gameCount["green"] = max(gameCount["green"], roundCount["green"])
		gameCount["blue"] = max(gameCount["blue"], roundCount["blue"])
	}
	return gameCount["red"] * gameCount["green"] * gameCount["blue"]
}
