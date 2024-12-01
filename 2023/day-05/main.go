package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

type MapList []Map

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	seedsSplice := strings.Split(strings.TrimPrefix(scanner.Text(), "seeds: "), " ")
	seeds := make([]int, len(seedsSplice))
	for i, seed := range seedsSplice {
		seeds[i], err = strconv.Atoi(seed)
		if err != nil {
			log.Fatal(err)
		}
	}

	scanner.Scan()
	scanner.Scan()

	var mappings []MapList
	var currentMapList MapList

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			mappings = append(mappings, currentMapList)
			currentMapList = make(MapList, 0)

			scanner.Scan()
			continue
		}

		line := scanner.Text()

		var destinationRangeStart int
		var sourceRangeStart int
		var rangeLength int
		_, err = fmt.Sscanf(line, "%d %d %d", &destinationRangeStart, &sourceRangeStart, &rangeLength)
		if err != nil {
			log.Fatal(err)
		}

		currentMapList = append(currentMapList, Map{
			destinationRangeStart: destinationRangeStart,
			sourceRangeStart:      sourceRangeStart,
			rangeLength:           rangeLength,
		})
	}

	if len(currentMapList) > 0 {
		mappings = append(mappings, currentMapList)
	}

	lowestPosition := math.MaxInt32
	for _, seed := range seeds {
		currentPosition := seed

		for _, mapping := range mappings {
			for _, mapItem := range mapping {
				if currentPosition >= mapItem.sourceRangeStart && currentPosition < mapItem.sourceRangeStart+mapItem.rangeLength {
					currentPosition = mapItem.destinationRangeStart + (currentPosition - mapItem.sourceRangeStart)
					break
				}
			}
		}

		if currentPosition < lowestPosition {
			lowestPosition = currentPosition
		}
	}

	fmt.Println("Lowest position:", lowestPosition)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	seedsSplice := strings.Split(strings.TrimPrefix(scanner.Text(), "seeds: "), " ")
	seeds := make([]int, len(seedsSplice))
	for i, seed := range seedsSplice {
		seeds[i], err = strconv.Atoi(seed)
		if err != nil {
			log.Fatal(err)
		}
	}

	scanner.Scan()
	scanner.Scan()

	var mappings []MapList
	var currentMapList MapList

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			mappings = append(mappings, currentMapList)
			currentMapList = make(MapList, 0)

			scanner.Scan()
			continue
		}

		line := scanner.Text()

		var destinationRangeStart int
		var sourceRangeStart int
		var rangeLength int
		_, err = fmt.Sscanf(line, "%d %d %d", &destinationRangeStart, &sourceRangeStart, &rangeLength)
		if err != nil {
			log.Fatal(err)
		}

		currentMapList = append(currentMapList, Map{
			destinationRangeStart: destinationRangeStart,
			sourceRangeStart:      sourceRangeStart,
			rangeLength:           rangeLength,
		})
	}

	if len(currentMapList) > 0 {
		mappings = append(mappings, currentMapList)
	}

	lowestPosition := math.MaxInt32
	for pos := 1; pos < 1000000000; pos++ {
		currentPosition := pos

		for i := len(mappings) - 1; i >= 0; i-- {
			for _, mapItem := range mappings[i] {
				if currentPosition >= mapItem.destinationRangeStart && currentPosition < mapItem.destinationRangeStart+mapItem.rangeLength {
					currentPosition = mapItem.sourceRangeStart + (currentPosition - mapItem.destinationRangeStart)
					break
				}
			}
		}

		foundResult := false
		for i := 0; i < len(seeds); i += 2 {
			seedFrom, seedsLen := seeds[i], seeds[i+1]

			if currentPosition >= seedFrom && currentPosition < seedFrom+seedsLen {
				lowestPosition = pos
				foundResult = true
				break
			}
		}

		if foundResult {
			break
		}
	}

	fmt.Println("Lowest position:", lowestPosition)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
