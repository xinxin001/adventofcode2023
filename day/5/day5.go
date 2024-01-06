package day5

import (
	"math"
	"strconv"
	"strings"
)

type SeedMap struct {
	mappings []SeedInterval
}

type SeedInterval struct {
	destinationRangeStart, sourceRangeStart, rangeLength int
}

func FindLowestLocationNumber(lines []string) (int, int) {
	lowestPart1, lowestPart2 := int(math.MaxInt), int(math.MaxInt)
	var seeds []int
	seedString := strings.Split(lines[0], ": ")[1]
	for _, val := range strings.Fields(seedString) {
		intVal, _ := strconv.Atoi(val)
		seeds = append(seeds, intVal)
	}
	var maps []SeedMap
	for _, mapping := range lines[1:] {
		mapString := strings.Split(strings.Split(mapping, ":\n")[1], "\n")
		seedMap := SeedMap{}
		for _, line := range mapString {
			fields := strings.Fields(line)
			destinationRangeStart, _ := strconv.Atoi(fields[0])
			sourceRangeStart, _ := strconv.Atoi(fields[1])
			rangeLength, _ := strconv.Atoi(fields[2])
			seedMap.mappings = append(seedMap.mappings, SeedInterval{destinationRangeStart, sourceRangeStart, rangeLength})
		}
		maps = append(maps, seedMap)
	}
	calc := func(seed int) int {
		for _, m := range maps {
			for _, r := range m.mappings {
				if seed >= r.sourceRangeStart && seed < r.sourceRangeStart+r.rangeLength {
					seed = r.destinationRangeStart + seed - r.sourceRangeStart
					break
				}
			}
		}
		return seed
	}
	for i, s := range seeds {
		lowestPart1 = min(calc(s), lowestPart1)
		if i%2 == 0 {
			for s := seeds[i]; s < seeds[i]+seeds[i+1]; s++ {
				lowestPart2 = min(calc(s), lowestPart2)
			}
		}
	}
	return lowestPart1, lowestPart2
}
