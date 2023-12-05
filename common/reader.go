package common

import (
	"bufio"
	"os"
)

func ReadPuzzle(path string) (*Puzzle, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Puzzle{Data: bytes}, nil
}

func ReadFileToGraph(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var graph [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		graph = append(graph, []rune(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}
