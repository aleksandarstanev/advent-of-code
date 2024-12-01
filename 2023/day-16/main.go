package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	Up = iota
	Down
	Left
	Right
)

func isValid(grid []string, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[0])
}

func resolveNextPositions(grid []string, row, col, direction int) [][3]int {
	unfilteredNextPositions := make([][3]int, 0)

	switch direction {
	case Up:
		if grid[row][col] == '|' || grid[row][col] == '.' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row - 1, col, Up})
		} else if grid[row][col] == '/' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col + 1, Right})
		} else if grid[row][col] == '\\' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col - 1, Left})
		} else if grid[row][col] == '-' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col + 1, Right})
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col - 1, Left})
		}
	case Down:
		if grid[row][col] == '|' || grid[row][col] == '.' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row + 1, col, Down})
		} else if grid[row][col] == '/' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col - 1, Left})
		} else if grid[row][col] == '\\' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col + 1, Right})
		} else if grid[row][col] == '-' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col + 1, Right})
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col - 1, Left})
		}
	case Left:
		if grid[row][col] == '-' || grid[row][col] == '.' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col - 1, Left})
		} else if grid[row][col] == '/' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row + 1, col, Down})
		} else if grid[row][col] == '\\' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row - 1, col, Up})
		} else if grid[row][col] == '|' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row + 1, col, Down})
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row - 1, col, Up})
		}
	case Right:
		if grid[row][col] == '-' || grid[row][col] == '.' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row, col + 1, Right})
		} else if grid[row][col] == '/' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row - 1, col, Up})
		} else if grid[row][col] == '\\' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row + 1, col, Down})
		} else if grid[row][col] == '|' {
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row + 1, col, Down})
			unfilteredNextPositions = append(unfilteredNextPositions, [3]int{row - 1, col, Up})
		}
	}

	nextPositions := make([][3]int, 0)
	for _, nextPosition := range unfilteredNextPositions {
		if isValid(grid, nextPosition[0], nextPosition[1]) {
			nextPositions = append(nextPositions, nextPosition)
		}
	}

	return nextPositions
}

func getEnergizedCells(grid []string, startRow, startCol, startDir int) int {
	rows, cols := len(grid), len(grid[0])

	visited := make([][][4]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([][4]bool, cols)
	}

	queue := make([][3]int, 0)
	queue = append(queue, [3]int{startRow, startCol, startDir})
	visited[startRow][startCol][startDir] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		row, col, direction := current[0], current[1], current[2]
		nextPositions := resolveNextPositions(grid, row, col, direction)

		for _, nextPosition := range nextPositions {
			if !visited[nextPosition[0]][nextPosition[1]][nextPosition[2]] {
				visited[nextPosition[0]][nextPosition[1]][nextPosition[2]] = true
				queue = append(queue, nextPosition)
			}
		}
	}

	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for k := 0; k < 4; k++ {
				if visited[i][j][k] {
					res++
					break
				}
			}
		}
	}

	return res
}

func solvePart1(grid []string) int {
	return getEnergizedCells(grid, 0, 0, Right)
}

func solvePart2(grid []string) int {
	rows, cols := len(grid), len(grid[0])

	ans := 0
	for i := 0; i < rows; i++ {
		cur := getEnergizedCells(grid, i, 0, Right)
		if cur > ans {
			ans = cur
		}

		cur = getEnergizedCells(grid, i, cols-1, Left)
		if cur > ans {
			ans = cur
		}
	}

	for i := 0; i < cols; i++ {
		cur := getEnergizedCells(grid, 0, i, Down)
		if cur > ans {
			ans = cur
		}

		cur = getEnergizedCells(grid, rows-1, i, Up)
		if cur > ans {
			ans = cur
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
		grid = append(grid, scanner.Text())
	}

	ans := solvePart2(grid)
	fmt.Println("Answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
