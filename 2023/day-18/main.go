package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	Up = iota
	Down
	Left
	Right
)

type Instruction struct {
	direction int
	steps     int
	color     string
}

func solvePart1(instructions []Instruction) int {
	const OFFSET = 1000

	rows, cols := 5000, 5000

	grid := make([][]bool, rows)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]bool, cols)
	}

	row, col := OFFSET, OFFSET
	for _, instruction := range instructions {
		for i := 0; i < instruction.steps; i++ {
			grid[row][col] = true

			switch instruction.direction {
			case Up:
				row--
			case Down:
				row++
			case Left:
				col--
			case Right:
				col++
			}
		}

		grid[row][col] = true
	}

	visited := make([][]bool, rows)
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, cols)
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})

	visited[0][0] = true

	count := 1

	for len(queue) > 0 {
		row, col := queue[0][0], queue[0][1]
		queue = queue[1:]

		if row > 0 && !visited[row-1][col] && !grid[row-1][col] {
			visited[row-1][col] = true
			count += 1
			queue = append(queue, [2]int{row - 1, col})
		}
		if row < rows-1 && !visited[row+1][col] && !grid[row+1][col] {
			visited[row+1][col] = true
			count += 1
			queue = append(queue, [2]int{row + 1, col})
		}
		if col > 0 && !visited[row][col-1] && !grid[row][col-1] {
			visited[row][col-1] = true
			count += 1
			queue = append(queue, [2]int{row, col - 1})
		}
		if col < cols-1 && !visited[row][col+1] && !grid[row][col+1] {
			visited[row][col+1] = true
			count += 1
			queue = append(queue, [2]int{row, col + 1})
		}
	}

	return rows*cols - count
}

func solvePart2(oldInstructions []Instruction) int64 {
	instructions := make([]Instruction, len(oldInstructions))
	for i := 0; i < len(oldInstructions); i++ {
		color := oldInstructions[i].color
		var direction int
		switch color[len(color)-1] {
		case '0':
			direction = Right
		case '1':
			direction = Down
		case '2':
			direction = Left
		case '3':
			direction = Up
		}

		steps, err := strconv.ParseInt(color[1:len(color)-1], 16, 64)
		if err != nil {
			log.Fatal(err)
		}

		instructions[i] = Instruction{
			direction: direction,
			steps:     int(steps),
		}
	}

	vertices := make([][2]int, 0)
	row, col := 0, 0
	circumference := 0
	for _, instruction := range instructions {
		switch instruction.direction {
		case Up:
			row -= instruction.steps
		case Down:
			row += instruction.steps
		case Left:
			col -= instruction.steps
		case Right:
			col += instruction.steps
		}

		circumference += instruction.steps

		vertices = append(vertices, [2]int{row, col})
	}

	// fmt.Println(vertices)

	area := int64(0)
	for i := 0; i < len(vertices); i++ {
		xi, xii := int64(vertices[i][0]), int64(vertices[(i+1)%len(vertices)][0])
		yi, yii := int64(vertices[i][1]), int64(vertices[(i+1)%len(vertices)][1])

		area += (yi + yii) * (xi - xii)
	}

	// fmt.Println(area)

	if area < 0 {
		return -(area / 2) + int64(circumference/2) + 1
	}

	return area/2 + int64(circumference/2) + 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()

		var directionStr string
		var steps int
		var color string
		_, err := fmt.Sscanf(line, "%s %d %s", &directionStr, &steps, &color)
		if err != nil {
			log.Fatal(err)
		}

		var direction int
		switch directionStr {
		case "U":
			direction = Up
		case "D":
			direction = Down
		case "L":
			direction = Left
		case "R":
			direction = Right
		}

		instructions = append(instructions, Instruction{
			direction: direction,
			steps:     steps,
			color:     color[1 : len(color)-1],
		})
	}

	ans := solvePart2(instructions)
	fmt.Println("Answer:", ans)
	// 952408144115

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
