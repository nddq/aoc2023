package main

import (
	"fmt"
	"strings"

	"github.com/nddq/aoc2023/common"
)

func SolveOne() {
	cards, err := common.ReadFileToStringSliceSingleSpace("input.txt")
	if err != nil {
		panic(err)
	}

	totalPoint := 0
	for _, card := range cards {
		points := 0
		numbers := strings.Split(strings.Split(card, ":")[1], "|")
		// strip front and back spaces for each item in numbers
		for i, number := range numbers {
			numbers[i] = strings.TrimSpace(number)
		}
		combine := strings.Join(numbers, " ")
		allNumbers := strings.Split(combine, " ")
		seen := make(map[string]bool)
		for _, number := range allNumbers {
			if _, ok := seen[number]; ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			} else {
				seen[number] = true
			}
		}
		totalPoint += points
	}
	fmt.Println(totalPoint)
}

func SolveTwo() {
	cards, err := common.ReadFileToStringSliceSingleSpace("input.txt")
	if err != nil {
		panic(err)
	}
	instanceOfCards := make(map[int]int)
	for i := range cards {
		instanceOfCards[i] = 1
	}
	for j, card := range cards {
		noOfDuplicate := 0
		numbers := strings.Split(strings.Split(card, ":")[1], "|")
		// strip front and back spaces for each item in numbers
		for i, number := range numbers {
			numbers[i] = strings.TrimSpace(number)
		}
		combine := strings.Join(numbers, " ")
		allNumbers := strings.Split(combine, " ")
		seen := make(map[string]bool)
		for _, number := range allNumbers {
			if _, ok := seen[number]; ok {
				noOfDuplicate++
			} else {
				seen[number] = true
			}
		}
		for k := 1; k <= noOfDuplicate; k++ {
			instanceOfCards[j+k] += instanceOfCards[j]
		}
	}
	// sum of all values in map
	totalCards := 0
	for _, v := range instanceOfCards {
		totalCards += v
	}
	fmt.Println(totalCards)

}

func main() {
	SolveTwo()
}
