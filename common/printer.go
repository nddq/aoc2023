package common

import "fmt"

func PrettyPrintStringsMatrix(matrix [][]string) {
	for i := range matrix {
		for j := range matrix[i] {
			fmt.Printf(matrix[i][j])
		}
		fmt.Printf("\n")
	}
}
