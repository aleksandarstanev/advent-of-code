package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start string
	end   string
}

type Input struct {
	ranges []Range
}

func getDigits(num int64) int {
	digits := 0
	for num > 0 {
		num /= 10
		digits++
	}

	return digits
}

func getInvalidIdsSum(start, end string) int64 {
	startNum, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		panic(err)
	}

	endNum, err := strconv.ParseInt(end, 10, 64)
	if err != nil {
		panic(err)
	}

	lowerRange, err := strconv.Atoi(start[:len(start)/2])
	if err != nil {
		lowerRange = 0
	}

	upperRange, err := strconv.Atoi(end[:(len(end)/2 + len(end)%2)])
	if err != nil {
		log.Fatal(err)
	}

	var res int64
	for i := lowerRange; i <= upperRange; i++ {
		digits := getDigits(int64(i))

		current := int64(i)
		for j := 0; j < digits; j++ {
			current *= 10
		}

		current += int64(i)

		if current >= startNum && current <= endNum {
			res += current
		}
	}

	return res
}

func part1(input Input) {
	var ans int64
	for _, rng := range input.ranges {
		ans += getInvalidIdsSum(rng.start, rng.end)
	}

	fmt.Println("Part 1:", ans)
}

func isInvalidId(id int64) bool {
	digits := getDigits(id)

	for i := 1; i < digits; i++ {
		if digits%i != 0 {
			continue
		}

		str := strconv.FormatInt(id, 10)

		works := true
		for j := 0; j < digits; j += i {
			for k := 0; k < i; k++ {
				if str[k] != str[j+k] {
					works = false
				}
			}
		}

		if works {
			return true
		}
	}

	return false
}

func part2(input Input) {
	var ans int64
	for _, rng := range input.ranges {
		startNum, err := strconv.ParseInt(rng.start, 10, 64)
		if err != nil {
			panic(err)
		}

		endNum, err := strconv.ParseInt(rng.end, 10, 64)
		if err != nil {
			panic(err)
		}

		for i := startNum; i <= endNum; i++ {
			if isInvalidId(i) {
				ans += i
			}
		}
	}

	fmt.Println("Part 2:", ans)
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var input Input
	ranges := strings.Split(string(file), ",")
	for _, rng := range ranges {
		parts := strings.Split(rng, "-")
		start, end := parts[0], parts[1]

		input.ranges = append(input.ranges, Range{
			start: start,
			end:   end,
		})
	}

	// part1(input)
	part2(input)
}
