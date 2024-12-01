package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Part struct {
	x int
	m int
	a int
	s int
}

type Ranges struct {
	fromX int
	toX   int

	fromM int
	toM   int

	fromA int
	toA   int

	fromS int
	toS   int
}

type Node struct {
	ranges Ranges
	label  string
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func (r Ranges) Combinations() int {
	return (r.toX - r.fromX + 1) * (r.toM - r.fromM + 1) * (r.toA - r.fromA + 1) * (r.toS - r.fromS + 1)
}

func (r Ranges) Split(variable string, sign string, number int) []Ranges {
	ranges := make([]Ranges, 0)

	switch variable {
	case "x":
		if sign == ">" {
			newFromX := max(number+1, r.fromX)

			if newFromX <= r.toX {
				ranges = append(ranges, Ranges{
					fromX: newFromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}

			if r.fromX < newFromX {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   newFromX - 1,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}
		}

		if sign == "<" {
			newToX := min(number-1, r.toX)

			if r.fromX <= newToX {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   newToX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}

			if newToX < r.toX {
				ranges = append(ranges, Ranges{
					fromX: newToX + 1,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}
		}
	case "m":
		if sign == ">" {
			newFromM := max(number+1, r.fromM)

			if newFromM <= r.toM {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: newFromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}

			if r.fromM < newFromM {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   newFromM - 1,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}
		}

		if sign == "<" {
			newToM := min(number-1, r.toM)

			if r.fromM <= newToM {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   newToM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}

			if newToM < r.toM {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: newToM + 1,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}
		}
	case "a":
		if sign == ">" {
			newFromA := max(number+1, r.fromA)

			if newFromA <= r.toA {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: newFromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}

			if r.fromA < newFromA {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   newFromA - 1,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}
		}

		if sign == "<" {
			newToA := min(number-1, r.toA)

			if r.fromA <= newToA {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   newToA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}

			if newToA < r.toA {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: newToA + 1,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   r.toS,
				})
			}
		}

	case "s":
		if sign == ">" {
			newFromS := max(number+1, r.fromS)

			if newFromS <= r.toS {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: newFromS,
					toS:   r.toS,
				})
			}

			if r.fromS < newFromS {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   newFromS - 1,
				})
			}
		}

		if sign == "<" {
			newToS := min(number-1, r.toS)

			if r.fromS <= newToS {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: r.fromS,
					toS:   newToS,
				})
			}

			if newToS < r.toS {
				ranges = append(ranges, Ranges{
					fromX: r.fromX,
					toX:   r.toX,
					fromM: r.fromM,
					toM:   r.toM,
					fromA: r.fromA,
					toA:   r.toA,
					fromS: newToS + 1,
					toS:   r.toS,
				})
			}
		}
	}

	return ranges
}

func isAccepted(workflows map[string]string, part Part) bool {
	currentNode := "in"
	for currentNode != "A" && currentNode != "R" {
		// fmt.Println(currentNode)
		workflow := workflows[currentNode]

		conditions := strings.Split(workflow, ",")
		for _, condition := range conditions {
			conditionParts := strings.Split(condition, ":")
			if len(conditionParts) == 1 {
				currentNode = condition
				break
			}

			variable := string(conditionParts[0][0])
			sign := string(conditionParts[0][1])
			number, err := strconv.Atoi(conditionParts[0][2:])
			if err != nil {
				log.Fatal(err)
			}

			// fmt.Println(string(variable), string(sign), number, part)
			newNode := false

			switch variable {
			case "x":
				if sign == ">" && part.x > number {
					currentNode = conditionParts[1]
					newNode = true
				}

				if sign == "<" && part.x < number {
					currentNode = conditionParts[1]
					newNode = true
				}

			case "m":
				if sign == ">" && part.m > number {
					currentNode = conditionParts[1]
					newNode = true
				}

				if sign == "<" && part.m < number {
					currentNode = conditionParts[1]
					newNode = true
				}

			case "a":
				if sign == ">" && part.a > number {
					currentNode = conditionParts[1]
					newNode = true
				}

				if sign == "<" && part.a < number {
					currentNode = conditionParts[1]
					newNode = true
				}

			case "s":
				if sign == ">" && part.s > number {
					currentNode = conditionParts[1]
					newNode = true
				}

				if sign == "<" && part.s < number {
					currentNode = conditionParts[1]
					newNode = true
				}
			}

			if newNode {
				break
			}
		}
	}

	return currentNode == "A"
}

func solvePart1(workflows map[string]string, parts []Part) int {
	ans := 0
	for _, part := range parts {
		if isAccepted(workflows, part) {
			ans += part.x
			ans += part.m
			ans += part.a
			ans += part.s
		}
	}

	return ans
}

func solvePart2(workflows map[string]string) int {
	startRanges := Ranges{
		fromX: 1,
		toX:   4000,
		fromM: 1,
		toM:   4000,
		fromA: 1,
		toA:   4000,
		fromS: 1,
		toS:   4000,
	}

	queue := make([]Node, 0)
	queue = append(queue, Node{
		ranges: startRanges,
		label:  "in",
	})
	ans := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// fmt.Println(node)

		ranges := node.ranges

		if node.label == "A" {
			ans += ranges.Combinations()
			continue
		}

		if node.label == "R" {
			continue
		}

		workflow := workflows[node.label]

		conditions := strings.Split(workflow, ",")
		for _, condition := range conditions {
			conditionParts := strings.Split(condition, ":")
			if len(conditionParts) == 1 {
				queue = append(queue, Node{
					ranges: ranges,
					label:  condition,
				})

				break
			}

			variable := string(conditionParts[0][0])
			sign := string(conditionParts[0][1])
			number, err := strconv.Atoi(conditionParts[0][2:])
			if err != nil {
				log.Fatal(err)
			}

			newRanges := ranges.Split(variable, sign, number)

			queue = append(queue, Node{
				ranges: newRanges[0],
				label:  conditionParts[1],
			})

			ranges = newRanges[1]
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

	workflows := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		parts := strings.Split(line, "{")
		workflows[parts[0]] = parts[1][:len(parts[1])-1]
	}

	parts := make([]Part, 0)
	for scanner.Scan() {
		line := scanner.Text()

		var x, m, a, s int
		_, err := fmt.Sscanf(line, "{x=%d,m=%d,a=%d,s=%d}", &x, &m, &a, &s)
		if err != nil {
			log.Fatal(err)
		}

		parts = append(parts, Part{x, m, a, s})
	}

	// fmt.Println(parts)

	ans := solvePart2(workflows)
	fmt.Println("Answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
