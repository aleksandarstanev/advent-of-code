package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	nums       [][]int
	operations []byte
}

type InputRaw struct {
	numberLines []string
	numberEnds  []int
	operations  []byte
}

func part1(input Input) {
	rows := len(input.nums)
	cols := len(input.nums[0])
	ans := 0

	for c := 0; c < cols; c++ {
		var res int
		if input.operations[c] == '+' {
			res = 0
		} else {
			res = 1
		}

		for r := 0; r < rows; r++ {
			if input.operations[c] == '+' {
				res += input.nums[r][c]
			} else {
				res *= input.nums[r][c]
			}
		}

		ans += res
	}

	fmt.Println("Part 1:", ans)
}

func part2(input InputRaw) {
	rows := len(input.numberLines)
	ans := 0

	pos := 0
	for i := 0; i < len(input.numberEnds); i++ {
		var res int
		if input.operations[i] == '+' {
			res = 0
		} else {
			res = 1
		}
		for pos < input.numberEnds[i] {
			cur := 0
			for r := 0; r < rows; r++ {
				c := input.numberLines[r][pos]
				if c >= '0' && c <= '9' {
					cur = cur*10 + (int(c) - '0')
				}
			}

			if input.operations[i] == '+' {
				res += cur
			} else {
				res *= cur
			}

			pos++
		}

		ans += res

		pos++
	}

	fmt.Println("Part 2", ans)
}

func parseInput(file *os.File) Input {
	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	var input Input
	input.nums = make([][]int, len(lines)-1)
	for i := 0; i < len(lines)-1; i++ {
		parts := strings.Split(lines[i], " ")
		input.nums[i] = make([]int, 0)
		for _, part := range parts {
			trimmed := strings.TrimSpace(part)
			if trimmed == "" {
				continue
			}

			num, err := strconv.Atoi(trimmed)
			if err != nil {
				log.Fatal(err)
			}

			input.nums[i] = append(input.nums[i], num)
		}
	}

	parts := strings.Split(lines[len(lines)-1], " ")
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}

		if trimmed[0] == '+' {
			input.operations = append(input.operations, '+')
		} else {
			input.operations = append(input.operations, '*')
		}
	}

	return input
}

func parseInputRaw(file *os.File) InputRaw {
	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	var input InputRaw
	input.numberLines = make([]string, len(lines)-1)
	for i := 0; i < len(lines)-1; i++ {
		input.numberLines[i] = lines[i]
	}

	cols := len(lines[0])
	for c := 0; c < cols; c++ {
		allSpaces := true
		for r := 0; r < len(lines)-1; r++ {
			if lines[r][c] != ' ' {
				allSpaces = false
				break
			}
		}

		if allSpaces {
			input.numberEnds = append(input.numberEnds, c)
		}
	}

	input.numberEnds = append(input.numberEnds, cols)

	parts := strings.Split(lines[len(lines)-1], " ")
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}

		if trimmed[0] == '+' {
			input.operations = append(input.operations, '+')
		} else {
			input.operations = append(input.operations, '*')
		}
	}

	return input
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// input := parseInput(file)
	// part1(input)

	input := parseInputRaw(file)
	part2(input)
}
