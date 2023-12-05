package first

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nddq/aoc2023/common"
)

const MaxInt = int(^uint(0) >> 1)

type Two struct {
}

func (t *Two) Solve(puzzle *common.Puzzle) {
	total := 0
	matrix := common.ParseBytesToStringSliceMatrix(puzzle.Data)
	for _, line := range matrix {
		// slide a window forward in the line
		winSize := 5
		if len(line) <= winSize {
			winSize = len(line)
		}
		first := 0
		last := 0
		for i := 0; i <= len(line)-winSize; i++ {
			// get the window
			window := line[i : i+winSize]
			// get the number from the window
			num := getNumFromWindow(window)
			if num != -1 {
				first = num
				break
			}
		}
		// slide a window backward in the line
		for i := len(line) - winSize; i >= 0; i-- {
			// get the window
			window := line[i : i+winSize]
			// get the number from the window
			num := getNumFromWindowReverse(window)
			if num != -1 {
				last = num
				break
			}
		}
		res := first*10 + last
		common.WriteToFile("output.txt", fmt.Sprint(res)+"\n")
		total += res
	}
	fmt.Println(total)
}

func getNumFromWindow(window []string) int {
	numStr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	str := strings.Join(window, "")
	index := MaxInt
	numToReturn := -1
	for i, c := range str {
		digit, err := strconv.Atoi(string(c))
		if err == nil {
			if i < index {
				numToReturn = digit
				index = i
			}
		}
	}
	for _, num := range numStr {
		i := strings.Index(str, num)
		if i != -1 && i < index {
			numToReturn = convertStrToNum(num)
			index = i
		}
	}

	return numToReturn
}

func getNumFromWindowReverse(window []string) int {
	numStr := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	str := strings.Join(window, "")
	index := -1
	numToReturn := -1
	for i, c := range str {
		digit, err := strconv.Atoi(string(c))
		if err == nil && i > index { // later index == good in the reverse case
			numToReturn = digit
			index = i
		}
	}
	for _, num := range numStr {
		i := strings.Index(str, num)
		if i > index {
			numToReturn = convertStrToNum(num)
			index = i
		}
	}

	return numToReturn
}

func convertStrToNum(num string) int {
	switch num {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}
