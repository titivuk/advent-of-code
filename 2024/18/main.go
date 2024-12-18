package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/titivuk/advent-of-code/utils"
)

func main() {
	ex1 := `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`
	testPart1(ex1, "ex1", 12, 6+1, 6+1)
	testPart2(ex1, "ex1", 6+1, 6+1)

	f, err := os.Open("./2024/18/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	_ = input
	testPart1(string(input), "input", 1024, 70+1, 70+1)
	testPart2(string(input), "input", 70+1, 70+1)
}

func testPart2(input, name string, m, n int) {
	fmt.Printf("%s p2: %s\n", name, part2(input, m, n))
}

func testPart1(input, name string, b, m, n int) {
	fmt.Printf("%s p1: %d\n", name, part1(input, b, m, n))
}

func part2(input string, m, n int) string {
	lines := strings.Split(input, "\n")
	l := 1
	r := len(lines)
	for l <= r {
		mid := (l + r) / 2

		steps := part1(input, mid, m, n)
		if steps == -1 {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return lines[l-1]
}

func part1(input string, b, m, n int) int {
	grid := make([][]byte, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			grid[i][j] = '.'
		}
	}

	lines := strings.Split(input, "\n")
	for k := 0; k < b; k++ {
		l := lines[k]
		parts := strings.SplitN(l, ",", 2)
		i, _ := strconv.Atoi(parts[1])
		j, _ := strconv.Atoi(parts[0])

		grid[i][j] = '#'
	}

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	steps := 0
	// starting point
	i, j := 0, 0
	visited[i][j] = true
	curVertices := []struct{ i, j int }{
		{i, j},
	}
	nextVertices := []struct{ i, j int }{}
	for len(curVertices) > 0 {
		for _, cur := range curVertices {

			if cur.i == m-1 && cur.j == n-1 {
				return steps
			}

			for _, d := range utils.StraightDirections {
				next := struct{ i, j int }{i: cur.i + d[0], j: cur.j + d[1]}

				if utils.InRange(m, n, next.i, next.j) && grid[next.i][next.j] != '#' && !visited[next.i][next.j] {
					visited[next.i][next.j] = true
					nextVertices = append(nextVertices, next)
				}
			}
		}

		steps += 1

		curVertices = nextVertices
		nextVertices = nil
	}

	return -1
}
