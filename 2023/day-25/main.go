package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func dfs(v string, graph map[string][]string, visited map[string]bool, countEdgeUsage map[string]int) {
	visited[v] = true

	for _, u := range randomShuffle(graph[v]) {
		if !visited[u] {
			str := ""
			if v < u {
				str = v + "->" + u
			} else {
				str = u + "->" + v
			}

			countEdgeUsage[str]++

			dfs(u, graph, visited, countEdgeUsage)
		}
	}
}

func bfs(v string, graph map[string][]string, visited map[string]bool, countEdgeUsage map[string]int) {
	queue := make([]string, 0)
	queue = append(queue, v)
	visited[v] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		neighbors := randomShuffle(graph[v])
		for _, u := range neighbors {
			if !visited[u] {
				str := ""
				if v < u {
					str = v + "->" + u
				} else {
					str = u + "->" + v
				}

				countEdgeUsage[str]++

				visited[u] = true
				queue = append(queue, u)
			}
		}
	}
}

func dfsCountVertices(v string, graph map[string][]string, visited map[string]bool) int {
	visited[v] = true

	count := 1
	for _, u := range graph[v] {
		if !visited[u] {
			count += dfsCountVertices(u, graph, visited)
		}
	}

	return count
}

func remove(slice []string, element string) []string {
	newSlice := make([]string, 0)
	for _, v := range slice {
		if v == element {
			continue
		}

		newSlice = append(newSlice, v)
	}

	return newSlice
}

func randomShuffle(slice []string) []string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]string, len(slice))
	perm := r.Perm(len(slice))
	for i, randIndex := range perm {
		ret[i] = slice[randIndex]
	}

	return ret
}

func solvePart1(graph map[string][]string, vertices []string) int {
	const EDGES_TO_REMOVE = 3

	// graph["hfx"] = remove(graph["hfx"], "pzl")
	// graph["pzl"] = remove(graph["pzl"], "hfx")

	// graph["bvb"] = remove(graph["bvb"], "cmg")
	// graph["cmg"] = remove(graph["cmg"], "bvb")

	// graph["nvd"] = remove(graph["nvd"], "jqt")
	// graph["jqt"] = remove(graph["jqt"], "nvd")

	// 1487 vertices, 6666 edges

	for i := 0; i < EDGES_TO_REMOVE; i++ {
		countEdgeUsage := make(map[string]int)

		for _, v := range vertices {
			visited := make(map[string]bool)
			bfs(v, graph, visited, countEdgeUsage)
		}

		max := 0
		mostUsedEdge := ""
		for k, v := range countEdgeUsage {
			if v > max {
				max = v
				mostUsedEdge = k
			}
		}

		// fmt.Println(countEdgeUsage)

		parts := strings.Split(mostUsedEdge, "->")
		u, v := parts[0], parts[1]

		// fmt.Println(graph[u])
		// fmt.Println(graph[v])

		fmt.Println("Removing edge:", u, v)

		graph[u] = remove(graph[u], v)
		graph[v] = remove(graph[v], u)

		// fmt.Println(graph[u])
		// fmt.Println(graph[v])
	}

	visited := make(map[string]bool)
	ans := 1
	for _, v := range vertices {
		if !visited[v] {
			vertices := dfsCountVertices(v, graph, visited)
			ans *= vertices
			fmt.Println(vertices)
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

	graph := make(map[string][]string)

	verticesMap := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ": ")

		v := parts[0]
		verticesMap[v] = true

		for _, u := range strings.Split(parts[1], " ") {
			graph[v] = append(graph[v], u)
			graph[u] = append(graph[u], v)

			verticesMap[u] = true
		}
	}

	vertices := make([]string, 0, len(verticesMap))
	for v := range verticesMap {
		vertices = append(vertices, v)
	}

	ans := solvePart1(graph, vertices)
	fmt.Println("Answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
