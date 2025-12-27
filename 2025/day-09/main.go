package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"sort"
)

type Tile struct {
	r, c int
}

func abs(n int) int {
	if n >= 0 {
		return n
	}

	return -n
}

func part1(tiles []Tile) {
	ans := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			dist := (abs(tiles[i].r-tiles[j].r) + 1) * (abs(tiles[i].c-tiles[j].c) + 1)

			ans = max(ans, dist)
		}
	}

	fmt.Println("Part 1:", ans)
}

func probe(simplifiedGrid [][]bool, r, c int, visited [][]bool) bool {
	rows := len(simplifiedGrid)
	cols := len(simplifiedGrid[0])

	if r < 0 || c < 0 || r >= rows || c >= cols {
		return false
	}

	if visited[r][c] {
		return true
	}

	if simplifiedGrid[r][c] {
		return true
	}

	visited[r][c] = true

	return probe(simplifiedGrid, r+1, c, visited) && probe(simplifiedGrid, r-1, c, visited) && probe(simplifiedGrid, r, c+1, visited) && probe(simplifiedGrid, r, c-1, visited)
}

func fill(simplifiedGrid [][]bool, r, c int) {
	rows := len(simplifiedGrid)
	cols := len(simplifiedGrid[0])

	if r < 0 || c < 0 || r >= rows || c >= cols || simplifiedGrid[r][c] {
		return
	}

	simplifiedGrid[r][c] = true

	fill(simplifiedGrid, r+1, c)
	fill(simplifiedGrid, r-1, c)
	fill(simplifiedGrid, r, c+1)
	fill(simplifiedGrid, r, c-1)
}

func part2(tiles []Tile) {
	seenRows := make(map[int]bool, 0)
	seenCols := make(map[int]bool, 0)

	for _, tile := range tiles {
		seenRows[tile.r] = true
		seenCols[tile.c] = true
	}

	rows := make([]int, 0)
	for k := range seenRows {
		rows = append(rows, k)
	}

	cols := make([]int, 0)
	for k := range seenCols {
		cols = append(cols, k)
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i] < rows[j]
	})

	sort.Slice(cols, func(i, j int) bool {
		return cols[i] < cols[j]
	})

	compressedRows := make(map[int]int, 0)
	compressedCols := make(map[int]int, 0)

	for i, row := range rows {
		compressedRows[row] = i
	}

	for i, col := range cols {
		compressedCols[col] = i
	}

	simplifiedGrid := make([][]bool, len(rows))
	for i := 0; i < len(rows); i++ {
		simplifiedGrid[i] = make([]bool, len(cols))
	}

	simplifiedGrid[compressedRows[tiles[0].r]][compressedCols[tiles[0].c]] = true

	for i := 1; i <= len(tiles); i++ {
		idx := i % len(tiles)
		simplifiedGrid[compressedRows[tiles[idx].r]][compressedCols[tiles[idx].c]] = true
		if tiles[idx].r == tiles[i-1].r {
			from := min(compressedCols[tiles[idx].c], compressedCols[tiles[i-1].c])
			to := max(compressedCols[tiles[idx].c], compressedCols[tiles[i-1].c])

			for j := from; j <= to; j++ {
				simplifiedGrid[compressedRows[tiles[idx].r]][j] = true
			}
		} else {
			from := min(compressedRows[tiles[idx].r], compressedRows[tiles[i-1].r])
			to := max(compressedRows[tiles[idx].r], compressedRows[tiles[i-1].r])

			for j := from; j <= to; j++ {
				simplifiedGrid[j][compressedCols[tiles[idx].c]] = true
			}
		}
	}

	tryR, tryC := rand.IntN(len(simplifiedGrid)), rand.IntN(len(simplifiedGrid[0]))
	visited := make([][]bool, len(simplifiedGrid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(simplifiedGrid[0]))
	}

	for {
		if simplifiedGrid[tryR][tryC] == true {
			tryR, tryC = rand.IntN(len(simplifiedGrid)), rand.IntN(len(simplifiedGrid[0]))
			continue
		}

		works := probe(simplifiedGrid, tryR, tryC, visited)
		if works {
			break
		}

		tryR, tryC = rand.IntN(len(simplifiedGrid)), rand.IntN(len(simplifiedGrid[0]))
		for i := 0; i < len(visited); i++ {
			for j := 0; j < len(visited[i]); j++ {
				visited[i][j] = false
			}
		}
	}

	fill(simplifiedGrid, tryR, tryC)

	ans := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			fromR := min(compressedRows[tiles[i].r], compressedRows[tiles[j].r])
			toR := max(compressedRows[tiles[i].r], compressedRows[tiles[j].r])

			fromC := min(compressedCols[tiles[i].c], compressedCols[tiles[j].c])
			toC := max(compressedCols[tiles[i].c], compressedCols[tiles[j].c])

			works := true
			for r := fromR; r <= toR; r++ {
				for c := fromC; c <= toC; c++ {
					if simplifiedGrid[r][c] == false {
						works = false
					}
				}
			}

			if works {
				dist := (abs(tiles[i].r-tiles[j].r) + 1) * (abs(tiles[i].c-tiles[j].c) + 1)
				ans = max(ans, dist)
			}
		}
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

	var tiles []Tile
	for scanner.Scan() {
		line := scanner.Text()

		var c, r int
		_, err := fmt.Sscanf(line, "%d,%d", &c, &r)
		if err != nil {
			log.Fatal(err)
		}

		tiles = append(tiles, Tile{
			r: r,
			c: c,
		})
	}

	// part1(tiles)
	part2(tiles)
}
