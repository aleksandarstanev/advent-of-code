package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isSymbol(character byte) bool {
	return character != '.' && (character < '0' || character > '9')
}

func isAdjacentToASymbol(curRow, curCol, totalRows, totalCols int, lines []string) bool {
	for deltaI := -1; deltaI <= 1; deltaI++ {
		for deltaJ := -1; deltaJ <= 1; deltaJ++ {
			if deltaI == 0 && deltaJ == 0 {
				continue
			}

			adjRow := curRow + deltaI
			adjCol := curCol + deltaJ

			if adjRow < 0 || adjRow >= totalRows || adjCol < 0 || adjCol >= totalCols {
				continue
			}

			if isSymbol(lines[adjRow][adjCol]) {
				return true
			}
		}
	}

	return false
}

func isAnyDigitAdjacentToASymbol(curRow int, curCols []int, totalRows, totalCols int, lines []string) bool {
	for _, curCol := range curCols {
		if isAdjacentToASymbol(curRow, curCol, totalRows, totalCols, lines) {
			return true
		}
	}

	return false
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	ans := 0
	curNumber := 0
	curCols := []int{}
	totalRows := len(lines)
	totalCols := len(lines[0])
	for row, line := range lines {
		for col, c := range line {
			if c >= '0' && c <= '9' {
				curNumber = curNumber*10 + int(c-'0')
				curCols = append(curCols, col)
			} else {
				if curNumber > 0 && isAnyDigitAdjacentToASymbol(row, curCols, totalRows, totalCols, lines) {
					ans += curNumber
					// fmt.Println(curNumber)
				}
				curNumber = 0
				curCols = []int{}
			}
		}

		if curNumber > 0 && isAnyDigitAdjacentToASymbol(row, curCols, totalRows, totalCols, lines) {
			ans += curNumber

			curNumber = 0
			curCols = []int{}
		}
	}

	fmt.Println(ans)
}
