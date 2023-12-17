package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func makeLengths(line string) []int {
	out := []int{}
	numStrs := strings.Split(line, ",")
	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		out = append(out, num)
	}

	return out
}

func readInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	out := []string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		out = append(out, line)
	}

	return out
}

func parseLinesToStrings(lines []string) ([]string, [][]int) {
	records := []string{}
	lengths := [][]int{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		records = append(records, parts[0])
		ls := makeLengths(parts[1])
		lengths = append(lengths, ls)
	}

	return records, lengths
}

func unfoldLength(lengths [][]int) [][]int {
	lout := [][]int{}
	for _, length := range lengths {
		newLength := []int{}
		for i := 0; i < 4; i++ {
			newLength = append(newLength, length...)
		}
		newLength = append(newLength, length...)
		lout = append(lout, newLength)
	}

	return lout
}

func unfoldString(record string) string {
	newString := ""

	for i := 0; i < 4; i++ {
		newString = newString + record + "?"
	}
	newString = newString + record
	return newString
}

func unfoldRecords(records []string) []string {
	newRecords := []string{}
	for _, record := range records {
		newRecords = append(newRecords, unfoldString(record))
	}
	return newRecords
}

func get(m map[int]int, key int, def int) int {
	if v, ok := m[key]; ok {
		return v
	} else {
		return def
	}
}

func matches(record string, ls []int) int {
	// This builds the list of possible states
	// i.e. the state representation of the NFA
	states := "."
	for _, num := range ls {
		for i := 0; i < num; i++ {
			states += "#"
		}
		states += "."
	}

	// We need two maps to keep track of the states that we see
	states_map := map[int]int{0: 1}
	new_map := map[int]int{}
	for _, char := range record {
		for state := range states_map {
			if char == '?' {
				if state+1 < len(states) {
					new_map[state+1] = get(new_map, state+1, 0) + states_map[state]
				}
				if states[state] == '.' {
					new_map[state] = get(new_map, state, 0) + states_map[state]
				}
			} else if char == '.' {
				if state+1 < len(states) && states[state+1] == '.' {
					new_map[state+1] = get(new_map, state+1, 0) + states_map[state]
				}
				if states[state] == '.' {
					new_map[state] = get(new_map, state, 0) + states_map[state]
				}
			} else if char == '#' {
				if state+1 < len(states) && states[state+1] == '#' {
					new_map[state+1] = get(new_map, state+1, 0) + states_map[state]
				}
			}
		}
		states_map = new_map
		new_map = make(map[int]int, 0)
	}
	l := len(states)
	out := get(states_map, l-1, 0) + get(states_map, l-2, 0)
	return out
}

func main() {
	useRealInput := flag.Bool("p", false, "Whether to use the example input or the real problem input")
	flag.Parse()

	fileStr := "./exampleInput.txt"
	if *useRealInput {
		fmt.Println("Using real input")
		fileStr = "./input.txt"
	}

	start := time.Now()

	lines := readInput(fileStr)
	records, lengths := parseLinesToStrings(lines)

	sum := 0
	for i, rec := range records {
		l := lengths[i]
		sum += matches(rec, l)
	}
	dur := time.Since(start)
	fmt.Println("Part 1: ", sum, " took ", dur)
	start = time.Now()

	records = unfoldRecords(records)
	lengths = unfoldLength(lengths)
	sum = 0
	for i, rec := range records {
		l := lengths[i]
		sum += matches(rec, l)
	}
	dur = time.Since(start)
	fmt.Println("Part 2: ", sum, " took ", dur)

}
