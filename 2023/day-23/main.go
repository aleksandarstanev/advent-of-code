package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func dfsPart1(row, col int, grid []string, visited [][]bool, steps int, maxSteps *int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '#' {
		return
	}

	if visited[row][col] {
		return
	}

	visited[row][col] = true
	defer func() {
		visited[row][col] = false
	}()

	if row == len(grid)-1 {
		fmt.Println(steps)
		if steps > *maxSteps {
			*maxSteps = steps
		}

		return
	}

	switch grid[row][col] {
	case '.':
		dfsPart1(row+1, col, grid, visited, steps+1, maxSteps)
		dfsPart1(row-1, col, grid, visited, steps+1, maxSteps)
		dfsPart1(row, col+1, grid, visited, steps+1, maxSteps)
		dfsPart1(row, col-1, grid, visited, steps+1, maxSteps)
	case '^':
		dfsPart1(row-1, col, grid, visited, steps+1, maxSteps)
	case '>':
		dfsPart1(row, col+1, grid, visited, steps+1, maxSteps)
	case 'v':
		dfsPart1(row+1, col, grid, visited, steps+1, maxSteps)
	case '<':
		dfsPart1(row, col-1, grid, visited, steps+1, maxSteps)
	}
}

func dfsPart2(row, col int, grid []string, visited [][]bool, steps int, maxSteps *int) {
	if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) || grid[row][col] == '#' {
		return
	}

	if visited[row][col] {
		return
	}

	visited[row][col] = true
	defer func() {
		visited[row][col] = false
	}()

	if row == len(grid)-1 {
		if steps > *maxSteps {
			*maxSteps = steps
		}

		return
	}

	dfsPart2(row+1, col, grid, visited, steps+1, maxSteps)
	dfsPart2(row-1, col, grid, visited, steps+1, maxSteps)
	dfsPart2(row, col+1, grid, visited, steps+1, maxSteps)
	dfsPart2(row, col-1, grid, visited, steps+1, maxSteps)
}

func solvePart1(grid []string) int {
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < cols; i++ {
		if grid[0][i] == '.' {
			visited := make([][]bool, rows)
			for i := range visited {
				visited[i] = make([]bool, cols)
			}

			maxSteps := 0
			dfsPart1(0, i, grid, visited, 0, &maxSteps)

			return maxSteps
		}
	}

	return -1
}

func solvePart2(grid []string) int {
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < cols; i++ {
		if grid[0][i] == '.' {
			visited := make([][]bool, rows)
			for i := range visited {
				visited[i] = make([]bool, cols)
			}

			maxSteps := 0
			dfsPart2(0, i, grid, visited, 0, &maxSteps)

			return maxSteps
		}
	}

	return -1
}

func main() {
	file, err := os.Open("example.txt")
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

	ans := solvePart2(grid)
	fmt.Println("Answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
