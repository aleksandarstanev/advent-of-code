package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(grid []string) {
	rows := len(grid)
	cols := len(grid[0])
	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cnt := 0
			for ii := -1; ii <= 1; ii++ {
				for jj := -1; jj <= 1; jj++ {
					if ii == 0 && jj == 0 {
						continue
					}

					newI := i + ii
					newJ := j + jj

					if newI >= 0 && newI < rows && newJ >= 0 && newJ < cols && grid[newI][newJ] == '@' {
						cnt++
					}
				}
			}

			if grid[i][j] == '@' && cnt < 4 {
				ans++
			}
		}
	}

	fmt.Println("Part 1:", ans)
}

func removePaper(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	removed := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			cnt := 0
			for ii := -1; ii <= 1; ii++ {
				for jj := -1; jj <= 1; jj++ {
					if ii == 0 && jj == 0 {
						continue
					}

					newI := i + ii
					newJ := j + jj

					if newI >= 0 && newI < rows && newJ >= 0 && newJ < cols && grid[newI][newJ] == '@' {
						cnt++
					}
				}
			}

			if grid[i][j] == '@' && cnt < 4 {
				grid[i][j] = '.'
				removed++
			}
		}
	}

	return removed
}

func part2(grid []string) {
	rows := len(grid)

	gridToBytes := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		gridToBytes[i] = []byte(grid[i])
	}

	removed := removePaper(gridToBytes)
	finalRemoved := removed
	for removed > 0 {
		removed = removePaper(gridToBytes)
		finalRemoved += removed
	}

	fmt.Println("Part 1:", finalRemoved)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid []string
	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, line)
	}

	// part1(grid)
	part2(grid)
}
