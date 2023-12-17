/*
The Elf says they've stopped producing snow because they aren't getting any water! He isn't sure why the water stopped; however, he can show you how to get to the water source to check it out for yourself. It's just up ahead!

As you continue your walk, the Elf poses a second question: in each game you played, what is the fewest number of cubes of each color that could have been in the bag to make the game possible?

Again consider the example games from earlier:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
Game 4 required at least 14 red, 3 green, and 15 blue cubes.
Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.
The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together. The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum 2286.

For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)


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
	
	// This map holds our "total dice put into the bag," by color string
	var numberOfDice map[string]int
	numberOfDice = make(map[string]int)

	// The sum of the game IDs we find are valid
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Format of line is encoded into the splits
		// First is getting the game ID ahead of the colon
		res1 := strings.Split(line, ":")

		// Start with a small number, since we're looking for maxima
		numberOfDice["red"]   = 0
		numberOfDice["green"] = 0
		numberOfDice["blue"]  = 0

		// Next, split the game by round (semicolons)
		res2 := strings.Split(res1[1], ";")
		for _, round := range res2 {
			// In each round, we have a number of different color cubes found, separated by comma
			res3 := strings.Split(round, ",")
			for _, chunk := range res3 {
				// Each "color found" is a number and a word, separated by space
				res4 := strings.Split(strings.TrimSpace(chunk), " ")
				numCubesSeen, _ := strconv.Atoi(res4[0])
				colorSeen := strings.TrimSpace(res4[1])

				if numCubesSeen > numberOfDice[colorSeen] {
					numberOfDice[colorSeen] = numCubesSeen
				}
			}
		}

		power := 1
		for _, num := range numberOfDice {
			power *= num
		}

		// Add the power of the cube set for this game
		sum += power
	}

	// Print the result, of course!
	fmt.Println("Sum of game IDs: ", sum)
}

