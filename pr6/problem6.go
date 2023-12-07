/*
--- Part Two ---
The engineer finds the missing part and installs it in the engine! As the engine springs to life, you jump in the closest gondola, finally ready to ascend to the water source.

You don't seem to be going very fast, though. Maybe something is still wrong? Fortunately, the gondola has a phone labeled help, so you pick it up and the engineer answers.

Before you can explain the situation, she suggests that you look out the window. There stands the engineer, holding a phone in one hand and waving with the other. You're going so slowly that you haven't even left the station. You exit the gondola.

The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

Consider the same engine schematic again:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

What is the sum of all of the gear ratios in your engine schematic?
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type gearRatio struct {
	num1 int
	num2 int
	ratio int
}

type coords struct {
	x int
	y int
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
	
	var gears []gearRatio
	var input []string

	for scanner.Scan() {
		// Scan through line, looking for '*'
		// Once found, check (up to) 8 adjacent spaces for numbers
		// If 2, expand those numbers fully
		// Append to gears
		line := strings.TrimSpace(scanner.Text())
		input = append(input, line)
	}
	maxI := len(input)
	maxJ := len(input[0])

	// Now, we process it
	for i, line := range input {
		for j, s := range line {
			if int(rune(s) - '*') == 0 {
				// We've found a star, now to look for numbers
				// Setup the bounding adjacent box coords (x0, y0) -> (xf, yf)
				x0 := j
				xf := j
				y0 := i
				yf := i

				if (i == 0) {
					y0 = 0
				} else {
					y0 = i - 1
				}

				if (i == maxI) {
					yf = i
				} else {
					yf = i + 1
				}

				if (j==0){
					x0 = 0
				} else {
					x0 = j - 1
				}

				if (j == maxJ) {
					xf = j
				} else {
					xf = j + 1
				}

				// Search adjacent characters for digits
				// With some deduplication for digits that are adjacent to each other
				y := y0
				var searchPoints []coords
				searchPoints = make([]coords, 0)

				for y <= yf {
					newLine := input[y]
					isAdjacent := false
					x := x0
					for x <= xf {
						if unicode.IsDigit(rune(newLine[x])) && !isAdjacent {
							searchPoints = append(searchPoints, coords{x:x, y:y})
							isAdjacent = true
						} else if !unicode.IsDigit(rune(newLine[x])){
							isAdjacent = false
						}
						x++
					}
					y++
				}

				// Check that we have the correct number of part numbers for a gear
				if len(searchPoints) < 2 {
					fmt.Println("Only one number found")
					continue
				}

				if len(searchPoints) > 2 {
					fmt.Println("Somehow we have three or more numbers, this is undefined.")
					continue
				}

				// Parse those numbers and get the gear ratio
				prod := 1
				for _, point := range searchPoints {
					searchLine := input[point.y]
					searchStart := point.x
					searchEnd := point.x
					for (searchStart > 0 && unicode.IsDigit(rune(searchLine[searchStart-1]))) {
						searchStart--
					}
					for (searchEnd < (maxJ) && unicode.IsDigit(rune(searchLine[searchEnd]))) {
						searchEnd++
					}
					num, _ := strconv.Atoi(searchLine[searchStart:searchEnd])
					fmt.Println("After searching, found number: ", num, " from string: ", searchLine[searchStart:searchEnd])
					prod = prod * num
				}
				fmt.Println("This gear has ratio: ", prod)
				gears = append(gears, gearRatio{num1:1, num2:1, ratio:prod})
			}
		}

	}

	sum := 0
	for _, gear := range gears {
		sum += gear.ratio
	}

	// Print the result, of course!
	fmt.Println("Sum of gear ratios: ", sum)
}


