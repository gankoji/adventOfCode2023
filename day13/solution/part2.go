package solution

func Part2(data string) int {
	patterns := parse(data)

	return calcSummation(patterns, solve2)
}

func solve2(pattern [][]byte) (int, bool) {
	rowCnts, colCnts := mapToCounts(pattern)
	if pos := findNextReflection(rowCnts, len(colCnts)); pos != -1 {
		return pos+1, true
	}
	return findNextReflection(colCnts, len(rowCnts))+1, false
}

func findNextReflection(cnts []int, cntSize int) int {
	for i := 1; i < len(cnts); i++ {
		if !isReflection(cnts, i-1, i) && canBeReflection(cnts, i-1, i, cntSize) {
			return i-1
		}
	}
	return -1
}

func canBeReflection(cnts []int, l, r int, cntSize int) bool {
	for l >= 0 && r < len(cnts) && (cnts[l] == cnts[r] || differBy1(cnts[l], cnts[r], cntSize)) {
		l--
		r++
	}
	return l < 0 || r == len(cnts)
}

func differBy1(n1, n2 int, size int) bool {
	if n1 == n2 {
		return false
	}
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	// n1 small, n2 greater
	for i := 0; i < size; i++ {
		diff := 1 << i
		newN1 := n1 | diff
		
		if newN1 == n2 {
			return true
		}
	}
	return false
}