package first

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/nddq/aoc2023/common"
)

type One struct {
}

func (o *One) Solve(puzzle *common.Puzzle) {
	totalCalibrationValue := 0
	matrix := common.ParseBytesToStringSliceMatrix(puzzle.Data)
	for _, line := range matrix {
		CalibrationValueStr := ""
		str := strings.Join(line, "")
		// Compile the regular expression
		re := regexp.MustCompile("[1-9]")

		// Find the first match
		firstDigit := re.FindString(str)
		CalibrationValueStr += firstDigit
		// reverse the string
		str = common.Reverse(str)
		// Find the first match
		lastDigit := re.FindString(str)
		CalibrationValueStr += lastDigit
		// convert to int
		i, err := strconv.Atoi(CalibrationValueStr)
		if err != nil {
			panic(err)
		}
		totalCalibrationValue += i
	}
	fmt.Println(totalCalibrationValue)
}
