package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var finalNode string = "ZZZ"

type Node struct {
	left  string
	right string
}

func part1(instructions string, stringToNode map[string]Node) {
	steps := 0
	currentNode := "AAA"
	for currentNode != finalNode {
		node := stringToNode[currentNode]

		if instructions[steps%len(instructions)] == 'L' {
			currentNode = node.left
		} else {
			currentNode = node.right
		}

		steps = steps + 1
	}

	fmt.Println(steps)
}

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}

	return gcd(b%a, a)
}

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

func lcmSlice(a []int) int {
	result := a[0]
	for i := 1; i < len(a); i++ {
		result = lcm(result, a[i])
	}

	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	instructions := scanner.Text()

	scanner.Scan() // blank line

	stringToNode := make(map[string]Node)

	currentNodes := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		current := line[0:3]
		left := line[7:10]
		right := line[12:15]

		if current[len(current)-1] == 'A' {
			currentNodes = append(currentNodes, current)
		}

		stringToNode[current] = Node{left, right}
	}

	steps := 0
	areAllFinal := false
	prevFinal := make([]int, len(currentNodes))
	period := make([]int, len(currentNodes))
	for !areAllFinal {
		areAllFinal = true

		for i := 0; i < len(currentNodes); i++ {
			node := stringToNode[currentNodes[i]]

			if instructions[steps%len(instructions)] == 'L' {
				currentNodes[i] = node.left
			} else {
				currentNodes[i] = node.right
			}

			if currentNodes[i][len(currentNodes[i])-1] != 'Z' {
				areAllFinal = false
			} else {
				if prevFinal[i] != 0 && period[i] == 0 {
					period[i] = steps - prevFinal[i]
				}

				prevFinal[i] = steps
			}
		}

		allHavePeriod := true
		for i := 0; i < len(currentNodes); i++ {
			if period[i] == 0 {
				allHavePeriod = false
			}
		}

		if allHavePeriod {
			break
		}

		steps = steps + 1
	}

	lcm := lcmSlice(period)

	fmt.Println(lcm)

	// steps = lcm - 1
	// areAllFinal = false
	// for !areAllFinal {
	// 	areAllFinal = true

	// 	for i := 0; i < len(currentNodes); i++ {
	// 		node := stringToNode[currentNodes[i]]

	// 		if instructions[steps%len(instructions)] == 'L' {
	// 			currentNodes[i] = node.left
	// 		} else {
	// 			currentNodes[i] = node.right
	// 		}

	// 		if currentNodes[i][len(currentNodes[i])-1] != 'Z' {
	// 			areAllFinal = false
	// 		}
	// 	}

	// 	steps = steps + 1
	// }

	// fmt.Println(steps)
}
