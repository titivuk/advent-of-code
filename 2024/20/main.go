package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/titivuk/advent-of-code/utils"
)

type cheat struct {
	start utils.Point
	end   utils.Point
}

func main() {
	ex1 := `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`
	testPart12(ex1, 2, "ex1")

	f, err := os.Open("./2024/20/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	testPart12(string(input), 100, "input")
}

func testPart12(input string, minDelta int, name string) {
	grid := parseInput(input)
	count := part1(grid, minDelta)
	fmt.Printf("%s p1: %d\n", name, count)
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

func part1(grid [][]byte, minDelta int) int {
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

	// also let's put how many steps it took to reach every cell
	// visited[start] = 0
	// visited[start + 1 step] = 1
	// ...
	// visited[end] = N  
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
		for j := 0; j < n; j++ {
			visited[i][j] = -1
		}
	}

	// BFS to traverse the grid and build the path
	// (since path is only on we could use DFS as well)
	curVertices := []utils.Point{
		start,
	}
	visited[start.Y][start.X] = 0
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

				if utils.InRange(m, n, next.Y, next.X) && grid[next.Y][next.X] != '#' && visited[next.Y][next.X] == -1 {
					visited[next.Y][next.X] = visited[cur.Y][cur.X] + 1
					nextVertices = append(nextVertices, next)
				}
			}
		}

		curVertices = nextVertices
		nextVertices = nil
	}

	// after we have built the path
	// lets try on every path step to go through the neigbor wall
	// if we cheat and the next cell after the wall has bigger visited[i][j] value
	// it means that we moved from current cell to the next cell and reduced our path by visited[nextCell] - visited[currentCell] - 2 (2 picoseconds to cheat)
	cheats := map[cheat]struct{}{}
	for _, p := range path {
		for _, sd := range utils.StraightDirections {
			next := utils.Point{X: p.X + sd[1], Y: p.Y + sd[0]}

			start := p
			// cheat path looks like this: start -> wall -> end
			// so next is middle cell which must be a wall
			if utils.InRange(m, n, next.Y, next.X) && grid[next.Y][next.X] == '#' {
				// try to move to every possible direction
				// it must be non-wall cell and it's path value must be bigger to actually win some time (minDelta)
				for _, ed := range utils.StraightDirections {
					end := utils.Point{X: next.X + ed[1], Y: next.Y + ed[0]}

					if utils.InRange(m, n, end.Y, end.X) && grid[end.Y][end.X] != '#' && visited[end.Y][end.X]-visited[p.Y][p.X] > minDelta {
						// remember every unique pair of start and end cells
						cheats[cheat{start: start, end: end}] = struct{}{}
					}
				}
			}
		}
	}

	return len(cheats)
}
