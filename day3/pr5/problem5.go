/*
You and the Elf eventually reach a gondola lift station; he says the gondola lift will take you up to the water source, but this is as far as he can bring you. You go inside.

It doesn't take long to find the gondolas, but there seems to be a problem: they're not moving.

"Aaah!"

You turn around to see a slightly-greasy Elf with a wrench and a look of surprise. "Sorry, I wasn't expecting anyone! The gondola lift isn't working right now; it'll still be a while before I can fix it." You offer to help.

The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

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

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func addDigit(old int, newDigit int) (int) {
	return 10*old + newDigit
}

type indexPair struct {
	x int
	y int
}

func buildAdjacents()([]indexPair) {
	var adj []indexPair
	adj = append(adj, indexPair{x: -1, y: -1})
	adj = append(adj, indexPair{x: -1, y:  0})
	adj = append(adj, indexPair{x: -1, y:  1})
	adj = append(adj, indexPair{x:  0, y: -1})
	adj = append(adj, indexPair{x:  0, y:  1})
	adj = append(adj, indexPair{x:  1, y: -1})
	adj = append(adj, indexPair{x:  1, y:  0})
	adj = append(adj, indexPair{x:  1, y:  1})
	return adj
}

func addPair(p1 indexPair, p2 indexPair) (indexPair) {
	return indexPair{x: p1.x + p2.x, y: p1.y + p2.y}
}

func getAdjacentIndices(current indexPair, max indexPair) ([]indexPair) {
	var pairs []indexPair

	adjs := buildAdjacents()
	for _, adj := range adjs {
		newPair := addPair(current, adj)
		if ((newPair.x < 0) || (newPair.y < 0) ||
		    (newPair.x >= max.x) || (newPair.y >= max.y)) {
			    continue
		}
		pairs = append(pairs, newPair)
	}
	return pairs
}

func checkAdjForSymbol(symbolMat [][]rune, curr indexPair, max indexPair) (bool) {
	adjs := getAdjacentIndices(curr, max)
	for _, adj := range adjs {
		sym := symbolMat[adj.x][adj.y]
		if (!unicode.IsDigit(sym) && sym != '.') {
			// We have an adjacent symbol
			return true
		}
	}
	return false
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
	
	// The sum of the part numbers we find are valid
	sum := 0

	// Internal representation of the puzzle input
	var symbolMat [][]rune
	var maxX,maxY int
	maxY = 0
	maxLines := 190
	lines := 0
	var input []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		input = append(input, line)
		chars := []rune(line)
		maxX = len(chars)
		symbolMat = append(symbolMat, chars)

		maxY += 1

		if (lines > maxLines) {
			break
		}
		lines += 1
	}
	fmt.Println("Dimensions: ", maxX, "x", maxY)
	dims := indexPair{x:maxX, y:maxY}

	// Now, we process it
	for i, runeLine := range symbolMat {
		inNumber := false
		isAdjacent := false
		currPartNumber := 0
		for j, r := range runeLine {
			curr := indexPair{x:i, y:j}
			if unicode.IsDigit(r) {
				fmt.Print(string(r))
				if inNumber {
					// Add to the current one
					currPartNumber = addDigit(currPartNumber, int(r-'0'))
				} else {
					// Start a new part number
					inNumber = true
					currPartNumber = int(r - '0')
				}
				if !isAdjacent {
					isAdjacent = checkAdjForSymbol(symbolMat, curr, dims)
				}
			} else {
				if isAdjacent {
					fmt.Println("New adj pn: ", currPartNumber)
					sum += currPartNumber
				}
				currPartNumber = 0
				inNumber = false
				isAdjacent = false
			}

		}
		if isAdjacent {
			fmt.Println("New adj pn: ", currPartNumber)
			sum += currPartNumber
		}
		currPartNumber = 0
	}

	// Print the result, of course!
	fmt.Println("Sum of part numbers: ", sum)
}

