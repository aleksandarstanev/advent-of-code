package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(springs string, expectedGroups []int) bool {
	cnt := 0
	groups := make([]int, 0)
	for _, spring := range springs {
		if spring == '#' {
			cnt += 1
		} else {
			if cnt > 0 {
				groups = append(groups, cnt)
				cnt = 0
			}
		}
	}

	if cnt > 0 {
		groups = append(groups, cnt)
	}

	if len(groups) != len(expectedGroups) {
		return false
	}

	for i := 0; i < len(groups); i++ {
		if groups[i] != expectedGroups[i] {
			return false
		}
	}

	return true
}

func solvePart1(springs string, groups []int, idx int, count *int) {
	if idx == len(springs) {
		if check(springs, groups) {
			*count += 1
		}
		return
	}

	if springs[idx] == '#' || springs[idx] == '.' {
		solvePart1(springs, groups, idx+1, count)
	} else {
		withDamagedSpring := springs[:idx] + "#" + springs[idx+1:]
		solvePart1(withDamagedSpring, groups, idx+1, count)

		withOperationalSpring := springs[:idx] + "." + springs[idx+1:]
		solvePart1(withOperationalSpring, groups, idx+1, count)
	}

}

func expandSprings(springs string) string {
	return springs + "?" + springs + "?" + springs + "?" + springs + "?" + springs
}

func expandGroups(groups []int) []int {
	expandedGroups := make([]int, 0)
	for i := 0; i < 5; i++ {
		expandedGroups = append(expandedGroups, groups...)
	}

	return expandedGroups
}

func solvePart2(springs string, groups []int) int {
	springs = expandSprings(springs)
	groups = expandGroups(groups)

	dp := make([][]int, len(groups)+1)
	for i := 0; i < len(groups)+1; i++ {
		dp[i] = make([]int, len(springs)+1)
	}

	firstPos := -1
	for i := 0; i < len(springs); i++ {
		if springs[i] == '#' && firstPos == -1 {
			firstPos = i
		}

		startIdx := i - groups[0] + 1
		if firstPos != -1 && startIdx > firstPos {
			break
		}

		if startIdx < 0 {
			continue
		}

		substring := springs[startIdx:(i + 1)]

		works := true
		for _, c := range substring {
			if c == '.' {
				works = false
				break
			}
		}

		if works {
			dp[0][i] = 1
		}
	}

	for i := 1; i < len(groups); i++ {
		for start := 0; start < len(springs); start++ {
			var rollingStr string
			firstPos := -1
			for end := start + 1; end < len(springs); end++ {
				if springs[end] == '#' && firstPos == -1 {
					firstPos = end
				}

				rollingStr += string(springs[end])

				startIdx := (end - groups[i] + 1)
				if firstPos != -1 && startIdx > firstPos {
					break
				}

				if len(rollingStr) <= groups[i] {
					continue
				}

				diff := len(rollingStr) - groups[i]

				works := true
				for idx, c := range rollingStr {
					if idx < diff {
						if c == '#' {
							works = false
							break
						}
					} else {
						if c == '.' {
							works = false
							break
						}
					}
				}

				if works {
					dp[i][end] = dp[i][end] + dp[i-1][start]
				}
			}
		}
	}

	ans := 0
	for i := len(springs) - 1; i >= 0; i-- {
		ans += dp[len(groups)-1][i]

		if springs[i] == '#' {
			break
		}
	}

	// for i := 0; i < len(groups); i++ {
	// 	for j := 0; j < len(springs); j++ {
	// 		fmt.Print(dp[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }

	return ans
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ans := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		springs := parts[0]
		groupPart := parts[1]

		groups := make([]int, 0)
		for _, group := range strings.Split(groupPart, ",") {
			num, err := strconv.Atoi(group)
			if err != nil {
				log.Fatal(err)
			}

			groups = append(groups, num)
		}

		count := solvePart2(springs, groups)

		fmt.Println(count)
		ans += count
	}

	fmt.Println("Answer:", ans)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
