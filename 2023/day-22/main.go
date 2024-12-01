package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Brick struct {
	x int
	y int
	z int
}

type Cube struct {
	brick1 Brick
	brick2 Brick
}

func parseLine(line string) Cube {
	var cube Cube
	fmt.Sscanf(line, "%d,%d,%d~%d,%d,%d", &cube.brick1.x, &cube.brick1.y, &cube.brick1.z, &cube.brick2.x, &cube.brick2.y, &cube.brick2.z)
	return cube
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func canFit(cube Cube, minZ int, occupied map[Brick]bool) bool {
	diff := max(cube.brick1.z, cube.brick2.z) - min(cube.brick1.z, cube.brick2.z)
	maxZ := minZ + diff

	for z := minZ; z <= maxZ; z++ {
		minX, maxX := min(cube.brick1.x, cube.brick2.x), max(cube.brick1.x, cube.brick2.x)
		minY, maxY := min(cube.brick1.y, cube.brick2.y), max(cube.brick1.y, cube.brick2.y)

		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				if occupied[Brick{x, y, z}] {
					return false
				}
			}
		}
	}

	return true
}

func placeCube(cube Cube, occupied map[Brick]bool) {
	minZ, maxZ := min(cube.brick1.z, cube.brick2.z), max(cube.brick1.z, cube.brick2.z)
	minX, maxX := min(cube.brick1.x, cube.brick2.x), max(cube.brick1.x, cube.brick2.x)
	minY, maxY := min(cube.brick1.y, cube.brick2.y), max(cube.brick1.y, cube.brick2.y)

	for z := minZ; z <= maxZ; z++ {
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				occupied[Brick{x, y, z}] = true
			}
		}
	}
}

func isCubeSupported(cube, supportedBy Cube) bool {
	if cube.brick1.z != supportedBy.brick2.z+1 {
		return false
	}

	minX, maxX := min(supportedBy.brick1.x, supportedBy.brick2.x), max(supportedBy.brick1.x, supportedBy.brick2.x)
	minY, maxY := min(supportedBy.brick1.y, supportedBy.brick2.y), max(supportedBy.brick1.y, supportedBy.brick2.y)

	minCubeX, maxCubeX := min(cube.brick1.x, cube.brick2.x), max(cube.brick1.x, cube.brick2.x)
	minCubeY, maxCubeY := min(cube.brick1.y, cube.brick2.y), max(cube.brick1.y, cube.brick2.y)

	occupies := make(map[[2]int]bool)
	for x := minCubeX; x <= maxCubeX; x++ {
		for y := minCubeY; y <= maxCubeY; y++ {
			occupies[[2]int{x, y}] = true
		}
	}

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if occupies[[2]int{x, y}] {
				return true
			}
		}
	}

	// if minCubeX >= minX && maxCubeX <= minX && minCubeY >= minY && maxCubeY <= minY ||
	// 	minCubeX >= maxX && maxCubeX <= maxX && minCubeY >= minY && maxCubeY <= minY ||
	// 	minCubeX >= minX && maxCubeX <= minX && minCubeY >= maxY && maxCubeY <= maxY ||
	// 	minCubeX >= maxX && maxCubeX <= maxX && minCubeY >= maxY && maxCubeY <= maxY {
	// 	return true
	// }

	return false
}

func solvePart1(cubes []Cube) int {
	fallenCubes := make([]Cube, 0)

	sort.Slice(cubes, func(i, j int) bool {
		return cubes[i].brick1.z < cubes[j].brick1.z
	})

	occupied := make(map[Brick]bool)
	for _, cube := range cubes {
		z, minZ := cube.brick1.z, cube.brick1.z
		for z > 0 {
			if canFit(cube, z, occupied) {
				minZ = z
			} else {
				break
			}

			z--
		}

		fallenCube := Cube{
			Brick{cube.brick1.x, cube.brick1.y, minZ},
			Brick{cube.brick2.x, cube.brick2.y, minZ + max(cube.brick1.z, cube.brick2.z) - min(cube.brick1.z, cube.brick2.z)},
		}

		fallenCubes = append(fallenCubes, fallenCube)

		placeCube(fallenCube, occupied)
	}

	supportedBy := make(map[int][]int)

	nonRemovable := make(map[int]bool)
	for i := 0; i < len(fallenCubes); i++ {
		for j := 0; j < len(fallenCubes); j++ {
			if i == j {
				continue
			}

			if isCubeSupported(fallenCubes[i], fallenCubes[j]) {
				supportedBy[i] = append(supportedBy[i], j)
			}
		}

		if len(supportedBy[i]) == 1 {
			nonRemovable[supportedBy[i][0]] = true
		}
	}

	ans := len(fallenCubes) - len(nonRemovable)
	// fmt.Println(fallenCubes)

	return ans
}

func solvePart2(cubes []Cube) int {
	fallenCubes := make([]Cube, 0)

	sort.Slice(cubes, func(i, j int) bool {
		return cubes[i].brick1.z < cubes[j].brick1.z
	})

	occupied := make(map[Brick]bool)
	for _, cube := range cubes {
		z, minZ := cube.brick1.z, cube.brick1.z
		for z > 0 {
			if canFit(cube, z, occupied) {
				minZ = z
			} else {
				break
			}

			z--
		}

		fallenCube := Cube{
			Brick{cube.brick1.x, cube.brick1.y, minZ},
			Brick{cube.brick2.x, cube.brick2.y, minZ + max(cube.brick1.z, cube.brick2.z) - min(cube.brick1.z, cube.brick2.z)},
		}

		fallenCubes = append(fallenCubes, fallenCube)

		placeCube(fallenCube, occupied)
	}

	supportedBy := make(map[int][]int)

	nonRemovable := make(map[int]bool)
	for i := 0; i < len(fallenCubes); i++ {
		for j := 0; j < len(fallenCubes); j++ {
			if i == j {
				continue
			}

			if isCubeSupported(fallenCubes[i], fallenCubes[j]) {
				supportedBy[i] = append(supportedBy[i], j)
			}
		}

		if len(supportedBy[i]) == 1 {
			nonRemovable[supportedBy[i][0]] = true
		}
	}

	ans := 0
	for i := range nonRemovable {
		removed := make(map[int]bool)
		removed[i] = true
		cur := 0
		hasFallen := true
		for hasFallen {
			hasFallen = false
			for j := 0; j < len(fallenCubes); j++ {
				// fmt.Println("j = ", j)
				if removed[j] {
					continue
				}

				falls := len(supportedBy[j]) > 0
				for _, k := range supportedBy[j] {
					if _, ok := removed[k]; !ok {
						falls = false
					}
				}

				if falls {
					removed[j] = true
					// fmt.Println(j)
					cur++
					hasFallen = true
				}
			}
		}

		ans += cur
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

	cubes := make([]Cube, 0)
	for scanner.Scan() {
		line := scanner.Text()
		cubes = append(cubes, parseLine(line))
	}

	// fmt.Println(parts)

	ans := solvePart2(cubes)
	fmt.Println("Answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
