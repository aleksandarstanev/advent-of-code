package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getResult(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	ans := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'O' {
				ans += (rows - r)
			}
		}
	}

	return ans
}

const (
	North = iota
	West
	South
	East
)

func tilt(grid [][]rune, direction int) {
	rows, cols := len(grid), len(grid[0])

	switch direction {
	case North:
		for c := 0; c < cols; c++ {
			for r := 0; r < rows; r++ {
				if grid[r][c] != 'O' {
					continue
				}

				rr := r - 1
				for rr >= 0 && grid[rr][c] == '.' {
					rr--
				}

				rr++
				if rr != r {
					tmp := grid[r][c]
					grid[r][c] = grid[rr][c]
					grid[rr][c] = tmp
				}
			}
		}
	case West:
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] != 'O' {
					continue
				}

				cc := c - 1
				for cc >= 0 && grid[r][cc] == '.' {
					cc--
				}

				cc++
				if cc != c {
					tmp := grid[r][c]
					grid[r][c] = grid[r][cc]
					grid[r][cc] = tmp
				}
			}
		}
	case South:
		for c := 0; c < cols; c++ {
			for r := rows - 1; r >= 0; r-- {
				if grid[r][c] != 'O' {
					continue
				}

				rr := r + 1
				for rr < rows && grid[rr][c] == '.' {
					rr++
				}

				rr--
				if rr != r {
					tmp := grid[r][c]
					grid[r][c] = grid[rr][c]
					grid[rr][c] = tmp
				}
			}
		}
	case East:
		for r := 0; r < rows; r++ {
			for c := cols - 1; c >= 0; c-- {
				if grid[r][c] != 'O' {
					continue
				}

				cc := c + 1
				for cc < cols && grid[r][cc] == '.' {
					cc++
				}

				cc--
				if cc != c {
					tmp := grid[r][c]
					grid[r][c] = grid[r][cc]
					grid[r][cc] = tmp
				}
			}
		}
	}
}

func solvePart1(grid [][]rune) int {
	tilt(grid, North)

	return getResult(grid)
}

func solvePart2(grid [][]rune) int {
	saved := make(map[string]int)
	i := 0
	for i = 0; i < 1000000000; i++ {
		for j := 0; j < 4; j++ {
			tilt(grid, j)
		}

		str := ""
		for r := 0; r < len(grid); r++ {
			str += string(grid[r])
		}

		if val, ok := saved[str]; ok {
			fmt.Println("Found", val, i)

			period := i - val
			i += ((1000000000-i)/period - 1) * period

			break
		}

		saved[str] = i
	}

	i += 1
	for i < 1000000000 {
		for j := 0; j < 4; j++ {
			tilt(grid, j)
		}

		i += 1
	}

	return getResult(grid)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ans := 0
	grid := make([][]rune, 0)
	for scanner.Scan() {
		line := []rune(scanner.Text())

		grid = append(grid, line)
	}

	ans += solvePart2(grid)

	fmt.Println("Result:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
