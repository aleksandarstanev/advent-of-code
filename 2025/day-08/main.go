package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Box struct {
	x, y, z int
}

type Distance struct {
	from, to Box
	dist     int
}

func getDistance(first, second Box) int {
	return (first.x-second.x)*(first.x-second.x) + (first.y-second.y)*(first.y-second.y) + (first.z-second.z)*(first.z-second.z)
}

func areEqual(first, second Box) bool {
	return first.x == second.x && first.y == second.y && first.z == second.z
}

func getRoot(box Box, parent map[Box]Box) Box {
	for !areEqual(box, parent[box]) {
		box = parent[box]
	}

	return box
}

func join(first, second Box, parent map[Box]Box, numChildren map[Box]int) {
	firstRoot := getRoot(first, parent)
	secondRoot := getRoot(second, parent)
	if areEqual(firstRoot, secondRoot) {
		return
	}

	parent[secondRoot] = firstRoot
	numChildren[firstRoot] += numChildren[secondRoot]
	numChildren[secondRoot] = 0
}

func getDistances(boxes []Box) []Distance {
	distances := make([]Distance, 0)
	for i := 0; i < len(boxes); i++ {
		for j := i + 1; j < len(boxes); j++ {
			distances = append(distances, Distance{
				from: boxes[i],
				to:   boxes[j],
				dist: getDistance(boxes[i], boxes[j]),
			})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	return distances
}

func part1(boxes []Box, connections int) {
	distances := getDistances(boxes)

	parent := make(map[Box]Box, 0)
	numChildren := make(map[Box]int, 0)

	for _, box := range boxes {
		parent[box] = box
		numChildren[box] = 1
	}

	for i := 0; i < connections; i++ {
		join(distances[i].from, distances[i].to, parent, numChildren)
	}

	sizes := make([]int, 0)
	for _, v := range numChildren {
		sizes = append(sizes, v)
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	ans := sizes[0] * sizes[1] * sizes[2]

	fmt.Println("Part 1:", ans)
}

func part2(boxes []Box) {
	distances := getDistances(boxes)

	parent := make(map[Box]Box, 0)
	numChildren := make(map[Box]int, 0)

	for _, box := range boxes {
		parent[box] = box
		numChildren[box] = 1
	}

	ans := 0
	for i := 0; ; i++ {
		join(distances[i].from, distances[i].to, parent, numChildren)

		root := getRoot(distances[i].from, parent)

		if numChildren[root] == len(boxes) {
			ans = distances[i].from.x * distances[i].to.x
			break
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

	var boxes []Box
	for scanner.Scan() {
		line := scanner.Text()

		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			log.Fatal(err)
		}

		boxes = append(boxes, Box{
			x: x,
			y: y,
			z: z,
		})
	}

	// part1(boxes, 1000)
	part2(boxes)
}
