package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	children []string
}

func Sum(arr []int) int {
	out := 0
	for _, el := range arr {
		out += el
	}
	return out
}

func getSeqFromLine(line string) []int {
	out := []int{}
	for _, numStr := range strings.Split(strings.TrimSpace(line), " ") {
		num, _ := strconv.Atoi(numStr)
		out = append(out, num)
	}
	return out
}

func getDiffSeq(arr []int) []int {
	out := []int{}
	for i:= 1; i<len(arr); i++ {
		out = append(out, arr[i] - arr[i-1])
	}
	return out
}

func allZeros(arr []int) bool {
	return Sum(arr) == 0
}

func main() {
	useRealInput := flag.Bool("real", false, "Whether to use the real problem input")
	flag.Parse()

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

	inputSeqs := [][]int{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		inputSeqs = append(inputSeqs, getSeqFromLine(line))

	}

	sampleSeq := 131
	fmt.Println(inputSeqs[sampleSeq])
	nextElements := []int{}
	for i, seq := range inputSeqs {
		tempSeqs := [][]int{}
		ogSeq := seq

		// Build a 'stack' (list we'll access backwards)
		// of the difference sequences until we hit 
		// an all zero seq
		for (len(seq) > 0) && !allZeros(seq) {
			tempSeqs = append(tempSeqs, seq)
			seq = getDiffSeq(seq)
			if i == sampleSeq {
				fmt.Println(seq)
			}
		}

		difference := 0
		for j:=len(tempSeqs)-1; j> 0; j-- {
			// Now, we pop the stack and sum
			delta := tempSeqs[j][len(tempSeqs[j])-1]
			if i == sampleSeq {
				fmt.Printf("Difference: %d, delta: %d, total: %d.\n", difference, delta, delta+difference)
			}
			difference += delta
		}

		nextElement := ogSeq[len(ogSeq)-1] + difference
		nextElements = append(nextElements, nextElement)
	}

	fmt.Printf("The next element in the sample sequence is: %d.\n", nextElements[sampleSeq])
	elSum := Sum(nextElements)
	fmt.Printf("Sum of the next elements is %d.\n", elSum)
}
