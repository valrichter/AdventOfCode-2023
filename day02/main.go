package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var possibleConfigs = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	lines := readInput("day02/input.txt")

	partOneResult := partOne(lines)
	partTwoResult := partTwo(lines)

	fmt.Println("Part one result:", partOneResult)
	fmt.Println("Part two result:", partTwoResult)
}

func readInput(file string) []string {
	rawFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(rawFile), "\n")
}

func getValues(line string) map[string]int {
	values := strings.Split(line, ":")
	colorValues := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	colorString := strings.Split(values[1], ";")
	for _, set := range colorString {
		items := strings.Split(set, ",")
		for _, item := range items {
			value := strings.Split(item, " ")

			colorValue := value[1]
			colorName := value[2]

			number, _ := strconv.Atoi(colorValue)

			colorValues[colorName] = getMaxValue(colorValues[colorName], number)
		}
	}
	return colorValues
}

func getMaxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func partOne(lines []string) int {
	var idSum int

	for id, line := range lines {
		lineValues := getValues(line)

		shouldSum := true
		for color, value := range lineValues {
			if value > possibleConfigs[color] {
				shouldSum = false
			}
		}

		if shouldSum {
			idSum += id + 1
		}
	}
	return idSum
}

func partTwo(lines []string) int {
	var acc int

	for _, line := range lines {
		colorValues := getValues(line)
		curr := 0
		for _, value := range colorValues {
			if curr == 0 {
				curr = value
				continue
			}
			curr *= value
		}
		acc += curr
	}

	return acc
}
