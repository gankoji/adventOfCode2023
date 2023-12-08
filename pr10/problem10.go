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
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	start int
	end int
}

type intervalPair struct {
	src interval
	dest interval
}

func rangesOverlap(a interval, b interval) bool {
	if b.start < a.start  {
		c := a
		a = b
		b = c
	}

	return a.end > b.start
}

func union(a interval, b interval) interval {
	if b.start < a.start  {
		c := a
		a = b
		b = c
	}

	return interval{start:a.start, end:b.end}
}

func intersect(a interval, b interval) interval {
	// Assumes intervals overlap
	points := []int{}
	points = append(points, a.start)
	points = append(points, b.start)
	points = append(points, a.end)
	points = append(points, b.end)

	sort.Slice(points, func(i, j int) bool {
		return points[i] < points[j]
	})

	return interval{start:points[1], end:points[2]}
}
	
func accessMap(theMap map[int][]int, yourKey int) int {
	for src, v := range theMap {
		dest := v[0]
		span := v[1]
		if (yourKey >= src) && (yourKey < (src + span)) {
			delta := yourKey - src
			return dest + delta
		}
	}

	return yourKey
}

func makeMap(lines []string) map[int][]int {
	var outMap map[int][]int
	outMap = make(map[int][]int, 0)

	for _, line := range lines {
		nums := strings.Split(strings.TrimSpace(line), " ")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		num3, _ := strconv.Atoi(nums[2])

		outMap[num2] = append(outMap[num2], num1) 
		outMap[num2] = append(outMap[num2], num3)
	}

	return outMap
}

func makeIntervalPair(lines []string) []intervalPair {
	outPairs := []intervalPair{}
	for _, line := range lines {
		nums := strings.Split(strings.TrimSpace(line), " ")
		dest, _ := strconv.Atoi(nums[0])
		src, _ := strconv.Atoi(nums[1])
		span, _ := strconv.Atoi(nums[2])

		int1 := interval{start:src, end:src+span}
		int2 := interval{start:dest, end:dest+span}
		outPairs = append(outPairs, intervalPair{src:int1, dest:int2})
	}

	return outPairs
}

func makeMapTest() bool {
	exampleLine := "300 200 10"
	exampleLines := []string{}
	exampleLines = append(exampleLines, exampleLine)
	newMap := makeMap(exampleLines)
	fmt.Println(newMap)

	val1 := accessMap(newMap, 200)
	val2 := accessMap(newMap, 210)
	if val1 != 300 {
		fmt.Printf("Expected: 300, Got: %d\n", val1)
		fmt.Println("Access start fails.")
		return false
	}
	if val2 != 210 {
		fmt.Println("Non-covered access fails.")
		return false
	}
	return true
}

func minFind(arr []int) int {
	min := arr[0]
	for _, el := range arr {
		if el < min {
			min = el
		}
	}
	return min
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

	var maps []map[int][]int
	maps = make([]map[int][]int, 0)
	pairs := []interval

	seedNums := [][]int{}
	seedInts := []interval{}

	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		fmt.Println(line)

		if row == 0 {
			// We have the seeds here, they're the different ones
			parts := strings.Split(line, ":")
			fmt.Println(parts)
			seedNumStrs := strings.Split(strings.TrimSpace(parts[1]), " ")
			for j:= 0; j < len(seedNumStrs); j += 2 {
				startStr := seedNumStrs[j]
				spanStr := seedNumStrs[j+1]
				start, _ := strconv.Atoi(startStr)
				span, _ := strconv.Atoi(spanStr)
				var list []int
				list = append(list, start)
				list = append(list, span)
				seedNums = append(seedNums, list)
				seedInts = append(seedInts, interval{start:start, end:(start + span)})
			}

			row++
			continue
		}
		
		if line == "" {
			if row == 1 {
				row++
				continue
			}

			newMap := makeMap(mapBuffer)
			maps = append(maps, newMap)
			mapBuffer = make([]string, 0)
		} else if (len(strings.Split(line, ":")) > 1) {
			continue
		} else {
			mapBuffer = append(mapBuffer, line)
		}

		row++

	}
	
	if len(mapBuffer) > 0 {
		newMap := makeMap(mapBuffer)
		maps = append(maps, newMap)
	}

	minLocation := 1000000000000
	for _, seedNumRange := range seedNums {
		start := seedNumRange[0]
		span := seedNumRange[1]
		for i:= start; i < start + span; i++ {
			seedNum := i
			for _, theMap := range maps {
				seedNum = accessMap(theMap, seedNum)
			}
			location := seedNum
			if (location < minLocation) {
				minLocation = location
			}
		}
	}

	if(makeMapTest()) {
		fmt.Println("Tests passed.")
	} else {
		fmt.Println("Tests failed.")
	}

	fmt.Println("Minimum location: ", minLocation)

}

