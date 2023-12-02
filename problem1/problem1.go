package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("./input1.txt")
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		theString := []rune(line)

		first := 0
		last := 0
		lookingFirst := true
		for _, r := range theString {
			if unicode.IsDigit(r) {
				if lookingFirst {
					first = int(r-'0')
					lookingFirst = false
				}
				last = int(r-'0')
			}
		}
		_ = lookingFirst
		twoDigitNumber := 10.0*first + last
		fmt.Println("The two digit number for calibration: ", twoDigitNumber)
		sum += twoDigitNumber
	}

	fmt.Println("Total calibration: ", sum)

	if err := scanner.Err(); err != nil {
	    log.Fatal(err)
	}
}
