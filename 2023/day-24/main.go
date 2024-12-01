package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Hailstone struct {
	px float64
	py float64
	pz float64

	vx float64
	vy float64
	vz float64
}

func crossAt(h1, h2 Hailstone) (bool, float64, float64) {
	// p1x + x * v1x = p2x + x * v2x
	// p1y + y * v1y = p2y + y * v2y

	b := (h1.vx*h2.py - h1.vx*h1.py - h1.vy*h2.px + h1.vy*h1.px) / (h2.vx*h1.vy - h1.vx*h2.vy)
	a := (h2.px - h1.px + b*h2.vx) / h1.vx

	return a >= 0 && b >= 0, h1.px + a*h1.vx, h1.py + a*h1.vy
}

func solvePart1(hailstones []Hailstone) int {
	const MIN_CROSS_X = 200000000000000
	const MIN_CROSS_Y = 200000000000000

	const MAX_CROSS_X = 400000000000000
	const MAX_CROSS_Y = 400000000000000

	ans := 0
	for i := 0; i < len(hailstones); i++ {
		for j := i + 1; j < len(hailstones); j++ {
			doesCross, crossX, crossY := crossAt(hailstones[i], hailstones[j])

			if doesCross && crossX >= MIN_CROSS_X && crossX <= MAX_CROSS_X && crossY >= MIN_CROSS_Y && crossY <= MAX_CROSS_X {
				// fmt.Println(i, j)
				ans += 1
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

	hailstones := make([]Hailstone, 0)
	for scanner.Scan() {
		line := scanner.Text()

		var hailstone Hailstone
		fmt.Sscanf(line, "%f, %f, %f @ %f, %f, %f",
			&hailstone.px, &hailstone.py, &hailstone.pz, &hailstone.vx, &hailstone.vy, &hailstone.vz)

		hailstones = append(hailstones, hailstone)
	}

	ans := solvePart1(hailstones)
	fmt.Println("Answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
