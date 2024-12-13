package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/titivuk/advent-of-code/utils"
)

func main() {
	// 80
	example1 := `AAAA
BBCD
BBCC
EEEC`
	testPart2(example1, "ex1")

	// 436
	example2 := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	testPart2(example2, "ex2")

	// 236
	example3 := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	testPart2(example3, "ex3")

	// 368
	example4 := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
	testPart2(example4, "ex4")

	f, err := os.Open("./2024/12/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, _ := io.ReadAll(f)

	testPart2(string(input), "input")
}

func testPart2(input string, name string) {
	p1, p2 := part12(parseInput(input))
	fmt.Printf("%s - p1: %d; p2: %d\n", name, p1, p2)
}

func parseInput(garderStr string) [][]string {
	rows := strings.Split(garderStr, "\n")
	garden := make([][]string, len(rows))
	for i, r := range rows {
		garden[i] = strings.Split(r, "")
	}

	return garden
}

func part12(garden [][]string) (int, int) {
	m := len(garden)
	n := len(garden[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	totalPrice1 := 0
	totalPrice2 := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				visited[i][j] = true
				area, perimeter, angles := traverseRegion(garden, visited, i, j)

				totalPrice1 += area * perimeter
				totalPrice2 += area * angles
			}
		}
	}

	return totalPrice1, totalPrice2
}

func traverseRegion(garden [][]string, visited [][]bool, i, j int) (int, int, int) {
	m := len(garden)
	n := len(garden[0])

	angles := 0

	// no refactoring attempts were made

	// utils.StraightDirections order change
	// might break the logic
	up := utils.StraightDirections[0]
	right := utils.StraightDirections[1]
	down := utils.StraightDirections[2]
	left := utils.StraightDirections[3]

	ui := i + up[0]
	uj := j + up[1]

	ri := i + right[0]
	rj := j + right[1]

	di := i + down[0]
	dj := j + down[1]

	li := i + left[0]
	lj := j + left[1]

	// up and right 90
	if (!utils.InRange(m, n, ui, uj) || garden[i][j] != garden[ui][uj]) &&
		(!utils.InRange(m, n, ri, rj) || garden[i][j] != garden[ri][rj]) {
		angles += 1
	}
	// up and right 270
	if (utils.InRange(m, n, ui, uj) && garden[i][j] == garden[ui][uj]) &&
		(utils.InRange(m, n, ri, rj) && garden[i][j] == garden[ri][rj]) &&
		(!utils.InRange(m, n, i-1, j+1) || garden[i][j] != garden[i-1][j+1]) {
		angles += 1
	}

	// right and down 90
	if (!utils.InRange(m, n, ri, rj) || garden[i][j] != garden[ri][rj]) &&
		(!utils.InRange(m, n, di, dj) || garden[i][j] != garden[di][dj]) {
		angles += 1
	}
	// right and down 270
	if (utils.InRange(m, n, ri, rj) && garden[i][j] == garden[ri][rj]) &&
		(utils.InRange(m, n, di, dj) && garden[i][j] == garden[di][dj]) &&
		(!utils.InRange(m, n, i+1, j+1) || garden[i][j] != garden[i+1][j+1]) {
		angles += 1
	}

	// down and left 90
	if (!utils.InRange(m, n, di, dj) || garden[i][j] != garden[di][dj]) &&
		(!utils.InRange(m, n, li, lj) || garden[i][j] != garden[li][lj]) {
		angles += 1
	}
	// down and left 270
	if (utils.InRange(m, n, di, dj) && garden[i][j] == garden[di][dj]) &&
		(utils.InRange(m, n, li, lj) && garden[i][j] == garden[li][lj]) &&
		(!utils.InRange(m, n, i+1, j-1) || garden[i][j] != garden[i+1][j-1]) {
		angles += 1
	}

	// left and up 90
	if (!utils.InRange(m, n, li, lj) || garden[i][j] != garden[li][lj]) &&
		(!utils.InRange(m, n, ui, uj) || garden[i][j] != garden[ui][uj]) {
		angles += 1
	}
	// left and up 270
	if (utils.InRange(m, n, li, lj) && garden[i][j] == garden[li][lj]) &&
		(utils.InRange(m, n, ui, uj) && garden[i][j] == garden[ui][uj]) &&
		(!utils.InRange(m, n, i-1, j-1) || garden[i][j] != garden[i-1][j-1]) {
		angles += 1
	}

	area := 1
	perimeter := 4
	for _, d := range utils.StraightDirections {
		nexti := i + d[0]
		nextj := j + d[1]
		if utils.InRange(m, n, nexti, nextj) && garden[i][j] == garden[nexti][nextj] {
			perimeter -= 1

			if !visited[nexti][nextj] {
				visited[nexti][nextj] = true
				a, p, ia := traverseRegion(garden, visited, nexti, nextj)
				area += a
				perimeter += p
				angles += ia
			}
		}
	}

	return area, perimeter, angles
}
