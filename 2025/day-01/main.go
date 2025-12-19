package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Instruction struct {
	direction rune
	rotation  int
}

type Input struct {
	instructions []Instruction
}

func part1(input Input) {
	position := 50
	zeroes := 0
	for _, instruction := range input.instructions {
		if instruction.direction == 'L' {
			position = (position + 100 - instruction.rotation) % 100
		} else {
			position = (position + instruction.rotation) % 100
		}

		if position == 0 {
			zeroes++
		}
	}

	fmt.Println("Part 1:", zeroes)
}

func part2(input Input) {
	position := 50
	ans := 0
	for _, instruction := range input.instructions {
		ans += instruction.rotation / 100
		rotation := instruction.rotation % 100
		if instruction.direction == 'L' {
			if position-rotation <= 0 && position != 0 {
				ans++
			}
			position = (position + 100 - rotation) % 100

		} else {
			if position+rotation >= 100 && position != 0 {
				ans++
			}
			position = (position + rotation) % 100
		}
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

		direction := line[0]
		rotation, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}

		input.instructions = append(input.instructions, Instruction{
			direction: rune(direction),
			rotation:  rotation,
		})
	}

	// part1(input)
	part2(input)
}
