package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getReflectionColumn(grid []string, skipCol int) int {
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < cols-1; i++ {
		if i == skipCol {
			continue
		}

		isReflective := true
		l, r := i, i+1
		for l >= 0 && r < cols && isReflective {
			for j := 0; j < rows; j++ {
				if grid[j][l] != grid[j][r] {
					isReflective = false
					break
				}
			}

			l -= 1
			r += 1
		}

		if isReflective {
			return i
		}
	}

	return -1
}

func getReflectionRow(grid []string, skipRow int) int {
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < rows-1; i++ {
		if i == skipRow {
			continue
		}

		isReflective := true
		l, r := i, i+1
		for l >= 0 && r < rows && isReflective {
			for j := 0; j < cols; j++ {
				if grid[l][j] != grid[r][j] {
					isReflective = false
					break
				}
			}

			l -= 1
			r += 1
		}

		if isReflective {
			return i
		}
	}

	return -1
}

func solvePart1(grid []string) int {
	return (getReflectionColumn(grid, -1) + 1) + (getReflectionRow(grid, -1)+1)*100
}

func solvePart2(grid []string) int {
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			reflectionCol := getReflectionColumn(grid, -1)
			reflectionRow := getReflectionRow(grid, -1)

			updatedGrid := make([]string, rows)
			copy(updatedGrid, grid)

			if grid[i][j] == '#' {
				updatedGrid[i] = updatedGrid[i][:j] + "." + updatedGrid[i][j+1:]
			} else {
				updatedGrid[i] = updatedGrid[i][:j] + "#" + updatedGrid[i][j+1:]
			}

			res := (getReflectionColumn(updatedGrid, reflectionCol) + 1) + (getReflectionRow(updatedGrid, reflectionRow)+1)*100
			if res > 0 {
				return res
			}
		}
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ans := 0
	grid := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			ans += solvePart2(grid)
			grid = make([]string, 0)
			continue
		}

		grid = append(grid, line)
	}

	ans += solvePart2(grid)

	fmt.Println("Result:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
