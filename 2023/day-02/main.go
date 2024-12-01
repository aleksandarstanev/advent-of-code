package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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

	lineNum := 1
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ";")

		works := true
		for _, part := range parts {
			// Regular expression to match "number color" pattern
			re := regexp.MustCompile(`(\d+) (\w+)`)

			matches := re.FindAllStringSubmatch(part, -1)

			blue, red, green := 0, 0, 0
			for _, match := range matches {
				number := match[1]
				color := match[2]

				val, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}

				if color == "blue" {
					blue += val
				} else if color == "red" {
					red += val
				} else if color == "green" {
					green += val
				}
			}

			if red > 12 || green > 13 || blue > 14 {
				works = false
				break
			}
		}

		if works {
			ans += lineNum
		}

		lineNum += 1
	}

	fmt.Println(ans)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNum := 1
	ans := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ";")

		maxBlue, maxRed, maxGreen := 0, 0, 0
		for _, part := range parts {
			// Regular expression to match "number color" pattern
			re := regexp.MustCompile(`(\d+) (\w+)`)

			matches := re.FindAllStringSubmatch(part, -1)

			blue, red, green := 0, 0, 0
			for _, match := range matches {
				number := match[1]
				color := match[2]

				val, err := strconv.Atoi(number)
				if err != nil {
					log.Fatal(err)
				}

				if color == "blue" {
					blue += val
				} else if color == "red" {
					red += val
				} else if color == "green" {
					green += val
				}
			}

			maxBlue = max(maxBlue, blue)
			maxRed = max(maxRed, red)
			maxGreen = max(maxGreen, green)
		}

		ans += maxBlue * maxRed * maxGreen

		lineNum += 1
	}

	fmt.Println(ans)
}
