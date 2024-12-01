package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

const STEPS_PART_2 = 26501365

func isValid(grid []string, row, col int) bool {
	rows, cols := len(grid), len(grid[0])

	if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] == '#' {
		return false
	}

	return true
}

func isValidPart2(grid []string, row, col int) bool {
	rows, cols := len(grid), len(grid[0])

	return grid[(rows+row%rows)%rows][(cols+col%cols)%cols] != '#'
}

func solvePart1(grid []string) int {
	const MAX_DISTANCE = 64

	rows, cols := len(grid), len(grid[0])
	dist := make([][]int, rows)

	startRow, startCol := 0, 0

	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'S' {
				startRow = i
				startCol = j
			}

			dist[i][j] = -1
		}
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{startRow, startCol})
	dist[startRow][startCol] = 0
	ans := 1

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		row, col := cur[0], cur[1]
		if dist[row][col] == MAX_DISTANCE {
			continue
		}

		if isValid(grid, row-1, col) && dist[row-1][col] == -1 {
			dist[row-1][col] = dist[row][col] + 1
			queue = append(queue, [2]int{row - 1, col})
			if dist[row-1][col]%2 == 0 {
				ans++
			}
		}

		if isValid(grid, row+1, col) && dist[row+1][col] == -1 {
			dist[row+1][col] = dist[row][col] + 1
			queue = append(queue, [2]int{row + 1, col})
			if dist[row+1][col]%2 == 0 {
				ans++
			}
		}

		if isValid(grid, row, col-1) && dist[row][col-1] == -1 {
			dist[row][col-1] = dist[row][col] + 1
			queue = append(queue, [2]int{row, col - 1})
			if dist[row][col-1]%2 == 0 {
				ans++
			}
		}

		if isValid(grid, row, col+1) && dist[row][col+1] == -1 {
			dist[row][col+1] = dist[row][col] + 1
			queue = append(queue, [2]int{row, col + 1})
			if dist[row][col+1]%2 == 0 {
				ans++
			}
		}

	}

	return ans
}

type Position struct {
	row int
	col int
}

func solvePart2Slow(grid []string, maxDistance int) int {
	rows, cols := len(grid), len(grid[0])
	dist := make(map[Position]int)

	startRow, startCol := 0, 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'S' {
				startRow = i
				startCol = j
			}
		}
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{startRow, startCol})
	dist[Position{startRow, startCol}] = 0
	ans := 0
	if maxDistance%2 == 0 {
		ans += 1
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		row, col := cur[0], cur[1]
		// fmt.Println(row, col)
		if dist[Position{row, col}] == maxDistance {
			continue
		}

		if isValidPart2(grid, row-1, col) {
			if _, ok := dist[Position{row - 1, col}]; !ok {
				dist[Position{row - 1, col}] = dist[Position{row, col}] + 1
				queue = append(queue, [2]int{row - 1, col})
				if dist[Position{row - 1, col}]%2 == maxDistance%2 {
					ans++
				}
			}
		}

		if isValidPart2(grid, row+1, col) {
			if _, ok := dist[Position{row + 1, col}]; !ok {
				dist[Position{row + 1, col}] = dist[Position{row, col}] + 1
				queue = append(queue, [2]int{row + 1, col})
				if dist[Position{row + 1, col}]%2 == maxDistance%2 {
					ans++
				}
			}
		}

		if isValidPart2(grid, row, col-1) {
			if _, ok := dist[Position{row, col - 1}]; !ok {
				dist[Position{row, col - 1}] = dist[Position{row, col}] + 1
				queue = append(queue, [2]int{row, col - 1})
				if dist[Position{row, col - 1}]%2 == maxDistance%2 {
					ans++
				}
			}
		}

		if isValidPart2(grid, row, col+1) {
			if _, ok := dist[Position{row, col + 1}]; !ok {
				dist[Position{row, col + 1}] = dist[Position{row, col}] + 1
				queue = append(queue, [2]int{row, col + 1})
				if dist[Position{row, col + 1}]%2 == maxDistance%2 {
					ans++
				}
			}
		}

	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func countEmptySquares(grid []string) int {
	const MAX_DISTANCE = 100

	rows, cols := len(grid), len(grid[0])

	startRow, startCol := 0, 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'S' {
				startRow = i
				startCol = j
			}
		}
	}

	res := 0
	for i := startRow - MAX_DISTANCE; i <= startRow+MAX_DISTANCE; i++ {
		for j := startCol - MAX_DISTANCE; j <= startCol+MAX_DISTANCE; j++ {
			dist := abs(i-startRow) + abs(j-startCol)
			if dist <= MAX_DISTANCE && dist%2 == MAX_DISTANCE%2 {
				if isValidPart2(grid, i, j) {
					res += 1
				}
			}
		}
	}

	return res
}

func enlargeGrid(grid []string) []string {
	rows, cols := len(grid), len(grid[0])

	newRows := rows * 3
	newCols := cols * 3

	newGrid := make([]string, newRows)

	for i := 0; i < newRows; i++ {
		newGrid[i] = ""
		for j := 0; j < newCols; j++ {
			newGrid[i] += string(grid[i%rows][j%cols])
		}
	}

	startRow, startCol := 0, 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'S' {
				startRow = i
				startCol = j
			}
		}
	}

	newGrid = newGrid[startRow+1 : newRows-startRow-1]
	for i := 0; i < len(newGrid); i++ {
		newGrid[i] = newGrid[i][startCol+1 : newCols-startCol-1]
	}

	return newGrid
}

func enlargeGridBy3(grid []string) []string {
	rows, cols := len(grid), len(grid[0])

	newRows := rows * 3
	newCols := cols * 3

	newGrid := make([]string, newRows)

	startRow, startCol := newRows/2, newCols/2
	for i := 0; i < newRows; i++ {
		newGrid[i] = ""
		for j := 0; j < newCols; j++ {
			if grid[i%rows][j%cols] == 'S' {
				if i == startRow && j == startCol {
					newGrid[i] += "S"
				} else {
					newGrid[i] += "."
				}
			} else {
				newGrid[i] += string(grid[i%rows][j%cols])
			}
		}
	}

	return newGrid
}

func solvePart2(grid []string, steps int) int {
	originalRows, originalCols := len(grid), len(grid[0])

	originalStartRow, originalStartCol := 0, 0
	for i := 0; i < originalRows; i++ {
		for j := 0; j < originalCols; j++ {
			if grid[i][j] == 'S' {
				originalStartRow = i
				originalStartCol = j
			}
		}
	}

	grid = enlargeGridBy3(grid)

	rows, cols := len(grid), len(grid[0])
	fmt.Println(rows, cols)

	// for _, line := range grid {
	// 	fmt.Println(line)
	// }

	startRow, startCol := 0, 0
	dist := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			if grid[i][j] == 'S' {
				startRow = i
				startCol = j
			}

			dist[i][j] = -1
		}
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{startRow, startCol})
	dist[startRow][startCol] = 0

	fmt.Println(startRow, startCol)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		row, col := cur[0], cur[1]

		if isValid(grid, row-1, col) && dist[row-1][col] == -1 {
			dist[row-1][col] = dist[row][col] + 1
			queue = append(queue, [2]int{row - 1, col})
		}

		if isValid(grid, row+1, col) && dist[row+1][col] == -1 {
			dist[row+1][col] = dist[row][col] + 1
			queue = append(queue, [2]int{row + 1, col})
		}

		if isValid(grid, row, col-1) && dist[row][col-1] == -1 {
			dist[row][col-1] = dist[row][col] + 1
			queue = append(queue, [2]int{row, col - 1})
		}

		if isValid(grid, row, col+1) && dist[row][col+1] == -1 {
			dist[row][col+1] = dist[row][col] + 1
			queue = append(queue, [2]int{row, col + 1})
		}

	}

	grid = grid[originalStartRow+1 : rows-originalStartRow-1]
	for i := 0; i < len(grid); i++ {
		grid[i] = grid[i][originalStartCol+1 : cols-originalStartCol-1]
	}

	rows, cols = len(grid), len(grid[0])
	distCopy := make([][]int, rows)
	newStartRow, newStartCol := 0, 0
	for i := 0; i < rows; i++ {
		distCopy[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			distCopy[i][j] = dist[i+originalStartRow+1][j+originalStartCol+1]
			if grid[i][j] == 'S' {
				newStartRow = i
				newStartCol = j
			}
		}
	}

	ans := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if distCopy[i][j] == -1 {
				continue
			}

			if i == newStartRow || j == newStartCol {
				continue
			}

			if distCopy[i][j]%2 == steps%2 {
				indicesSum := int(math.Floor(float64(steps-distCopy[i][j]) / float64(originalRows)))
				if indicesSum >= 0 {
					halved := indicesSum / 2
					toAdd1 := ((halved + 1) * (halved + 2) / 2)
					toAdd2 := (halved * (halved + 1) / 2)

					ans = ans + toAdd1 + toAdd2
				}
			} else {
				indicesSum := int(math.Floor(float64(steps-distCopy[i][j]) / float64(originalRows)))
				if indicesSum >= 0 {
					halved := (indicesSum - 1) / 2
					toAdd1 := ((halved + 1) * (halved + 2) / 2)
					toAdd2 := ((halved + 1) * (halved + 2) / 2)

					ans = ans + toAdd1 + toAdd2
				}
			}

			// for ii := 0; ii <= originalRows; ii++ {
			// 	for jj := 0; jj <= originalCols; jj++ {
			// 		curDist := distCopy[i][j] + ii*originalRows + jj*originalCols

			// 		if curDist <= steps && curDist%2 == steps%2 {
			// 			// if curDist == steps {
			// 			// 	r, c := i+ii*originalRows, j+jj*originalCols
			// 			// 	fmt.Println(r, c, string(grid[r%originalRows][c%originalCols]), i, j, dist[i][j])
			// 			// }
			// 			ans++
			// 		}
			// 	}
			// }
		}
	}

	// for i := -1000; i <= 1000; i++ {
	// 	for j := -1000; j <= 1000; j++ {
	// 		if i%originalRows != 0 && j%originalCols != 0 {
	// 			continue
	// 		}

	// 		d := abs(i) + abs(j)
	// 		if d <= steps && d%2 == steps%2 {
	// 			ans++
	// 		}

	// 	}
	// }

	tempSum := 0
	for i := -steps; i <= steps; i++ {
		absI := abs(i)

		// var delta int
		// var start int
		if i%originalRows == 0 {
			if absI%2 == steps%2 {
				tempSum += (1 + (steps - absI))
			} else {
				tempSum += (steps - absI + 1)
			}

			continue
		} else {
			if absI%2 == steps%2 {
				tempSum += (1 + (((steps-absI)/originalRows)/2)*2)
			} else {
				tempSum += (((steps - absI + originalRows) / originalRows / 2) * 2)
			}

			// delta = originalCols
			// start = ((-steps+absI)/originalCols)*originalCols - originalCols
		}

		// for j := start; j <= steps-absI; j += delta {
		// 	if i%originalRows != 0 && j%originalCols != 0 {
		// 		continue
		// 	}

		// 	d := abs(i) + abs(j)
		// 	if d <= steps && d%2 == steps%2 {
		// 		tempSum++
		// 	}
		// }
	}

	// fmt.Println(tempSum)

	ans = ans + tempSum

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

	steps := STEPS_PART_2
	// ans := solvePart2Slow(grid, steps)
	// fmt.Println("Slow solution answer:", ans)

	ans := solvePart2(grid, steps)
	fmt.Println("Fast solution answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// 528192865877841 --> too low
// 621944727930768
