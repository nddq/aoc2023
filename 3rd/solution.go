package main

import (
	"fmt"
	"strconv"

	"github.com/nddq/aoc2023/common"
)

func search(i, j int, graph [][]rune) string {
	// out-of-bound
	if i < 0 || i >= len(graph) || j < 0 || j >= len(graph[0]) {
		return ""
	}
	// if not a digit
	if graph[i][j] < '0' || graph[i][j] > '9' {
		return ""
	}
	tmp := graph[i][j]
	graph[i][j] = '.'
	left := search(i, j-1, graph)
	right := search(i, j+1, graph)
	return left + string(tmp) + right

}

func SolveOne(graph [][]rune) {
	res := 0
	for i := range graph {
		for j := range graph[i] {
			if (graph[i][j] < '0' || graph[i][j] > '9') && graph[i][j] != '.' {
				// make a copy of the graph
				graphCopy := make([][]rune, len(graph))
				for i := range graph {
					graphCopy[i] = make([]rune, len(graph[i]))
					copy(graphCopy[i], graph[i])
				}

				left, right, up, down, upleft, upright, downleft, downright := search(i, j-1, graphCopy), search(i, j+1, graphCopy), search(i-1, j, graphCopy), search(i+1, j, graphCopy), search(i-1, j-1, graphCopy), search(i-1, j+1, graphCopy), search(i+1, j-1, graphCopy), search(i+1, j+1, graphCopy)
				fmt.Println(string(graph[i][j]))
				fmt.Printf("left: %s\n", left)
				fmt.Printf("right: %s\n", right)
				fmt.Printf("up: %s\n", up)
				fmt.Printf("down: %s\n", down)
				fmt.Printf("upleft: %s\n", upleft)
				fmt.Printf("upright: %s\n", upright)
				fmt.Printf("downleft: %s\n", downleft)
				fmt.Printf("downright: %s\n", downright)
				fmt.Println()
				// convert to int
				leftInt, _ := strconv.Atoi(left)
				rightInt, _ := strconv.Atoi(right)
				upInt, _ := strconv.Atoi(up)
				downInt, _ := strconv.Atoi(down)
				upleftInt, _ := strconv.Atoi(upleft)
				uprightInt, _ := strconv.Atoi(upright)
				downleftInt, _ := strconv.Atoi(downleft)
				downrightInt, _ := strconv.Atoi(downright)
				res += leftInt + rightInt + upInt + downInt + upleftInt + uprightInt + downleftInt + downrightInt
			}
		}
	}
	fmt.Println(res)
}

func SolveTwo(graph [][]rune) {
	res := 0
	for i := range graph {
		for j := range graph[i] {
			if (graph[i][j] < '0' || graph[i][j] > '9') && graph[i][j] == '*' {
				// make a copy of the graph
				graphCopy := make([][]rune, len(graph))
				for i := range graph {
					graphCopy[i] = make([]rune, len(graph[i]))
					copy(graphCopy[i], graph[i])
				}

				left, right, up, down, upleft, upright, downleft, downright := search(i, j-1, graphCopy), search(i, j+1, graphCopy), search(i-1, j, graphCopy), search(i+1, j, graphCopy), search(i-1, j-1, graphCopy), search(i-1, j+1, graphCopy), search(i+1, j-1, graphCopy), search(i+1, j+1, graphCopy)
				fmt.Println(string(graph[i][j]))
				fmt.Printf("left: %s\n", left)
				fmt.Printf("right: %s\n", right)
				fmt.Printf("up: %s\n", up)
				fmt.Printf("down: %s\n", down)
				fmt.Printf("upleft: %s\n", upleft)
				fmt.Printf("upright: %s\n", upright)
				fmt.Printf("downleft: %s\n", downleft)
				fmt.Printf("downright: %s\n", downright)
				fmt.Println()
				numOfAdjNumber := 0
				if left != "" {
					numOfAdjNumber++
				}
				if right != "" {
					numOfAdjNumber++
				}
				if up != "" {
					numOfAdjNumber++
				}
				if down != "" {
					numOfAdjNumber++
				}
				if upleft != "" {
					numOfAdjNumber++
				}
				if upright != "" {
					numOfAdjNumber++
				}
				if downleft != "" {
					numOfAdjNumber++
				}
				if downright != "" {
					numOfAdjNumber++
				}
				if numOfAdjNumber >= 2 {
					// convert to int
					leftInt, err := strconv.Atoi(left)
					if err != nil {
						leftInt = 1
					}
					rightInt, err := strconv.Atoi(right)
					if err != nil {
						rightInt = 1
					}
					upInt, err := strconv.Atoi(up)
					if err != nil {
						upInt = 1
					}
					downInt, err := strconv.Atoi(down)
					if err != nil {
						downInt = 1
					}
					upleftInt, err := strconv.Atoi(upleft)
					if err != nil {
						upleftInt = 1
					}
					uprightInt, err := strconv.Atoi(upright)
					if err != nil {
						uprightInt = 1
					}
					downleftInt, err := strconv.Atoi(downleft)
					if err != nil {
						downleftInt = 1
					}
					downrightInt, err := strconv.Atoi(downright)
					if err != nil {
						downrightInt = 1
					}
					res += leftInt * rightInt * upInt * downInt * upleftInt * uprightInt * downleftInt * downrightInt
				}
			}
		}
	}
	fmt.Println(res)
}

func main() {
	graph, err := common.ReadFileToGraph("input.txt")
	if err != nil {
		panic(err)
	}
	SolveTwo(graph)

}
