package first

import (
	"testing"

	"github.com/nddq/aoc2023/common"
)

func TestTwo(t *testing.T) {
	puzzle, err := common.ReadPuzzle("input.txt")
	if err != nil {
		panic(err)
	}
	one := &Two{}
	one.Solve(puzzle)
}
