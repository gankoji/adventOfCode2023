package main

import (
	"fmt"
	"log"
	"bufio"
	"os"
	"regexp"
	"unicode"
)


func main() {
	type replacers struct {
		regex *regexp.Regexp
		repl string
	}
	var rs [10]replacers
	rs[0] = replacers{ regex: regexp.MustCompile("zero"), repl: "zero0zero"}
	rs[1] = replacers{ regex: regexp.MustCompile("one"),  repl: "one1one"}
	rs[2] = replacers{ regex: regexp.MustCompile("two"),  repl: "two2two"}
	rs[3] = replacers{ regex: regexp.MustCompile("three"),repl: "three3three"}
	rs[4] = replacers{ regex: regexp.MustCompile("four"), repl: "four4four"}
	rs[5] = replacers{ regex: regexp.MustCompile("five"), repl: "five5five"}
	rs[6] = replacers{ regex: regexp.MustCompile("six"),  repl: "six6six"}
	rs[7] = replacers{ regex: regexp.MustCompile("seven"),repl: "seven7seven"}
	rs[8] = replacers{ regex: regexp.MustCompile("eight"),repl: "eight8eight"}
	rs[9] = replacers{ regex: regexp.MustCompile("nine"), repl: "nine9nine"}

	file, err := os.Open("./input.txt")
	if err != nil {
	    log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		for _, replacer := range rs {
			line = replacer.regex.ReplaceAllString(line, replacer.repl)
		}

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
