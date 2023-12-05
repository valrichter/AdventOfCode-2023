package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/Goldziher/go-utils/sliceutils"
)

type Card struct {
	card    string
	winning []int
	numbers []int
}

type Cards []Card

func main() {
	var cards Cards
	file, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := strings.Split(scanner.Text(), ": ")
		split := strings.Split(card[1], " | ")
		numbers := sortNums(split[1])
		winning := sortNums(split[0])
		cardInfo := Card{
			card:    card[0],
			winning: winning,
			numbers: numbers,
		}
		cards = append(cards, cardInfo)
	}

	total := 0

	for _, val := range cards {
		points := 0
		for _, int := range val.numbers {
			if slices.Contains(val.winning, int) && points == 0 {
				points = 1
			} else if slices.Contains(val.winning, int) {
				points = points * 2
			}
		}
		total = total + points
	}

	fmt.Println(total)
	part2()
}

func sortNums(numStr string) []int {
	var ints []int
	numSlice := strings.Fields(numStr)
	for i := range numSlice {
		num, err := strconv.Atoi(numSlice[i])
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, num)
	}
	sort.Ints(ints)
	return ints
}

func part2() {
	lines := countLines("day04/input.txt")
	file, err := os.Open("day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var played []int

	for i := 0; i < lines; i++ {
		played = append(played, 0)
	}

	index := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		played[index] += 1

		game := strings.Split(scanner.Text(), ": ")[1]

		winning := strings.Split(game, " | ")[0]

		ours := strings.Split(game, " | ")[1]

		winnum := numSlice(winning)
		ournum := numSlice(ours)

		intersection := sliceutils.Intersection(winnum, ournum)

		for w := range intersection {
			played[index+w+1] += played[index]
		}

		index++
	}
	total := 0
	for i := range played {
		total += played[i]
	}
	fmt.Println(total)
}

func countLines(file string) int {
	lines := 0
	content, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	scan := bufio.NewScanner(content)
	for scan.Scan() {
		lines++
	}
	return lines
}

func numSlice(nss string) []int {
	numstrsl := strings.Fields(nss)

	var numintsl []int

	for _, n := range numstrsl {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		numintsl = append(numintsl, num)
	}

	return numintsl
}
