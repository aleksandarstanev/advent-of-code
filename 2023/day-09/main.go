package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func calculateNextNumber(numbers []int) int {
	numSlices := make([][]int, 0)
	numSlices = append(numSlices, numbers)
	lastSlice := numbers

	areAllZeros := false
	for !areAllZeros {
		areAllZeros = true
		newSlice := make([]int, len(lastSlice)-1)

		for i := 0; i < len(lastSlice)-1; i++ {
			newSlice[i] = lastSlice[i+1] - lastSlice[i]

			if newSlice[i] != 0 {
				areAllZeros = false
			}
		}

		lastSlice = newSlice
		numSlices = append(numSlices, newSlice)
	}

	ans := 0
	for i := len(numSlices) - 2; i >= 0; i-- {
		ans = numSlices[i][len(numSlices[i])-1] + ans
	}

	return ans
}

func calculateNextNumberPart2(numbers []int) int {
	numSlices := make([][]int, 0)
	numSlices = append(numSlices, numbers)
	lastSlice := numbers

	areAllZeros := false
	for !areAllZeros {
		areAllZeros = true
		newSlice := make([]int, len(lastSlice)-1)

		for i := 0; i < len(lastSlice)-1; i++ {
			newSlice[i] = lastSlice[i+1] - lastSlice[i]

			if newSlice[i] != 0 {
				areAllZeros = false
			}
		}

		lastSlice = newSlice
		numSlices = append(numSlices, newSlice)
	}

	ans := 0
	for i := len(numSlices) - 2; i >= 0; i-- {
		ans = numSlices[i][0] - ans
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

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()

		values := strings.Split(line, " ")

		numbers := make([]int, 0)
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}

			numbers = append(numbers, num)
		}

		nextNum := calculateNextNumberPart2(numbers)

		// fmt.Println(nextNum)
		ans += nextNum
	}

	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
