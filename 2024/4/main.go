package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	word       = "XMAS"
	directions = [8][3][2]int{
		{{-1, 0}, {-2, 0}, {-3, 0}},    // top
		{{1, 0}, {2, 0}, {3, 0}},       // bot
		{{0, -1}, {0, -2}, {0, -3}},    // left
		{{0, 1}, {0, 2}, {0, 3}},       // right
		{{-1, -1}, {-2, -2}, {-3, -3}}, // top-left diag
		{{1, 1}, {2, 2}, {3, 3}},       // bottom-right diag
		{{-1, 1}, {-2, 2}, {-3, 3}},    // top-right diag
		{{1, -1}, {2, -2}, {3, -3}},    // bottom-left diag
	}
)

func main() {
	f, err := os.Open("./2024/4/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	puzzle := strings.Split(string(input), "\n")
	fmt.Println("part 1", part1(puzzle))
	fmt.Println("part 2", part2(puzzle))
}

func part1(puzzle []string) int {
	sum := 0

	m := len(puzzle)
	n := len(puzzle[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if puzzle[i][j] == 'X' {
				for _, d := range directions {

					isValid := true
					for chIdx, dxdy := range d {
						ii := i + dxdy[0]
						jj := j + dxdy[1]

						if !validRange(m, n, ii, jj) || word[chIdx+1] != puzzle[ii][jj] {
							isValid = false
							break
						}

					}

					if isValid {
						sum += 1
					}
				}
			}
		}
	}

	return sum
}

func part2(puzzle []string) int {

	validDiagChar := func(ch byte) bool {
		return ch == 'M' || ch == 'S'
	}

	sum := 0

	m := len(puzzle)
	n := len(puzzle[0])
	// no point to check edges
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if puzzle[i][j] == 'A' {
				topLeft := puzzle[i-1][j-1]
				bottomRight := puzzle[i+1][j+1]
				topRight := puzzle[i-1][j+1]
				bottomLeft := puzzle[i+1][j-1]

				if validDiagChar(topLeft) && validDiagChar(bottomRight) && topLeft != bottomRight &&
					validDiagChar(topRight) && validDiagChar(bottomLeft) && topRight != bottomLeft {
					sum += 1
				}
			}
		}
	}

	return sum
}

func validRange(m, n, i, j int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}
