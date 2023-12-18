package solution

import "strings"

func parse(data string) [][][]byte {
	blocks := strings.Split(data, "\n\n")

	var patterns [][][]byte
	for _, b := range blocks {
		patterns = append(patterns, getPattern(b))
	}
	return patterns
}

func getPattern(block string) [][]byte {
	lines := strings.Split(block, "\n")

	var pattern [][]byte 
	for _, l := range lines {
		pattern = append(pattern, []byte(l))
	}
	return pattern
}