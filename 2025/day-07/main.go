package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func part1(grid []string) {
	s := map[int]bool{}
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 'S' {
			s[i] = true
		}
	}

	ans := 0
	for i := 1; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if s[j] && grid[i][j] == '^' {
				ans++
				s[j] = false
				if j > 0 && grid[i][j-1] != '^' {
					s[j-1] = true
				}
				if j < len(grid[i])-1 && grid[i][j+1] != '^' {
					s[j+1] = true
				}
			}
		}
	}

	fmt.Println("Part 1:", ans)
}

func part2(grid []string) {
	s := map[int]int{}
	for i := 0; i < len(grid[0]); i++ {
		if grid[0][i] == 'S' {
			s[i] = 1
		}
	}

	ans := 0
	for i := 1; i < len(grid); i++ {
		newS := map[int]int{}
		for k, v := range s {
			newS[k] = v
		}

		for j := 0; j < len(grid[i]); j++ {
			if s[j] > 0 && grid[i][j] == '^' {
				newS[j] = 0
				if j > 0 && grid[i][j-1] != '^' {
					newS[j-1] += s[j]
				}
				if j < len(grid[i])-1 && grid[i][j+1] != '^' {
					newS[j+1] += s[j]
				}
			}
		}

		s = newS
	}

	for _, v := range s {
		ans += v
	}

	fmt.Println("Part 2:", ans)
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

	// part1(grid)
	part2(grid)
}
