package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Input struct {
	banks []string
}

func part1(input Input) {
	ans := 0
	for _, bank := range input.banks {
		maxJolt := 0
		maxDigit := int(bank[0]) - '0'
		for i := 1; i < len(bank); i++ {
			maxJolt = max(maxJolt, maxDigit*10+int(bank[i])-'0')
			maxDigit = max(maxDigit, int(bank[i])-'0')
		}

		ans += maxJolt
	}

	fmt.Println("Part 1:", ans)
}

func part2(input Input) {
	var ans int64
	for _, bank := range input.banks {
		maxNums := make([]int64, 12)
		for i := 0; i < 12; i++ {
			maxNums[i] = -1
		}

		maxNums[0] = int64(bank[0]) - '0'
		for i := 1; i < len(bank); i++ {
			for j := 10; j >= 0; j-- {
				if maxNums[j] != -1 {
					maxNums[j+1] = max(maxNums[j+1], maxNums[j]*10+int64(bank[i])-'0')
				}
			}

			maxNums[0] = max(maxNums[0], int64(bank[i])-'0')
		}

		ans += maxNums[11]
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

		input.banks = append(input.banks, line)
	}

	// part1(input)
	part2(input)
}
