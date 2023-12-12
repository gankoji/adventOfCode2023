package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"
	"strconv"
)

const (
	HighCard = iota
	OnePair = iota
	TwoPair = iota
	ThreeOfAKind = iota
	FullHouse = iota
	FourOfAKind = iota
	FiveOfAKind = iota
)

var CardToValue = map[rune]int{
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'J': 9,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type hand struct {
	cards string
	bid int
}

type Game []hand

func (a Game) Len() int 		{ return len(a) }
func (a Game) Swap(i,j int) 		{ a[i], a[j] = a[j], a[i] }
func (a Game) Less(i,j int) bool 	{ return compareHands(a[i], a[j]) }

func handHistogram(hand string) map[rune]int {
	chars := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
	out := make(map[rune]int, 0)
	for _, char := range chars {
		out[char] = 0
		for _, card := range []rune(hand) {
			if card == char {
				out[char] += 1
			}
		}
	}

	return out
}

func getQty(hist map[rune]int) []int {
	out := []int{}

	for _, v := range hist {
		if v > 0 {
			out = append(out, v)
		}
	}
	sort.Ints(out)
	slices.Reverse(out)
	return out
}

func findHandType(hand string) int {
	h := handHistogram(hand)
	qtys := getQty(h)

	if qtys[0] == 5 {
		return FiveOfAKind
	} else if qtys[0] == 4 {
		return FourOfAKind
	} else if qtys[0] == 3 {
		if qtys[1] == 2 {
			return FullHouse
		} else {
			return ThreeOfAKind
		}
	} else if qtys[0] == 2 {
		if qtys[1] == 2 {
			return TwoPair
		} else {
			return OnePair
		}
	} else {
		return HighCard
	}
}

func compareCards(a hand, b hand) bool {
	for i:=0; i<len(a.cards); i++ {
		runea := rune(a.cards[i])
		runeb := rune(b.cards[i])

		if runea == runeb {
			continue
		}
		return CardToValue[runea] < CardToValue[runeb]
	}

	return false
}

func compareHands(a hand, b hand) bool {
	typea := findHandType(a.cards)
	typeb := findHandType(b.cards)
	
	if typea == typeb {
		return compareCards(a, b)
	}

	return typea < typeb
}

func scoreGame(hands []hand) int {
	out := 0
	for i, hand := range hands {
		out += hand.bid*(i+1)
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

	hands := []hand{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, " ")

		bid, _ := strconv.Atoi(parts[1])

		hands = append(hands, hand{cards:parts[0], bid:bid})
	}

	fmt.Println(hands)
	sort.Sort(Game(hands))
	fmt.Println(hands)
	fmt.Println(scoreGame(hands))

	// Examples
	// fmt.Println("Examples")
	// hand1 := "KKQQQ"
	// hand2 := "22456"
	// h1 := hand{cards:hand1, bid: 20}
	// h2 := hand{cards:hand2, bid: 30}
	// fmt.Println(compareHands(h1, h2))

	// fmt.Println(hand1, findHandType(hand1))
	// fmt.Println(hand2, findHandType(hand2))
	// fmt.Println(findHandType(hand1) > findHandType(hand2))
	// for _, card := range hand1 {
	// 	val := CardToValue[card]
	// 	bigger := val > CardToValue['Q']
	// 	fmt.Println(card, " has val: ", val, ". Is bigger than Q? ", bigger)
	// }
}

