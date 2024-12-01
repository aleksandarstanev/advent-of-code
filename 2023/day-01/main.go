package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// func part1() {
// 	file, err := os.Open("example.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	totalSum := 0
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		firstDigit, lastDigit := -1, -1
// 		for i := 0; i < len(line); i++ {
// 			if line[i] >= '0' && line[i] <= '9' {
// 				lastDigit = int(line[i] - '0')

// 				if firstDigit == -1 {
// 					firstDigit = int(line[i] - '0')
// 				}
// 			}
// 		}

// 		if firstDigit == lastDigit {
// 			totalSum += firstDigit
// 		}
// 	}

// 	fmt.Println("Total sum:", totalSum)

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalSum := 0
	for scanner.Scan() {
		line := scanner.Text()

		firstDigit, lastDigit := -1, -1

		digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				lastDigit = int(line[i] - '0')

				if firstDigit == -1 {
					firstDigit = int(line[i] - '0')
				}
			}

			for num, digit := range digits {
				digitLen := len(digit)

				if i >= digitLen-1 {
					tryDigit := line[i-digitLen+1 : i+1]

					if tryDigit == digit {
						lastDigit = num + 1

						if firstDigit == -1 {
							firstDigit = num + 1
						}
					}
				}
			}
		}

		res := firstDigit*10 + lastDigit
		totalSum += res

		fmt.Println(res)
	}

	fmt.Println("Total sum: ", totalSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
