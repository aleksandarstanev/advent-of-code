package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func isValid(grid []string, row int, col int) bool {
	rows, cols := len(grid), len(grid[0])
	return row >= 0 && row < rows && col >= 0 && col < cols
}

func solvePart1(grid []string, startRow int, startCol int) int {
	rows, cols := len(grid), len(grid[0])
	dist := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)

		for j := 0; j < cols; j++ {
			dist[i][j] = -1
		}
	}

	queue := make([][2]int, 0)

	queue = append(queue, [2]int{startRow, startCol})
	dist[startRow][startCol] = 0

	ans := 0
	for len(queue) > 0 {
		top := queue[0]
		currentRow, currentCol := top[0], top[1]
		curDist := dist[currentRow][currentCol]
		ans = max(ans, curDist)

		queue = queue[1:]

		if grid[currentRow][currentCol] == '.' {
			continue
		}

		newPositions := make([][2]int, 0)

		if grid[currentRow][currentCol] == '|' {
			newPositions = append(newPositions, [2]int{currentRow - 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow + 1, currentCol})
		}

		if grid[currentRow][currentCol] == '-' {
			newPositions = append(newPositions, [2]int{currentRow, currentCol - 1})
			newPositions = append(newPositions, [2]int{currentRow, currentCol + 1})
		}

		if grid[currentRow][currentCol] == 'L' {
			newPositions = append(newPositions, [2]int{currentRow - 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow, currentCol + 1})
		}

		if grid[currentRow][currentCol] == 'J' {
			newPositions = append(newPositions, [2]int{currentRow - 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow, currentCol - 1})
		}

		if grid[currentRow][currentCol] == '7' {
			newPositions = append(newPositions, [2]int{currentRow + 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow, currentCol - 1})
		}

		if grid[currentRow][currentCol] == 'F' {
			newPositions = append(newPositions, [2]int{currentRow + 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow, currentCol + 1})
		}

		if grid[currentRow][currentCol] == 'S' {
			if isValid(grid, currentRow-1, currentCol) {
				c := grid[currentRow-1][currentCol]

				if c == '7' || c == 'F' || c == '|' {
					newPositions = append(newPositions, [2]int{currentRow - 1, currentCol})
				}
			}

			if isValid(grid, currentRow+1, currentCol) {
				c := grid[currentRow+1][currentCol]

				if c == 'L' || c == 'J' || c == '|' {
					newPositions = append(newPositions, [2]int{currentRow + 1, currentCol})
				}
			}

			if isValid(grid, currentRow, currentCol-1) {
				c := grid[currentRow][currentCol-1]

				if c == 'L' || c == 'F' || c == '-' {
					newPositions = append(newPositions, [2]int{currentRow, currentCol - 1})
				}
			}

			if isValid(grid, currentRow, currentCol+1) {
				c := grid[currentRow][currentCol+1]

				if c == '7' || c == 'J' || c == '-' {
					newPositions = append(newPositions, [2]int{currentRow, currentCol + 1})
				}
			}
		}

		for _, pos := range newPositions {
			newRow, newCol := pos[0], pos[1]

			if isValid(grid, newRow, newCol) && dist[newRow][newCol] == -1 {
				dist[newRow][newCol] = curDist + 1
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
	}

	return ans
}

func expand(grid []string) []string {
	rows, cols := len(grid), len(grid[0])
	newGrid := make([]string, 2*rows-1)

	for i := 0; i < 2*rows-1; i++ {
		row := make([]rune, 2*cols-1)

		if i%2 == 0 {
			for j := 0; j < 2*cols-1; j++ {
				if j%2 == 0 {
					row[j] = rune(grid[i/2][j/2])
				} else {
					row[j] = '.'
				}
			}
		} else {
			for j := 0; j < 2*cols-1; j++ {
				row[j] = '.'
			}
		}

		newGrid[i] = string(row)
	}

	for i := 0; i < 2*rows-1; i++ {
		for j := 0; j < 2*cols-1; j++ {
			if i%2 == 1 && i > 0 && i < 2*rows-2 &&
				strings.Contains("7F|", string(newGrid[i-1][j])) && strings.Contains("LJ|", string(newGrid[i+1][j])) {
				newGrid[i] = newGrid[i][:j] + string('|') + newGrid[i][j+1:]
			}

			if j%2 == 1 && j > 0 && j < 2*cols-2 &&
				strings.Contains("LF-", string(newGrid[i][j-1])) && strings.Contains("7J-", string(newGrid[i][j+1])) {
				newGrid[i] = newGrid[i][:j] + string('-') + newGrid[i][j+1:]
			}
		}
	}

	return newGrid
}

func solvePart2(grid []string, startRow int, startCol int) int {
	rows, cols := len(grid), len(grid[0])
	dist := make([][]int, rows)

	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)

		for j := 0; j < cols; j++ {
			dist[i][j] = -1
		}
	}

	queue := make([][2]int, 0)

	queue = append(queue, [2]int{startRow, startCol})
	dist[startRow][startCol] = 0

	for len(queue) > 0 {
		top := queue[0]
		currentRow, currentCol := top[0], top[1]
		curDist := dist[currentRow][currentCol]

		queue = queue[1:]

		if grid[currentRow][currentCol] == '.' {
			continue
		}

		newPositions := make([][2]int, 0)

		if grid[currentRow][currentCol] == '|' {
			newPositions = append(newPositions, [2]int{currentRow - 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow + 1, currentCol})
		}

		if grid[currentRow][currentCol] == '-' {
			newPositions = append(newPositions, [2]int{currentRow, currentCol - 1})
			newPositions = append(newPositions, [2]int{currentRow, currentCol + 1})
		}

		if grid[currentRow][currentCol] == 'L' {
			newPositions = append(newPositions, [2]int{currentRow - 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow, currentCol + 1})
		}

		if grid[currentRow][currentCol] == 'J' {
			newPositions = append(newPositions, [2]int{currentRow - 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow, currentCol - 1})
		}

		if grid[currentRow][currentCol] == '7' {
			newPositions = append(newPositions, [2]int{currentRow + 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow, currentCol - 1})
		}

		if grid[currentRow][currentCol] == 'F' {
			newPositions = append(newPositions, [2]int{currentRow + 1, currentCol})
			newPositions = append(newPositions, [2]int{currentRow, currentCol + 1})
		}

		if grid[currentRow][currentCol] == 'S' {
			isPossible := make(map[rune]bool)
			isPossible['7'] = true
			isPossible['F'] = true
			isPossible['|'] = true
			isPossible['L'] = true
			isPossible['J'] = true
			isPossible['-'] = true

			if isValid(grid, currentRow-1, currentCol) {
				c := grid[currentRow-1][currentCol]

				if c == '7' || c == 'F' || c == '|' {
					newPositions = append(newPositions, [2]int{currentRow - 1, currentCol})
					isPossible['7'] = false
					isPossible['F'] = false
					isPossible['-'] = false
				}
			}

			if isValid(grid, currentRow+1, currentCol) {
				c := grid[currentRow+1][currentCol]

				if c == 'L' || c == 'J' || c == '|' {
					newPositions = append(newPositions, [2]int{currentRow + 1, currentCol})
					isPossible['L'] = false
					isPossible['J'] = false
					isPossible['-'] = false
				}
			}

			if isValid(grid, currentRow, currentCol-1) {
				c := grid[currentRow][currentCol-1]

				if c == 'L' || c == 'F' || c == '-' {
					newPositions = append(newPositions, [2]int{currentRow, currentCol - 1})
					isPossible['L'] = false
					isPossible['F'] = false
					isPossible['|'] = false
				}
			}

			if isValid(grid, currentRow, currentCol+1) {
				c := grid[currentRow][currentCol+1]

				if c == '7' || c == 'J' || c == '-' {
					newPositions = append(newPositions, [2]int{currentRow, currentCol + 1})
					isPossible['7'] = false
					isPossible['J'] = false
					isPossible['|'] = false
				}
			}

			for k, v := range isPossible {
				if v {
					grid[currentRow] = grid[currentRow][:currentCol] + string(k) + grid[currentRow][currentCol+1:]
				}
			}
		}

		for _, pos := range newPositions {
			newRow, newCol := pos[0], pos[1]

			if isValid(grid, newRow, newCol) && dist[newRow][newCol] == -1 {
				dist[newRow][newCol] = curDist + 1
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dist[i][j] == -1 {
				grid[i] = grid[i][:j] + string('.') + grid[i][j+1:]
			}
		}
	}

	newGrid := expand(grid)

	// for i := 0; i < len(newGrid); i++ {
	// 	fmt.Println(newGrid[i])
	// }

	expandedRows := len(newGrid)
	expandedCols := len(newGrid[0])
	ans := 0

	visited := make([][]bool, expandedRows)
	for i := 0; i < expandedRows; i++ {
		visited[i] = make([]bool, expandedCols)
		for j := 0; j < expandedCols; j++ {
			visited[i][j] = false
		}
	}

	for i := 0; i < expandedRows; i++ {
		for j := 0; j < expandedCols; j++ {
			if newGrid[i][j] == '.' && !visited[i][j] {
				isEnclosed := true
				tilesCount := 0

				visited[i][j] = true

				queue := make([][2]int, 0)
				queue = append(queue, [2]int{i, j})
				for len(queue) > 0 {
					top := queue[0]
					currentRow, currentCol := top[0], top[1]

					queue = queue[1:]

					if currentRow%2 == 0 && currentCol%2 == 0 {
						tilesCount += 1
					}

					if currentRow == 0 || currentRow == expandedRows-1 || currentCol == 0 || currentCol == expandedCols-1 {
						isEnclosed = false
					}

					if isValid(newGrid, currentRow-1, currentCol) && !visited[currentRow-1][currentCol] && newGrid[currentRow-1][currentCol] == '.' {
						visited[currentRow-1][currentCol] = true
						queue = append(queue, [2]int{currentRow - 1, currentCol})
					}

					if isValid(newGrid, currentRow+1, currentCol) && !visited[currentRow+1][currentCol] && newGrid[currentRow+1][currentCol] == '.' {
						visited[currentRow+1][currentCol] = true
						queue = append(queue, [2]int{currentRow + 1, currentCol})
					}

					if isValid(newGrid, currentRow, currentCol-1) && !visited[currentRow][currentCol-1] && newGrid[currentRow][currentCol-1] == '.' {
						visited[currentRow][currentCol-1] = true
						queue = append(queue, [2]int{currentRow, currentCol - 1})
					}

					if isValid(newGrid, currentRow, currentCol+1) && !visited[currentRow][currentCol+1] && newGrid[currentRow][currentCol+1] == '.' {
						visited[currentRow][currentCol+1] = true
						queue = append(queue, [2]int{currentRow, currentCol + 1})
					}
				}

				if isEnclosed {
					ans += tilesCount
				}
			}
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
	startRow, startCol := 0, 0
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)

		for j, c := range line {
			if c == 'S' {
				startRow = i
				startCol = j
			}
		}

		i += 1
	}

	fmt.Println(solvePart2(grid, startRow, startCol))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
