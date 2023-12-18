package solution

func Part1(data string) int {
	patterns := parse(data)

	return calcSummation(patterns, solve1)
}

func calcSummation(patterns [][][]byte, solve func (pattern [][]byte) (int, bool)) int {
	sum := 0

	for _, p := range patterns {
		cnt, isRow := solve(p)
		if isRow {
			cnt = 100 * cnt
		}
		sum += cnt
	}
	return sum
} 

func solve1(pattern [][]byte) (int, bool) {
	rowCnt, colCnt := mapToCounts(pattern)
	if pos := findReflection(rowCnt); pos != -1 {
		return pos+1, true
	} 
	return findReflection(colCnt)+1, false
}

func mapToCounts(pattern [][]byte) ([]int, []int) {
	rowCnt := make([]int, len(pattern))
	colCnt := make([]int, len(pattern[0]))

	for i, r := range pattern {
		for j, v := range r {
			if v == '#' {
				rowCnt[i] |= 1 << j
				colCnt[j] |= 1 << i
			}
		}
	}
	/*
	Counts Result for example input
	RowCnt- [205 180 259 259 180 204 181] ColCnt- [77 12 115 33 82 82 33 115 12]
	RowCnt- [305 289 460 223 223 460 289] ColCnt- [91 24 60 60 25 67 60 60 103]
	*/
	return rowCnt, colCnt
}

func findReflection(cnts []int) int {
	for i := 1; i < len(cnts); i++ {
		if isReflection(cnts, i-1, i) {
			return i-1
		}
	}
	return -1
}

func isReflection(cnts []int, l, r int) bool {
	for l >= 0 && r < len(cnts) && cnts[l] == cnts[r] {
		l--
		r++
	}
	return l < 0 || r == len(cnts)
}