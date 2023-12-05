package second

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nddq/aoc2023/common"
)

type game struct {
	noOfRed   int
	noOfGreen int
	noOfBlue  int
}

func parseGame(line string) game {
	sets := strings.Split(line, ",")
	g := game{}
	for _, set := range sets {
		str := strings.Split(set, " ")
		switch str[len(str)-1] {
		case "red":
			noOfRed, err := strconv.Atoi(str[len(str)-2])
			if err != nil {
				panic("failed to parse number of red cubes")
			}
			g.noOfRed = noOfRed
		case "green":
			noOfGreen, err := strconv.Atoi(str[len(str)-2])
			if err != nil {
				panic("failed to parse number of green cubes")
			}
			g.noOfGreen = noOfGreen
		case "blue":
			noOfBlue, err := strconv.Atoi(str[len(str)-2])
			if err != nil {
				panic("failed to parse number of blue cubes")
			}
			g.noOfBlue = noOfBlue
		default:
			println("???")
		}
	}
	return g
}

func gamePossible(g game) bool {
	if g.noOfBlue > 14 || g.noOfGreen > 13 || g.noOfRed > 12 {
		return false
	}
	return true
}

func SolveOne(puzzle *common.Puzzle) {
	games := common.ParseBytesToStringSlices(puzzle.Data)
	fmt.Println(len(games))
	res := 0
	for _, g := range games {
		str := strings.Split(g, ":")
		id, err := strconv.Atoi(strings.Split(str[0], " ")[1])
		if err != nil {
			panic("failed to parse game id")
		}
		sets := str[1]
		possible := true
		for _, set := range strings.Split(sets, ";") {
			smallerGame := parseGame(set)
			if !gamePossible(smallerGame) {
				possible = false
				break
			}
		}
		if possible {
			fmt.Printf("Game %d possible\n", id)
			res += id
		}
	}
	fmt.Println(res)
}

func SolveTwo(puzzle *common.Puzzle) {
	games := common.ParseBytesToStringSlices(puzzle.Data)
	res := 0
	for _, g := range games {
		str := strings.Split(g, ":")
		_, err := strconv.Atoi(strings.Split(str[0], " ")[1])
		if err != nil {
			panic("failed to parse game id")
		}
		sets := str[1]
		var minRed, minGreen, minBlue int
		for _, set := range strings.Split(sets, ";") {
			smallerGame := parseGame(set)
			minRed = max(smallerGame.noOfRed, minRed)
			minGreen = max(smallerGame.noOfGreen, minGreen)
			minBlue = max(smallerGame.noOfBlue, minBlue)
		}
		pw := minRed * minGreen * minBlue
		res += pw
	}
	fmt.Println(res)
}
