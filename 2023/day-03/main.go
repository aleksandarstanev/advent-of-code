package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getAllAdjacentGears(curRow int, curCols []int, totalRows, totalCols int, lines []string) [][]int {
	adjacentGears := [][]int{}
	for _, curCol := range curCols {
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

				if lines[adjRow][adjCol] == '*' {
					adjacentGears = append(adjacentGears, []int{adjRow, adjCol})
				}
			}
		}
	}

	return adjacentGears
}

func main() {
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
	gearToCount := make(map[int]int)
	gearToProduct := make(map[int]int)
	for row, line := range lines {
		for col, c := range line {
			if c >= '0' && c <= '9' {
				curNumber = curNumber*10 + int(c-'0')
				curCols = append(curCols, col)
			} else {
				if curNumber > 0 {
					adjacentGears := getAllAdjacentGears(row, curCols, totalRows, totalCols, lines)

					used := make(map[int]bool)
					for _, gear := range adjacentGears {
						gearHash := gear[0]*totalCols + gear[1]
						if used[gearHash] {
							continue
						}

						gearToCount[gearHash]++

						if gearToCount[gearHash] == 1 {
							gearToProduct[gearHash] = curNumber
						} else {
							gearToProduct[gearHash] *= curNumber
						}

						used[gearHash] = true
					}
				}
				curNumber = 0
				curCols = []int{}
			}
		}

		if curNumber > 0 {
			if curNumber > 0 {
				adjacentGears := getAllAdjacentGears(row, curCols, totalRows, totalCols, lines)

				for _, gear := range adjacentGears {
					gearHash := gear[0]*totalCols + gear[1]
					gearToCount[gearHash]++

					if gearToCount[gearHash] == 1 {
						gearToProduct[gearHash] = curNumber
					} else {
						gearToProduct[gearHash] *= curNumber
					}
				}
			}
			curNumber = 0
			curCols = []int{}
		}
	}

	for hash, product := range gearToProduct {
		if gearToCount[hash] == 2 {
			ans += product
		}
	}

	fmt.Println(ans)
}
