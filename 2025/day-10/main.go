package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Machine struct {
	diagram      string
	buttons      [][]int
	requirements []int
}

type Equation struct {
	variables []int
	result    int
}

func toState(lights []int, sz int) string {
	var sb strings.Builder

	idx := 0
	for i := 0; i < sz; i++ {
		if idx < len(lights) && lights[idx] == i {
			sb.WriteString("#")
			idx++
		} else {
			sb.WriteString(".")
		}
	}

	return sb.String()
}

func mergeState(state1, state2 string) string {
	var sb strings.Builder
	for i := 0; i < len(state1); i++ {
		if (state1[i] == '#' && state2[i] == '.') || (state1[i] == '.' && state2[i] == '#') {
			sb.WriteString("#")
		} else {
			sb.WriteString(".")
		}
	}

	return sb.String()
}

func part1(machines []Machine) {
	ans := 0

	for _, machine := range machines {
		dp := make(map[string]int, 0)
		for _, buttonParts := range machine.buttons {
			state := toState(buttonParts, len(machine.diagram))

			for k, v := range dp {
				newState := mergeState(state, k)

				cur, found := dp[newState]
				if found {
					dp[newState] = min(cur, v+1)
				} else {
					dp[newState] = v + 1
				}
			}

			dp[state] = 1
		}

		ans += dp[machine.diagram]
	}

	fmt.Println("Part 1:", ans)
}

// Fraction represents a rational number for exact arithmetic
type Fraction struct {
	num, den int
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func newFraction(num, den int) Fraction {
	if den == 0 {
		panic("division by zero")
	}
	if den < 0 {
		num, den = -num, -den
	}
	g := gcd(num, den)
	if g != 0 {
		num, den = num/g, den/g
	}
	return Fraction{num, den}
}

func (f Fraction) add(other Fraction) Fraction {
	return newFraction(f.num*other.den+other.num*f.den, f.den*other.den)
}

func (f Fraction) sub(other Fraction) Fraction {
	return newFraction(f.num*other.den-other.num*f.den, f.den*other.den)
}

func (f Fraction) mul(other Fraction) Fraction {
	return newFraction(f.num*other.num, f.den*other.den)
}

func (f Fraction) div(other Fraction) Fraction {
	return newFraction(f.num*other.den, f.den*other.num)
}

func (f Fraction) isZero() bool {
	return f.num == 0
}

func (f Fraction) toInt() (int, bool) {
	if f.den == 1 || f.den == -1 {
		return f.num * f.den, true
	}
	if f.num%f.den == 0 {
		return f.num / f.den, true
	}
	return 0, false
}

// Matrix row: coefficients for each variable + constant term (last element)
// Represents: coef[0]*x0 + coef[1]*x1 + ... = constant
type Row []Fraction

func gaussianElimination(matrix []Row, numVars int) ([]Row, []int, []int) {
	rows := len(matrix)
	if rows == 0 {
		return matrix, nil, nil
	}

	// Copy matrix
	m := make([]Row, rows)
	for i := range matrix {
		m[i] = make(Row, len(matrix[i]))
		copy(m[i], matrix[i])
	}

	pivotCols := make([]int, 0) // which columns are pivot columns
	freeVars := make([]int, 0)  // which variables are free

	currentRow := 0
	for col := 0; col < numVars && currentRow < rows; col++ {
		// Find pivot
		pivotRow := -1
		for r := currentRow; r < rows; r++ {
			if !m[r][col].isZero() {
				pivotRow = r
				break
			}
		}

		if pivotRow == -1 {
			freeVars = append(freeVars, col)
			continue
		}

		// Swap rows
		m[currentRow], m[pivotRow] = m[pivotRow], m[currentRow]

		// Scale pivot row
		pivot := m[currentRow][col]
		for j := range m[currentRow] {
			m[currentRow][j] = m[currentRow][j].div(pivot)
		}

		// Eliminate column in other rows
		for r := 0; r < rows; r++ {
			if r == currentRow || m[r][col].isZero() {
				continue
			}
			factor := m[r][col]
			for j := range m[r] {
				m[r][j] = m[r][j].sub(factor.mul(m[currentRow][j]))
			}
		}

		pivotCols = append(pivotCols, col)
		currentRow++
	}

	// Add remaining variables as free
	for col := 0; col < numVars; col++ {
		if !slices.Contains(pivotCols, col) && !slices.Contains(freeVars, col) {
			freeVars = append(freeVars, col)
		}
	}

	return m, pivotCols, freeVars
}

// Given values for free variables, compute pivot variables
// Returns nil if solution is invalid (non-integer or negative)
func computeSolution(matrix []Row, pivotCols []int, freeVars []int, freeValues []int, numVars int) []int {
	solution := make([]int, numVars)

	// Set free variables
	for i, fv := range freeVars {
		solution[fv] = freeValues[i]
	}

	// Compute pivot variables from bottom up
	for i := len(pivotCols) - 1; i >= 0; i-- {
		pivotCol := pivotCols[i]
		// Find the row with this pivot
		row := -1
		for r := 0; r < len(matrix); r++ {
			if !matrix[r][pivotCol].isZero() {
				// Check if this is a pivot (coefficient is 1)
				if v, ok := matrix[r][pivotCol].toInt(); ok && v == 1 {
					row = r
					break
				}
			}
		}
		if row == -1 {
			continue
		}

		// value = constant - sum(coef[j] * solution[j]) for j != pivotCol
		value := matrix[row][numVars] // constant term
		for j := 0; j < numVars; j++ {
			if j != pivotCol {
				value = value.sub(matrix[row][j].mul(newFraction(solution[j], 1)))
			}
		}

		intVal, ok := value.toInt()
		if !ok || intVal < 0 {
			return nil
		}
		solution[pivotCol] = intVal
	}

	return solution
}

func solveMachineGaussian(buttons [][]int, requirements []int) int {
	numCounters := len(requirements)
	numButtons := len(buttons)

	if numButtons == 0 {
		// Check if all requirements are 0
		for _, r := range requirements {
			if r != 0 {
				return math.MaxInt
			}
		}
		return 0
	}

	// Build matrix: each row is a counter equation
	// coef[j] = 1 if button j affects this counter, 0 otherwise
	matrix := make([]Row, numCounters)
	for i := 0; i < numCounters; i++ {
		matrix[i] = make(Row, numButtons+1)
		for j := 0; j <= numButtons; j++ {
			matrix[i][j] = newFraction(0, 1)
		}
		matrix[i][numButtons] = newFraction(requirements[i], 1)
	}

	for btnIdx, btn := range buttons {
		for _, counter := range btn {
			if counter < numCounters {
				matrix[counter][btnIdx] = newFraction(1, 1)
			}
		}
	}

	// Perform Gaussian elimination
	reduced, pivotCols, freeVars := gaussianElimination(matrix, numButtons)

	// Now brute force over free variables only
	best := math.MaxInt

	// If no free variables, just compute the unique solution
	if len(freeVars) == 0 {
		solution := computeSolution(reduced, pivotCols, freeVars, []int{}, numButtons)
		if solution == nil {
			return math.MaxInt
		}
		total := 0
		for _, v := range solution {
			if v < 0 {
				return math.MaxInt
			}
			total += v
		}
		return total
	}

	// Compute upper bounds for free variables
	// Use max requirement as a safe upper bound
	upperBounds := make([]int, len(freeVars))
	for i := range freeVars {
		upperBounds[i] = 0
		for _, req := range requirements {
			upperBounds[i] = max(upperBounds[i], req)
		}
	}

	// Recursive search over free variables
	var search func(idx int, freeValues []int, currentSum int)
	search = func(idx int, freeValues []int, currentSum int) {
		if currentSum >= best {
			return
		}

		if idx == len(freeVars) {
			solution := computeSolution(reduced, pivotCols, freeVars, freeValues, numButtons)
			if solution == nil {
				return
			}
			total := 0
			for _, v := range solution {
				if v < 0 {
					return
				}
				total += v
			}
			best = min(best, total)
			return
		}

		for v := 0; v <= upperBounds[idx]; v++ {
			freeValues[idx] = v
			search(idx+1, freeValues, currentSum+v)
		}
	}

	freeValues := make([]int, len(freeVars))
	search(0, freeValues, 0)

	return best
}

func part2(machines []Machine) {
	ans := 0

	for _, machine := range machines {
		res := solveMachineGaussian(machine.buttons, machine.requirements)
		ans += res
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

	machines := make([]Machine, 0)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		diagram := strings.Trim(parts[0], "[]")
		buttons := make([][]int, 0)
		for i := 1; i < len(parts)-1; i++ {
			buttonStr := strings.Trim(parts[i], "()")
			buttonParts := strings.Split(buttonStr, ",")

			lights := make([]int, 0)
			for _, part := range buttonParts {
				num, err := strconv.Atoi(part)
				if err != nil {
					log.Fatal(err)
				}
				lights = append(lights, num)
			}

			sort.Slice(lights, func(i, j int) bool {
				return lights[i] < lights[j]
			})

			buttons = append(buttons, lights)
		}

		requirements := make([]int, 0)
		requriementsStr := strings.Trim(parts[len(parts)-1], "{}")
		requirementsParts := strings.Split(requriementsStr, ",")
		for _, part := range requirementsParts {
			num, err := strconv.Atoi(part)
			if err != nil {
				log.Fatal(err)
			}
			requirements = append(requirements, num)
		}

		machines = append(machines, Machine{
			diagram:      diagram,
			buttons:      buttons,
			requirements: requirements,
		})
	}

	part2(machines)
}
