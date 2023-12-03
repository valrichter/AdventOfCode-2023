package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	input, err := os.Open("day-01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var sum1 int
	var sum2 int

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		var arr1 []string
		var arr2 []string
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				arr1 = append(arr1, string(line[i]))
				arr2 = append(arr2, string(line[i]))
				continue
			}

			for str, val := range m {
				if strings.HasPrefix(line[i:], str) {
					arr2 = append(arr2, strconv.Itoa(val))
				}
			}
		}

		addToSum(&sum1, arr1)
		addToSum(&sum2, arr2)
	}

	output, err := os.Create("day-01/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	_, err2 := output.WriteString(fmt.Sprintf("Part 1: %d\nPart 2: %d", sum1, sum2))
	if err2 != nil {
		log.Fatal(err2)
	}
}

func addToSum(sum *int, arr []string) {
	val, err := strconv.Atoi(arr[0] + arr[len(arr)-1])
	if err != nil {
		log.Fatal(err)
	}

	*sum += val
}
