package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strings"
)

type Node struct {
	children []string
}

func allZees(nodes []string) bool {
	for _, node := range nodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func findStepsOfSinglePath(start string, dirs []int, graph map[string]Node) int {
	numSteps := 0
	dirIdx := -1

	curr := start
	for curr[2] != 'Z' {
		numSteps++
		dirIdx++

		if dirIdx >= len(dirs) {
			dirIdx = 0
		}

		// Take out direction at this step
		dir := dirs[dirIdx]

		// And use it to get the next node on the graph
		curr = graph[curr].children[dir]
	}
	
	return numSteps
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
			result = LCM(result, integers[i])
	}

	return result
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var useRealInput = flag.Bool("p", false, "Whether to use the real input.")

func main() {
	flag.Parse()

    if *cpuprofile != "" {
        f, err := os.Create(*cpuprofile)
        if err != nil {
            log.Fatal("could not create CPU profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        if err := pprof.StartCPUProfile(f); err != nil {
            log.Fatal("could not start CPU profile: ", err)
        }
        defer pprof.StopCPUProfile()
    }

	fileStr := "./exampleInput.txt"
	if *useRealInput {
		fmt.Println("Using real input")
		fileStr = "./input.txt"
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
			re := regexp.MustCompile(`[0-9A-Z]{3}`)
			nodeIds := re.FindAllString(line, -1)
			
			if nodeIds != nil {
				newOne := Node{[]string{nodeIds[1], nodeIds[2]}}
				graph[nodeIds[0]] = newOne
			}
		}


		lineNum++
	}

	// Special nodes in the graph that we're about to walk
	start := []string{}
	for k, _ := range graph {
		if k[2] == 'A' {
			start = append(start, k)
		}
	}
	fmt.Printf("Found start nodes: %s.\n", start)

	stepsPerPath := []int{}
	for _, node := range start {
		steps := findStepsOfSinglePath(node, dirs, graph)
		stepsPerPath = append(stepsPerPath, steps)
	}

	fmt.Printf("Found steps per path: %s.\n", stepsPerPath)

	lcm := stepsPerPath[0]
	for i:=1; i<len(stepsPerPath); i++ {
		lcm = LCM(lcm, stepsPerPath[i])
	}

	fmt.Printf("The answer is: %d.\n", lcm)
	if *memprofile != "" {
        f, err := os.Create(*memprofile)
        if err != nil {
            log.Fatal("could not create memory profile: ", err)
        }
        defer f.Close() // error handling omitted for example
        // runtime.GC() // get up-to-date statistics
        if err := pprof.WriteHeapProfile(f); err != nil {
            log.Fatal("could not write memory profile: ", err)
        }
    }
}
