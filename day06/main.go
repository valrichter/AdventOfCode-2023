package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Race struct {
	time     int
	distance int
}

func main() {
	// Input
	lines := readInput("day06/input.txt")
	fmt.Println(lines)

	// Part 1
	part1 := Part1(lines)
	fmt.Println("Part 1: ", part1)

	// Part 2
	timeCode(lines)
	part2 := Part2(lines)
	fmt.Println("Part 2: ", part2)
}

func timeCode(lines []string) {
	// Registra el tiempo antes de ejecutar el código
	startTime := time.Now()

	// Código que deseas medir
	Part2(lines)

	// Registra el tiempo después de ejecutar el código
	endTime := time.Now()

	// Calcula la diferencia de tiempo
	elapsedTime := endTime.Sub(startTime)

	// Imprime el tiempo transcurrido
	fmt.Printf("Tiempo transcurrido: %s\n", elapsedTime)
}

func Part2(lines []string) int {
	time := lines[0]
	distance := lines[1]

	// Regex for numbers
	regex := regexp.MustCompile(`\d+`)

	// Find the numbers
	timeFind := regex.FindAllString(time, -1)
	distanceFind := regex.FindAllString(distance, -1)

	// Convert to int
	timeNumber := strings.Join(timeFind, "")
	distanceNumber := strings.Join(distanceFind, "")

	timeInt, _ := strconv.Atoi(timeNumber)
	distanceInt, _ := strconv.Atoi(distanceNumber)

	race := Race{
		time:     timeInt,
		distance: distanceInt,
	}
	totalWaysRace := totalWaysToWinTheRace(race)

	return totalWaysRace
}

func Part1(lines []string) int {
	// Set Races
	races := setRaces(lines)
	fmt.Println(races)

	// Total ways of all races
	numberOfWaysRecord := 1
	for i, race := range races {
		totalWaysRace := totalWaysToWinTheRace(race)
		numberOfWaysRecord = numberOfWaysRecord * totalWaysRace
		fmt.Println("Total ways in Race", i, ":", totalWaysRace)
	}

	return numberOfWaysRecord
}

func setRaces(lines []string) []Race {
	races := []Race{}
	times := lines[0]
	distances := lines[1]

	// Regex for numbers
	regex := regexp.MustCompile(`\d+`)

	// Find the numbers
	timesFinds := regex.FindAllString(times, -1)
	distancesFinds := regex.FindAllString(distances, -1)

	// Convert to int
	for i := 0; i < len(timesFinds); i++ {
		time, _ := strconv.Atoi(timesFinds[i])
		distance, _ := strconv.Atoi(distancesFinds[i])
		race := Race{
			time:     time,
			distance: distance,
		}
		races = append(races, race)
	}

	return races
}

func totalWaysToWinTheRace(race Race) int {
	totalWays := 0

	max := 0
	for holdTime := 1; holdTime <= race.time; holdTime++ {
		restDistance := race.time - holdTime
		newRecord := holdTime * restDistance

		if newRecord > race.distance {
			totalWays = totalWays + 1
		}

		// Quadratic solution
		max = newRecord
		if max > newRecord {
			return totalWays
		}

	}
	return totalWays
}

func readInput(file string) []string {
	rawFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawFile), "\n")
}
