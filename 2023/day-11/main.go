package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func calculateDistances(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])

	rowContainsGalaxy := make([]bool, rows)
	colContainsGalaxy := make([]bool, cols)

	galaxies := make([][2]int, 0)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '#' {
				rowContainsGalaxy[i] = true
				colContainsGalaxy[j] = true

				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}

	ans := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			minRow := min(galaxies[i][0], galaxies[j][0])
			maxRow := max(galaxies[i][0], galaxies[j][0])

			minCol := min(galaxies[i][1], galaxies[j][1])
			maxCol := max(galaxies[i][1], galaxies[j][1])

			dist := maxRow - minRow + maxCol - minCol

			for k := minRow + 1; k < maxRow; k++ {
				if !rowContainsGalaxy[k] {
					dist += 999999
				}
			}

			for k := minCol + 1; k < maxCol; k++ {
				if !colContainsGalaxy[k] {
					dist += 999999
				}
			}

			ans += dist
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

	grid := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		grid = append(grid, line)
	}

	fmt.Println(calculateDistances(grid))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
