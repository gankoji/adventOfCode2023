package intervals

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
	Span  int
	Empty bool
}

func SortIntervals(arr []Interval) []Interval {
	out := arr
	sort.Slice(out, func(i,j int) bool {

		startEqual := out[i].Start == out[j].Start
		starts := out[i].Start < out[j].Start
		ends := out[i].End < out[j].End

		if startEqual {
			return ends
		} else {
			return starts
		}
	})

	return out
}

func MakeInt(start int, end int) Interval {
	return Interval{start, end, end - start, false}
}

func MakeEmptyInt() Interval {
	return Interval{0, 0, 0, true}
}

func IsEmpty(a Interval) bool {
	return a.Empty
}

func Overlap(a Interval, b Interval) bool {
	if (a.Start < b.Start) {
		return a.End > b.Start
	} else {
		return b.End < a.Start
	}
}

func Intersect(a Interval, b Interval) Interval {
	if Overlap(a,b) {
		endpoints := []int{a.Start, a.End, b.Start, b.End}
		sort.Ints(endpoints)

		return MakeInt(endpoints[1], endpoints[2])
	} else {
		return MakeEmptyInt()
	}
}

func Union(a Interval, b Interval) []Interval {
	out := []Interval{}
	if Overlap(a,b) {
		endpoints := []int{a.Start, a.End, b.Start, b.End}
		sort.Ints(endpoints)

		out = append(out, MakeInt(endpoints[0], endpoints[3]))
	} else {
		out = append(out, a)
		out = append(out, b)
	}

	return out
}

func SetMinus(a Interval, b Interval) []Interval {
	out := []Interval{}

	if Overlap(a,b) {
		if (b.Start <= a.Start) {
			if (b.End >= a.End) {
				out = append(out, MakeEmptyInt())
			} else {
				out = append(out, MakeInt(b.End, a.End))
			}
		} else {
			out = append(out, MakeInt(a.Start, b.Start))
			if (b.End < a.End) {
				out = append(out, MakeInt(b.End, a.End))
			}
		}
	} else {
		out = append(out, a)
	}

	return out
}

func RemoveItemFromOrderedSet(arr []Interval, idx int) []Interval {
	return append(arr[:idx], arr[idx+1:]...)
}

func SinglePassCompress(arr []Interval) []Interval {
	temp := arr

	for i, pr := range temp {
		for j := i; j < len(temp); j++ {
			if Overlap(pr, temp[j]) {
				// Combine the two
				b := temp[j]
				temp = RemoveItemFromOrderedSet(temp, j)
				temp[0] = Union(pr,b)[0]
			}
		}
	}

	return temp
}
func CompressOrderedSet(a []Interval) []Interval {
	temp := a

	oldLen := 0
	newLen := len(temp)
	for (newLen - oldLen) > 0 {
		oldLen = newLen
		temp = SinglePassCompress(temp)
		temp = SortIntervals(temp)
		newLen = len(temp)
	}

	return temp
}

