package common

import (
	"bufio"
	"os"
	"regexp"
)

func ReadPuzzle(path string) (*Puzzle, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return &Puzzle{Data: bytes}, nil
}

func ReadFileToStringSliceSingleSpace(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	reg := regexp.MustCompile("\\s+")

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		modifiedLine := reg.ReplaceAllString(line, " ")
		lines = append(lines, modifiedLine)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
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
