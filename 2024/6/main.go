package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var (
	directions = map[string][2]int{
		"^": {-1, 0},
		">": {0, 1},
		"v": {1, 0},
		"<": {0, -1},
	}
	turns = map[string]string{
		"^": ">",
		">": "v",
		"v": "<",
		"<": "^",
	}
	dir2VisitIdx = map[string]int{
		"^": 0,
		">": 1,
		"v": 2,
		"<": 3,
	}
)

func main() {
	f, err := os.Open("./2024/6/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	area := make([][]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		area = append(area, strings.Split(l, ""))
	}

	fmt.Println(part1(area))
	fmt.Println(part2(area))
}

func part1(area [][]string) int {
	m := len(area)
	n := len(area[0])

	// find guard
	gi := 0
	gj := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if area[i][j] == "^" || area[i][j] == ">" ||
				area[i][j] == "v" || area[i][j] == "<" {
				gi = i
				gj = j
				break
			}
		}
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	count := 0
	for inRange(gi, gj, m, n) {

		if !visited[gi][gj] {
			count += 1
		}
		visited[gi][gj] = true

		dir := directions[area[gi][gj]]

		nexti := gi + dir[0]
		nextj := gj + dir[1]

		// guard is leaving the area
		if !inRange(nexti, nextj, m, n) {
			break
		}

		// turn
		if area[nexti][nextj] == "#" {
			area[gi][gj] = turns[area[gi][gj]]
		} else {
			area[nexti][nextj] = area[gi][gj]
			area[gi][gj] = "."
			gi = nexti
			gj = nextj
		}
	}

	return count
}

// wrong number, still slow
// func qeqe(area [][]string) int {
// 	m := len(area)
// 	n := len(area[0])

// 	// find guard
// 	gi := 0
// 	gj := 0
// 	for i := 0; i < m; i++ {
// 		for j := 0; j < n; j++ {
// 			if area[i][j] == "^" || area[i][j] == ">" ||
// 				area[i][j] == "v" || area[i][j] == "<" {
// 				gi = i
// 				gj = j
// 				break
// 			}
// 		}
// 	}

// 	visited := make([][][]string, m)
// 	for i := 0; i < m; i++ {
// 		visited[i] = make([][]string, n)
// 		for j := 0; j < n; j++ {
// 			visited[i][j] = make([]string, 4)
// 		}
// 	}

// 	obstacles := make([][]bool, m)
// 	for i := 0; i < m; i++ {
// 		obstacles[i] = make([]bool, n)
// 	}

// 	count := 0
// 	for inRange(gi, gj, m, n) {
// 		// printMatrix(area)
// 		// fmt.Println("======================")

// 		arrow := area[gi][gj]

// 		dir := directions[arrow]
// 		nexti := gi + dir[0]
// 		nextj := gj + dir[1]

// 		// guard is leaving the area
// 		if !inRange(nexti, nextj, m, n) {
// 			break
// 		}

// 		// turn
// 		if area[nexti][nextj] == "#" {
// 			area[gi][gj] = turns[arrow]
// 		} else {
// 			// test what if obstacle on the next cell
// 			guard := area[gi][gj]
// 			area[nexti][nextj] = "#"
// 			visitedCopy := copyVisited(visited)
// 			if checkCycle(area, gi, gj, visitedCopy) {
// 				if !obstacles[nexti][nextj] {
// 					count++
// 					obstacles[nexti][nextj] = true
// 				}
// 			}
// 			area[nexti][nextj] = "."
// 			area[gi][gj] = guard

// 			area[nexti][nextj] = arrow
// 			area[gi][gj] = "."
// 			gi = nexti
// 			gj = nextj
// 		}
// 	}

// 	return count
// }

func part2(area [][]string) int {
	m := len(area)
	n := len(area[0])

	// find guard
	gi := 0
	gj := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if area[i][j] == "^" || area[i][j] == ">" ||
				area[i][j] == "v" || area[i][j] == "<" {
				gi = i
				gj = j
				break
			}
		}
	}

	// brute force
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if area[i][j] == "." {
				visited := make([][][]string, m)
				for i := 0; i < m; i++ {
					visited[i] = make([][]string, n)
					for j := 0; j < n; j++ {
						visited[i][j] = make([]string, 4)
					}
				}

				// place obstacle
				curr := area[i][j]
				guard := area[gi][gj]
				area[i][j] = "#"
				if checkCycle(area, gi, gj, visited) {
					count++
				}
				area[gi][gj] = guard
				area[i][j] = curr
			}
		}
	}

	return count
}

func checkCycle(area [][]string, gi, gj int, visited [][][]string) bool {
	m := len(area)
	n := len(area[0])

	for inRange(gi, gj, m, n) {
		arrow := area[gi][gj]

		visited[gi][gj][dir2VisitIdx[arrow]] = arrow

		dir := directions[arrow]
		nexti := gi + dir[0]
		nextj := gj + dir[1]

		// guard is leaving the area
		if !inRange(nexti, nextj, m, n) {
			// "restore" initial pos
			area[gi][gj] = "."
			break
		}

		if slices.Index(visited[nexti][nextj], arrow) != -1 {
			// "restore" initial pos
			area[gi][gj] = "."
			return true
		}

		// turn
		if area[nexti][nextj] == "#" {
			area[gi][gj] = turns[arrow]
		} else {
			area[nexti][nextj] = arrow
			area[gi][gj] = "."
			gi = nexti
			gj = nextj
		}
	}

	return false
}

func inRange(i, j, m, n int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}

func printMatrix[E any](m [][]E) {
	for _, v := range m {
		fmt.Println(v)
	}
}

func copyVisited(visited [][][]string) [][][]string {
	m := len(visited)
	n := len(visited[0])

	cp := make([][][]string, m)
	for i := 0; i < m; i++ {
		cp[i] = make([][]string, n)
		for j := 0; j < n; j++ {
			cp[i][j] = make([]string, 4)
			copy(cp[i][j], visited[i][j])
		}
	}

	return cp
}
