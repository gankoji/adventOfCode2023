package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	children []string
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
	dirs := []int{}
	graph := make(map[string]Node, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if lineNum == 0 {
			dirStrs := []rune(line)
			for _, dirRune := range dirStrs {
				if (dirRune == 'R') {
					dirs = append(dirs, 1)
				} else {
					dirs = append(dirs, 0)
				}
			}
		} else if lineNum > 1 {
			re := regexp.MustCompile(`\w\w\w`)
			nodeIds := re.FindAllString(line, -1)
			
			if nodeIds != nil {
				newOne := Node{[]string{nodeIds[1], nodeIds[2]}}
				graph[nodeIds[0]] = newOne
			}
		}


		lineNum++
	}


	fmt.Println(graph)
	
	// Special nodes in the graph that we're about to walk
	start := "AAA"
	end := "ZZZ"
	numSteps := 0
	dirIdx := -1

	curr := start
	for curr != end {
		numSteps++
		dirIdx++

		if dirIdx >= len(dirs) {
			dirIdx = 0
		}

		// Take out direction at this step
		dir := dirs[dirIdx]

		// And use it to get the next node on the graph
		prev := curr
		curr = graph[curr].children[dir]

		// For debugging purposes, let's print out our decisions. 
		fmt.Printf("At node: %s, going to node: %s, found direction: %d. \n", prev, curr, dir)
	}
	
	fmt.Printf("Got to finish node ZZZ in %d steps.\n", numSteps)
}
