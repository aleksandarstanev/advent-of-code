package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int64
	end   int64
}

type Input struct {
	ranges      []Range
	ingredients []int64
}

func part1(input Input) {
	ans := 0
	for _, ingredient := range input.ingredients {
		for _, rng := range input.ranges {
			if ingredient >= rng.start && ingredient <= rng.end {
				ans++
				break
			}
		}
	}

	fmt.Println("Part 1:", ans)
}

func part2(input Input) {
	var ans int64

	sort.Slice(input.ranges, func(i, j int) bool {
		return input.ranges[i].start < input.ranges[j].start
	})

	var currentMax int64
	currentMax = math.MinInt64
	for _, rng := range input.ranges {
		currentMax = max(currentMax, rng.start)
		if currentMax <= rng.end {
			ans += (rng.end - currentMax + 1)
		}

		currentMax = max(currentMax, rng.end+1)
	}

	fmt.Println("Part 2:", ans)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input Input
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, "-")
		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		input.ranges = append(input.ranges, Range{
			start: start,
			end:   end,
		})
	}

	for scanner.Scan() {
		line := scanner.Text()

		ingredient, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		input.ingredients = append(input.ingredients, ingredient)
	}

	// part1(input)
	part2(input)
}
