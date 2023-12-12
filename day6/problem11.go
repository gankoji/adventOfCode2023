package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
    "regexp"
	"strings"
	"strconv"
)

func simulateRace(length int, record int) []int {
	out := []int{}

	for i := 0; i <= length; i++ {
		speed := i
		timeLeft := length - speed
		distance := timeLeft * speed
		if distance > record {
			out = append(out, distance)
		}
	} 

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

	lineNum := 0
	times := []int{}
	dists := []int{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		numbers := strings.Split(line, ":")[1]
		parts := strings.Split(numbers, " ")

        numeric := regexp.MustCompile(`\d`)

		for _, numStr := range parts {
            if !numeric.MatchString(numStr) {
               continue
            }
			num, _ := strconv.Atoi(numStr)
			if lineNum == 0 {
				times = append(times, num)
			} else {
				dists = append(dists, num)
			}
		}
		lineNum += 1
	}

	output := []int{}
	prod := 1
	for i := 0; i < len(times); i++ {
		tmp := simulateRace(times[i], dists[i])
		output = append(output, len(tmp))
		prod = prod * len(tmp)
	}

    fmt.Println("Answer to part 1: ", prod)

    // Part 2: combine all times/dists and try again
    timeStr := ""
    distStr := ""
    for i, num := range times {
        timeStr = fmt.Sprintf("%s%d", timeStr, num)
        distStr = fmt.Sprintf("%s%d", distStr, dists[i])
    }
    time, _ := strconv.Atoi(timeStr)
    dist, _ := strconv.Atoi(distStr)
    fmt.Println(timeStr, " ", distStr)
    fmt.Println("Answer to part 2: ", len(simulateRace(time, dist)))

 }
