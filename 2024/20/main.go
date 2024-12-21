package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"

	"github.com/titivuk/advent-of-code/utils"
)

func main() {
	f, err := os.Open("./2024/20/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	testPart1(string(input), 100, "input")
	testPart2(string(input), 100, "input")
}

func testPart1(input string, minDelta int, name string) {
	grid := parseInput(input)
	count := part2(grid, 2, minDelta)
	fmt.Printf("%s p1: %d\n", name, count)
}

func testPart2(input string, minDelta int, name string) {
	grid := parseInput(input)
	count := part2(grid, 20, minDelta)
	fmt.Printf("%s p2: %d\n", name, count)
}

func parseInput(input string) [][]byte {
	lines := strings.Split(input, "\n")

	grid := make([][]byte, len(lines))
	for i := 0; i < len(lines); i++ {
		grid[i] = make([]byte, len(lines[i]))
		for j := 0; j < len(lines[i]); j++ {
			grid[i][j] = lines[i][j]
		}
	}

	return grid
}

func part2(grid [][]byte, cheatTime, minDelta int) int {
	m, n := len(grid), len(grid[0])

	// find start
	start := utils.Point{}
	finish := utils.Point{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'S' {
				start.X = j
				start.Y = i
			} else if grid[i][j] == 'E' {
				finish.X = j
				finish.Y = i
			}
		}
	}

	// since there is only one valid path
	// lets traverse the grid and build this path
	path := make([]utils.Point, 0)

	// BFS to traverse the grid and build the path
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	curVertices := []utils.Point{
		start,
	}
	visited[start.Y][start.X] = true
	nextVertices := []utils.Point{}
outer:
	for len(curVertices) > 0 {
		for _, cur := range curVertices {
			path = append(path, cur)
			if cur == finish {
				break outer
			}

			for _, d := range utils.StraightDirections {
				next := utils.Point{X: cur.X + d[1], Y: cur.Y + d[0]}

				if utils.InRange(m, n, next.Y, next.X) && grid[next.Y][next.X] != '#' && !visited[next.Y][next.X] {
					visited[next.Y][next.X] = true
					nextVertices = append(nextVertices, next)
				}
			}
		}

		curVertices = nextVertices
		nextVertices = nil
	}

	// after long debugging I found out that cheating can start from any cell
	// (no need step on the wall to start cheating path)
	// so, for every cheat start `s` check cheat end `e`
	// formula is simplified from
	// path_to_cheat_start + distance_between_start_and_end + remaining_path <= initial_path - min_delta
	// path_to_cheat_start = sPath
	// distance_between_start_and_end - distance between `s` and `e` points. Since we are cheating, we can go through any cell
	// remaining_path = initial_path - path_to_cheat_end = initial_path - ePath
	count := 0
	for sPath, s := range path {
		for ePath, e := range path {
			se_path := pointDistance(s, e)

			if se_path <= cheatTime && ePath-sPath-se_path >= minDelta {
				count += 1
			}
		}
	}

	return count
}

func pointDistance(a, b utils.Point) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}
