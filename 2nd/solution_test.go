package second

import (
	"testing"

	"github.com/nddq/aoc2023/common"
)

func TestSolveOne(t *testing.T) {
	puzzle, err := common.ReadPuzzle("input.txt")
	if err != nil {
		panic(err)
	}
	SolveOne(puzzle)
}

func TestSolveTwo(t *testing.T) {
	puzzle, err := common.ReadPuzzle("input.txt")
	if err != nil {
		panic(err)
	}
	SolveTwo(puzzle)
}
