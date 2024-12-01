package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPoints := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		numbersPart := strings.TrimSpace(parts[1])

		numberSets := strings.Split(numbersPart, "|")
		numbersLeft := strings.Fields(strings.TrimSpace(numberSets[0]))
		numbersRight := strings.Fields(strings.TrimSpace(numberSets[1]))

		luckyNumbers := make(map[int]bool)
		for _, number := range numbersLeft {
			intNumber, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			luckyNumbers[intNumber] = true
		}

		currentPoints := 0
		for _, number := range numbersRight {
			intNumber, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			if _, ok := luckyNumbers[intNumber]; ok {
				if currentPoints == 0 {
					currentPoints = 1
				} else {
					currentPoints *= 2
				}
			}
		}

		totalPoints += currentPoints
	}

	log.Println("Total points:", totalPoints)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalCards := 0
	cards := make(map[int]int)
	idx := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		numbersPart := strings.TrimSpace(parts[1])

		numberSets := strings.Split(numbersPart, "|")
		numbersLeft := strings.Fields(strings.TrimSpace(numberSets[0]))
		numbersRight := strings.Fields(strings.TrimSpace(numberSets[1]))

		luckyNumbers := make(map[int]bool)
		for _, number := range numbersLeft {
			intNumber, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			luckyNumbers[intNumber] = true
		}

		cards[idx] += 1
		totalCards += cards[idx]

		matching := 0
		for _, number := range numbersRight {
			intNumber, err := strconv.Atoi(number)
			if err != nil {
				log.Fatal(err)
			}

			if _, ok := luckyNumbers[intNumber]; ok {
				matching += 1
			}
		}

		for i := 1; i <= matching; i++ {
			cards[idx+i] += cards[idx]
		}

		idx += 1

	}

	log.Println("Total scratchcards:", totalCards)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
