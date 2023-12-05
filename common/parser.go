package common

func ParseBytesToStringSliceMatrix(bytes []byte) [][]string {
	var matrix [][]string
	var row []string
	for _, b := range bytes {
		if b == '\n' {
			matrix = append(matrix, row)
			row = []string{}
		} else {
			row = append(row, string(b))
		}
	}
	return matrix
}

func ParseBytesToStringSlices(bytes []byte) []string {
	var result []string
	var line string
	for _, b := range bytes {
		if b == '\n' {
			result = append(result, line)
			line = ""
		} else {
			line += string(b)
		}
	}
	return result
}

func Reverse(s string) string {
	var result string
	for _, v := range s {
		result = string(v) + result
	}
	return result
}
