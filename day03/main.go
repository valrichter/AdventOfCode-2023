package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readInput("day03/input.txt")
	matrix := getMatrix(lines)
	// printMatrix(matrix)
	sum := 0

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[0]); j++ {
			getAdjacentElementsOfNumbersRecursive(matrix, i, globalColMatrix)
			if containsSymbols(globalAdjacentElements) {
				// fmt.Println("Numero: ", globalNumber)
				// fmt.Println("Elementos adyacentes", globalAdjacentElements)
				numInt, _ := strconv.Atoi(globalNumber)
				sum = sum + numInt
			}

			globalNumber = ""
			globalAdjacentElements = []string{}
			globalColMatrix = globalColMatrix + 1
		}

		globalNumber = ""
		globalAdjacentElements = []string{}
		globalColMatrix = 0

	}

	fmt.Println("Part 1: ", sum)
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
}
