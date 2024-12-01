package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
)

const MAX_CONSECUTIVE_STEPS_PART_1 = 3
const MAX_CONSECUTIVE_STEPS_PART_2 = 10
const MIN_STEPS_BEFORE_TURNING = 4

const (
	Up = iota
	Down
	Left
	Right
)

type State struct {
	row              int
	col              int
	direction        int
	consecutiveSteps int
	heatLost         int
}

type StateNoCost struct {
	row              int
	col              int
	direction        int
	consecutiveSteps int
}

func isValid(grid []string, row int, col int) bool {
	rows, cols := len(grid), len(grid[0])
	return row >= 0 && row < rows && col >= 0 && col < cols
}

func (s State) resolveNextStatesPart1(grid []string) []*State {
	possibleNextPositions := make([][4]int, 0)

	switch s.direction {
	case Up:
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row - 1, s.col, Up, s.consecutiveSteps + 1})
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col - 1, Left, 1})
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col + 1, Right, 1})
	case Down:
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row + 1, s.col, Down, s.consecutiveSteps + 1})
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col - 1, Left, 1})
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col + 1, Right, 1})
	case Left:
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col - 1, Left, s.consecutiveSteps + 1})
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row - 1, s.col, Up, 1})
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row + 1, s.col, Down, 1})
	case Right:
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col + 1, Right, s.consecutiveSteps + 1})
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row - 1, s.col, Up, 1})
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row + 1, s.col, Down, 1})
	}

	nextStates := make([]*State, 0)
	for _, nextPosition := range possibleNextPositions {
		row, col, direction, consecutiveSteps := nextPosition[0], nextPosition[1], nextPosition[2], nextPosition[3]
		if consecutiveSteps <= MAX_CONSECUTIVE_STEPS_PART_1 && isValid(grid, row, col) {
			nextStates = append(nextStates, &State{row, col, direction, consecutiveSteps, s.heatLost + int(grid[row][col]-'0')})
		}
	}

	return nextStates
}

func (s State) resolveNextStatesPart2(grid []string) []*State {
	possibleNextPositions := make([][4]int, 0)

	switch s.direction {
	case Up:
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row - 1, s.col, Up, s.consecutiveSteps + 1})
		if s.consecutiveSteps >= MIN_STEPS_BEFORE_TURNING {
			possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col - 1, Left, 1})
			possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col + 1, Right, 1})
		}
	case Down:
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row + 1, s.col, Down, s.consecutiveSteps + 1})
		if s.consecutiveSteps >= MIN_STEPS_BEFORE_TURNING {
			possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col - 1, Left, 1})
			possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col + 1, Right, 1})
		}
	case Left:
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col - 1, Left, s.consecutiveSteps + 1})
		if s.consecutiveSteps >= MIN_STEPS_BEFORE_TURNING {
			possibleNextPositions = append(possibleNextPositions, [4]int{s.row - 1, s.col, Up, 1})
			possibleNextPositions = append(possibleNextPositions, [4]int{s.row + 1, s.col, Down, 1})
		}
	case Right:
		possibleNextPositions = append(possibleNextPositions, [4]int{s.row, s.col + 1, Right, s.consecutiveSteps + 1})
		if s.consecutiveSteps >= MIN_STEPS_BEFORE_TURNING {
			possibleNextPositions = append(possibleNextPositions, [4]int{s.row - 1, s.col, Up, 1})
			possibleNextPositions = append(possibleNextPositions, [4]int{s.row + 1, s.col, Down, 1})
		}
	}

	nextStates := make([]*State, 0)
	for _, nextPosition := range possibleNextPositions {
		row, col, direction, consecutiveSteps := nextPosition[0], nextPosition[1], nextPosition[2], nextPosition[3]
		if consecutiveSteps <= MAX_CONSECUTIVE_STEPS_PART_2 && isValid(grid, row, col) {
			nextStates = append(nextStates, &State{row, col, direction, consecutiveSteps, s.heatLost + int(grid[row][col]-'0')})
		}
	}

	return nextStates
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLost < pq[j].heatLost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(state any) {
	*pq = append(*pq, state.(*State))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	state := old[n-1]
	*pq = old[0 : n-1]
	return *state
}

func solvePart1(grid []string) int {
	rows, cols := len(grid), len(grid[0])
	pq := make(PriorityQueue, 0)

	visited := make(map[StateNoCost]bool)

	heap.Push(&pq, &State{0, 0, Right, 0, 0})
	heap.Push(&pq, &State{0, 0, Down, 0, 0})

	visited[StateNoCost{0, 0, Right, 0}] = true
	visited[StateNoCost{0, 0, Down, 0}] = true

	for pq.Len() > 0 {
		state := heap.Pop(&pq).(State)

		if state.row == rows-1 && state.col == cols-1 {
			return state.heatLost
		}

		nextStates := state.resolveNextStatesPart1(grid)
		for _, nextState := range nextStates {
			stateNoCost := StateNoCost{nextState.row, nextState.col, nextState.direction, nextState.consecutiveSteps}
			if _, ok := visited[stateNoCost]; !ok {
				heap.Push(&pq, nextState)
				visited[stateNoCost] = true
			}
		}
	}

	return -1
}

func solvePart2(grid []string) int {
	rows, cols := len(grid), len(grid[0])
	pq := make(PriorityQueue, 0)

	visited := make(map[StateNoCost]bool)

	heap.Push(&pq, &State{0, 0, Right, 0, 0})
	heap.Push(&pq, &State{0, 0, Down, 0, 0})

	visited[StateNoCost{0, 0, Right, 0}] = true
	visited[StateNoCost{0, 0, Down, 0}] = true

	for pq.Len() > 0 {
		state := heap.Pop(&pq).(State)

		if state.row == rows-1 && state.col == cols-1 && state.consecutiveSteps >= MIN_STEPS_BEFORE_TURNING {
			return state.heatLost
		}

		nextStates := state.resolveNextStatesPart2(grid)
		for _, nextState := range nextStates {
			stateNoCost := StateNoCost{nextState.row, nextState.col, nextState.direction, nextState.consecutiveSteps}
			if _, ok := visited[stateNoCost]; !ok {
				heap.Push(&pq, nextState)
				visited[stateNoCost] = true
			}
		}
	}

	return -1
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
		grid = append(grid, scanner.Text())
	}

	ans := solvePart2(grid)
	fmt.Println("Answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
