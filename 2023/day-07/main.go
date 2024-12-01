package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

var powerMap = map[byte]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

var powerMapWithJokers = map[byte]int{
	'J': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'Q': 12,
	'K': 13,
	'A': 14,
}

type Hand struct {
	cards string
	bid   int
}

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func getHandType(hand Hand) HandType {
	countsByLabel := make(map[rune]int)
	for _, char := range hand.cards {
		countsByLabel[char]++
	}

	countsByCount := make(map[int]int)
	for _, count := range countsByLabel {
		countsByCount[count]++
	}

	if countsByCount[5] == 1 {
		return FiveOfAKind
	}

	if countsByCount[4] == 1 {
		return FourOfAKind
	}

	if countsByCount[3] == 1 && countsByCount[2] == 1 {
		return FullHouse
	}

	if countsByCount[3] == 1 {
		return ThreeOfAKind
	}

	if countsByCount[2] == 2 {
		return TwoPair
	}

	if countsByCount[2] == 1 {
		return OnePair
	}

	return HighCard
}

func getHandTypeWithJokers(hand Hand) HandType {
	countsByLabel := make(map[rune]int)
	mostCommonNonJLabel := '0'
	for _, char := range hand.cards {
		countsByLabel[char]++

		if char != 'J' && countsByLabel[char] > countsByLabel[mostCommonNonJLabel] {
			mostCommonNonJLabel = char
		}
	}

	if mostCommonNonJLabel != '0' {
		countsByLabel[mostCommonNonJLabel] += countsByLabel['J']
		countsByLabel['J'] = 0
	}

	countsByCount := make(map[int]int)
	for _, count := range countsByLabel {
		countsByCount[count]++
	}

	if countsByCount[5] == 1 {
		return FiveOfAKind
	}

	if countsByCount[4] == 1 {
		return FourOfAKind
	}

	if countsByCount[3] == 1 && countsByCount[2] == 1 {
		return FullHouse
	}

	if countsByCount[3] == 1 {
		return ThreeOfAKind
	}

	if countsByCount[2] == 2 {
		return TwoPair
	}

	if countsByCount[2] == 1 {
		return OnePair
	}

	return HighCard
}

func compareLabels(a, b byte, powerMap map[byte]int) int {
	return powerMap[a] - powerMap[b]
}

func compareHands(a, b Hand) int {
	handTypeA := getHandType(a)
	handTypeB := getHandType(b)

	if handTypeA != handTypeB {
		return int(handTypeA) - int(handTypeB)
	}

	for i := 0; i < len(a.cards); i++ {
		if a.cards[i] != b.cards[i] {
			return compareLabels(a.cards[i], b.cards[i], powerMap)
		}
	}

	return 0
}

func compareHandsWithJokers(a, b Hand) int {
	handTypeA := getHandTypeWithJokers(a)
	handTypeB := getHandTypeWithJokers(b)

	if handTypeA != handTypeB {
		return int(handTypeA) - int(handTypeB)
	}

	for i := 0; i < len(a.cards); i++ {
		if a.cards[i] != b.cards[i] {
			return compareLabels(a.cards[i], b.cards[i], powerMapWithJokers)
		}
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := make([]Hand, 0)

	for scanner.Scan() {
		line := scanner.Text()

		var cards string
		var bid int

		fmt.Sscanf(line, "%s %d", &cards, &bid)

		hands = append(hands, Hand{cards, bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHandsWithJokers(hands[i], hands[j]) < 0
	})

	ans := 0
	for i, hand := range hands {
		ans = ans + hand.bid*(i+1)

		fmt.Println(hand)
	}

	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
