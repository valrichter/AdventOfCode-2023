package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Part 2

type NumberArea struct {
	adyacents []string
	number    int
}

func main() {
	lines := readInput("day03/input.txt")
	matrix := getMatrix(lines)
	printMatrix(matrix)
	sum := 0

	// Part 2
	allNumberAreas := []NumberArea{}

	// Part 2

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[0]); j++ {
			getAdjacentElementsOfNumbersRecursive(matrix, i, globalColMatrix)

			if containsSymbols(globalAdjacentElements) {
				numInt, _ := strconv.Atoi(globalNumber)
				sum = sum + numInt

				// Part 2

				numArea := NumberArea{
					adyacents: globalAdjacentElements,
					number:    numInt,
				}
				allNumberAreas = append(allNumberAreas, numArea)
				// Part 2
			}

			globalNumber = ""
			globalAdjacentElements = []string{}
			globalColMatrix = globalColMatrix + 1
		}

		globalNumber = ""
		globalAdjacentElements = []string{}
		globalColMatrix = 0
	}

	// Part 2
	fmt.Println("allNumberAreas:")
	for _, numArea := range allNumberAreas {
		fmt.Println(numArea.number, " : ", numArea.adyacents)
	}
	fmt.Println()

	// Data filtering
	// Definition of regex
	gearMap := make(map[string]string)
	regexPosition := regexp.MustCompile(`\((\d+,\d+)\)`)
	for _, numArea := range allNumberAreas {
		// elems adyacents
		inputString := strings.Join(numArea.adyacents, "")

		// Find matches
		match := regexPosition.FindString(inputString)

		// Map de gears values wiht pattern
		if match != "" {
			newKey := match
			newValue := strconv.Itoa(numArea.number)

			if gearMap[newKey] != "" {
				gearMap[newKey] = gearMap[newKey] + "*" + newValue
			} else {
				gearMap[newKey] = newValue
			}
		}
	}

	fmt.Println("gearMap:")
	for key, value := range gearMap {
		fmt.Println(key, " : ", value)
	}
	fmt.Println()

	var allGears []int
	// Filter and multiply by the values required
	for key, value := range gearMap {
		if strings.Contains(value, "*") {
			regexNum := regexp.MustCompile(`\d+`)
			coincidences := regexNum.FindAllString(value, -1)

			result := 1
			for _, numero := range coincidences {
				num, _ := strconv.Atoi(numero)
				result = result * num
			}
			allGears = append(allGears, result)
			fmt.Println(key, coincidences)

		} else {
			delete(gearMap, key)
		}
	}
	fmt.Println(allGears)

	// Sum all gears (multiplcations)
	gears := 0
	fmt.Println()
	for _, numero := range allGears {
		gears += numero
	}

	fmt.Println()
	fmt.Println("Part 1: ", sum)
	fmt.Println("Part 2: ", gears)
}

var globalAdjacentElements []string
var globalNumber string
var globalColMatrix int

func getAdjacentElementsOfNumbersRecursive(matrix [][]string, i, j int) {

	if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) && isNumber(matrix[i][j]) {
		elements := getAdjacentElements(matrix, i, j)
		globalAdjacentElements = append(globalAdjacentElements, elements...)
		globalNumber = globalNumber + matrix[i][j]
		globalColMatrix = j + 1
		getAdjacentElementsOfNumbersRecursive(matrix, i, j+1)
	}

}

func containsSymbols(arr []string) bool {
	for _, elem := range arr {
		if elem != "." && !isNumber(elem) && elem != " " {
			return true
		}
	}
	return false
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getAdjacentElements(matrix [][]string, i, j int) []string {
	adjacent := []string{}

	// Coordenadas relativas de los elementos adyacentes
	coordinates := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 0}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, coord := range coordinates {
		newI, newJ := i+coord[0], j+coord[1]

		// Verificar que las coordenadas estén dentro de los límites de la matriz
		if newI >= 0 && newI < len(matrix) && newJ >= 0 && newJ < len(matrix[0]) {
			elem := matrix[newI][newJ]

			// Part 2
			if elem == "*" {
				newi := strconv.Itoa(newI)
				newj := strconv.Itoa(newJ)
				pos := "(" + newi + "," + newj + ")"
				adjacent = append(adjacent, pos)
			}
			// Part 2

			adjacent = append(adjacent, elem)
		} else {
			adjacent = append(adjacent, " ")
		}
	}

	return adjacent
}

func getMatrix(line []string) [][]string {
	matrix := make([][]string, len(line))

	for i := 0; i < len(line); i++ {
		matrix[i] = readChars(line[i])
	}
	return matrix
}

func readChars(line string) []string {
	return strings.Split(line, "")
}

func readInput(file string) []string {
	rawFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(rawFile), "\n")
}

func printMatrix(matrix [][]string) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}
