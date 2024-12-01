package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeLine := scanner.Text()

	times := strings.Fields(timeLine[5:])

	scanner.Scan()
	distanceLine := scanner.Text()
	distances := strings.Fields(distanceLine[9:])

	ans := 1
	eps := 0.0000001
	for i := 0; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			log.Fatal(err)
		}

		dist, err := strconv.Atoi(distances[i])
		if err != nil {
			log.Fatal(err)
		}

		minAns := int(math.Ceil((float64(time)-math.Sqrt(float64(time*time-4*dist)))/2 + eps))
		maxAns := int(math.Floor((float64(time)+math.Sqrt(float64(time*time-4*dist)))/2 - eps))

		ans = ans * (maxAns - minAns + 1)
	}

	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	timeLine := scanner.Text()

	times := strings.Fields(timeLine[5:])

	scanner.Scan()
	distanceLine := scanner.Text()
	distances := strings.Fields(distanceLine[9:])

	eps := 0.0000001
	time, err := strconv.Atoi(strings.Join(times, ""))
	if err != nil {
		log.Fatal(err)
	}

	dist, err := strconv.Atoi(strings.Join(distances, ""))
	if err != nil {
		log.Fatal(err)
	}

	minAns := int(math.Ceil((float64(time)-math.Sqrt(float64(time*time-4*dist)))/2 + eps))
	maxAns := int(math.Floor((float64(time)+math.Sqrt(float64(time*time-4*dist)))/2 - eps))

	ans := (maxAns - minAns + 1)

	fmt.Println(ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
