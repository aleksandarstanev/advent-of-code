package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Lens struct {
	label       string
	focalLength int
}

type Box struct {
	lenses []Lens
}

func findLabel(box Box, label string) int {
	for i, lens := range box.lenses {
		if lens.label == label {
			return i
		}
	}

	return -1
}

func calculateHash(str string) int {
	res := 0
	for _, c := range str {
		res = res + int(c)
		res *= 17
		res %= 256
	}

	return res
}

func solvePart1(input string) int {
	parts := strings.Split(input, ",")

	total := 0
	for _, part := range parts {
		hash := calculateHash(part)

		total += hash
	}

	return total
}

func solvePart2(input string) int {
	parts := strings.Split(input, ",")

	boxes := make([]Box, 256)
	for _, part := range parts {
		len := len(part)
		if part[len-1] == '-' {
			label := part[:(len - 1)]

			boxIndex := calculateHash(label)
			box := &boxes[boxIndex]

			labelIndex := findLabel(*box, label)
			if labelIndex != -1 {
				box.lenses = append(box.lenses[:labelIndex], box.lenses[labelIndex+1:]...)
			}
		} else {
			label := part[:(len - 2)]
			focalLength := int(part[len-1] - '0')

			boxIndex := calculateHash(label)
			box := &boxes[boxIndex]

			labelIndex := findLabel(*box, label)
			if labelIndex != -1 {
				box.lenses[labelIndex].focalLength = focalLength
			} else {
				box.lenses = append(box.lenses, Lens{label: label, focalLength: focalLength})
			}
		}
	}

	ans := 0
	for boxIdx, box := range boxes {
		for lensIdx, lens := range box.lenses {
			ans += ((boxIdx + 1) * (lensIdx + 1) * lens.focalLength)
		}
	}

	return ans
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	input := scanner.Text()

	fmt.Println("Result:", solvePart2(input))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
