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

func Sum(arr []int) int {
	out := 0
	for _, el := range arr {
		out += el
	}
	return out
}

func ReadLines(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}

	return strings.Split(strings.TrimSpace(string(content)), "\n")
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
	for i := 1; i < len(arr); i++ {
		out = append(out, arr[i]-arr[i-1])
	}
	return out
}

func allZeros(arr []int) bool {
	for _, el := range arr {
		if el != 0 {
			return false
		}
	}
	return true
}

func PredictNext(nums []int) int {
	next := 0
	cur := nums
	for {
		cur = getDiffSeq(cur)
		if allZeros(cur) {
			break
		}
		next += cur[len(cur)-1]
	}
	return nums[len(nums)-1] + next
}

func PredictPrev(nums []int) int {
	cur := nums
	sub := 0
	mul := -1
	for {
		cur = getDiffSeq(cur)
		if allZeros(cur) {
			break
		}
		sub += cur[0] * mul
		mul *= -1
	}
	return nums[0] + sub
}

func ParseNums(line string) []int {
	nums := []int{}
	for _, num := range strings.Split(line, " ") {
		if num == "" {
			continue
		}
		n, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
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

	part1_total := 0
	part2_total := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		seq := ParseNums(line)
		part1_total += PredictNext(seq)
		part2_total += PredictPrev(seq)
	}

	fmt.Printf("Sum of the next elements is %d.\n", part1_total)
	fmt.Printf("Sum of the previous elements is %d.\n", part2_total)
}
