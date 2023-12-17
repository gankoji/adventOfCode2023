/*
--- Part Two ---
Everyone will starve if you only plant such a small number of seeds. Re-reading the almanac, it looks like the seeds: line actually describes ranges of seed numbers.

The values on the initial seeds: line come in pairs. Within each pair, the first value is the start of the range and the second value is the length of the range. So, in the first line of the example above:

seeds: 79 14 55 13
This line describes two ranges of seed numbers to be planted in the garden. The first range starts with seed number 79 and contains 14 values: 79, 80, ..., 91, 92. The second range starts with seed number 55 and contains 13 values: 55, 56, ..., 66, 67.

Now, rather than considering four seed numbers, you need to consider a total of 27 seed numbers.

In the above example, the lowest location number can be obtained from seed number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77, temperature 45, humidity 46, and location 46. So, the lowest location number is 46.

Consider all of the initial seed numbers listed in the ranges on the first line of the almanac. What is the lowest location number that corresponds to any of the initial seed numbers?
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"pr10/v2/intervals"
	"strconv"
	"strings"
)

type intervalPair struct {
	src  intervals.Interval
	dest intervals.Interval
}

func makeIntervalPairSlice(buffer []string) []intervalPair {
	out := []intervalPair{}

	for _, line := range buffer {
		numStr := strings.Split(strings.TrimSpace(line), " ")
		dest, _ := strconv.Atoi(numStr[0])
		src, _ := strconv.Atoi(numStr[1])
		span, _ := strconv.Atoi(numStr[2])
		srcInt := intervals.MakeInt(src, src+span)
		destInt := intervals.MakeInt(dest, dest+span)
		out = append(out, intervalPair{srcInt, destInt})
	}
	return out
}

func transformIntervalWithReplacementPairs(a intervals.Interval, pairs []intervalPair) []intervals.Interval {
	out := []intervals.Interval{}
	for _, intPair := range pairs {
		inter := intervals.Intersect(a, intPair.src)
		if !intervals.IsEmpty(inter) {
			newInt := intervals.MakeInt(intPair.dest.Start, intPair.dest.Start + inter.Span)
			out = append(out, newInt)

			al := intervals.SetMinus(a, inter)
			for _, leftover := range al {
				if !intervals.IsEmpty(leftover) {
					out = append(out, leftover)
				}
			}
		} else {
			out = append(out, a)
		}
	}

	out = intervals.SortIntervals(out)

	return out
}

func main() {
	useExample := flag.Bool("example", false, "Whether to use the example input or the real problem input")
	flag.Parse()

	fileStr := "./input.txt"
	if *useExample {
		fmt.Println("Using example input")
		fileStr = "./exampleInput.txt"
	}

	file, err := os.Open(fileStr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0

	var mapBuffer []string
	mapBuffer = make([]string, 0)

	example := intervals.MakeInt(1,2)
	fmt.Println(example)
	pairs := [][]intervalPair{}

	seedInts := []intervals.Interval{}

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		fmt.Println(line)

		if row == 0 {
			// We have the seeds here, they're the different ones
			parts := strings.Split(line, ":")
			fmt.Println(parts)
			seedNumStrs := strings.Split(strings.TrimSpace(parts[1]), " ")
			for j := 0; j < len(seedNumStrs); j += 2 {
				startStr := seedNumStrs[j]
				spanStr := seedNumStrs[j+1]
				start, _ := strconv.Atoi(startStr)
				span, _ := strconv.Atoi(spanStr)
				seedInts = append(seedInts, intervals.MakeInt(start, (start + span)))
			}

			row++
			continue
		}

		if line == "" {
			if row == 1 {
				row++
				continue
			}

			newPairSlice := makeIntervalPairSlice(mapBuffer)
			pairs = append(pairs, newPairSlice)
			mapBuffer = make([]string, 0)
		} else if len(strings.Split(line, ":")) > 1 {
			continue
		} else {
			mapBuffer = append(mapBuffer, line)
		}

		row++

	}

	if len(mapBuffer) > 0 {
		newPairSlice := makeIntervalPairSlice(mapBuffer)
		pairs = append(pairs, newPairSlice)
	}

	minLocation := 0

	fmt.Println(len(seedInts))
	for _, pairList := range pairs {
		fmt.Println("Pair list: ", pairList)
		tempIntervals := []intervals.Interval{}

		for _, interval := range seedInts {
			// Apply the map to the interval and add to temp
			fmt.Println("New input interval: ", interval)
			outputInts := transformIntervalWithReplacementPairs(interval, pairList)
			fmt.Println("Output after transform: ", outputInts)
			for _, out := range outputInts {
				if !intervals.IsEmpty(out) {
					inTemp := false
					for _, pr := range tempIntervals {
						if  (pr.Start == out.Start) && (pr.End == out.End) {
							inTemp = true
						}
					}
					if len(tempIntervals) == 0 || !inTemp {
						tempIntervals = append(tempIntervals, out)
					}
				}
			}
			fmt.Println("Temp intervals: ", tempIntervals)
		}
		tempIntervals = intervals.SortIntervals(tempIntervals)
		seedInts = intervals.CompressOrderedSet(tempIntervals)
	}


	fmt.Println("Break")
	fmt.Println(seedInts[0])
	seedInts = intervals.SortIntervals(seedInts)
	fmt.Println(seedInts[0])

	minLocation = seedInts[0].Start
	fmt.Println("Minimum location: ", minLocation)

}
