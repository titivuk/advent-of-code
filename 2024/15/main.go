package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

var (
	m2d = map[byte][2]int{
		'^': {-1, 0},
		'>': {0, 1},
		'v': {1, 0},
		'<': {0, -1},
	}
)

func main() {
	f, err := os.Open("./2024/15/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	testPart1(string(input), "input")
	testPart2(string(input), "input")
}

func testPart1(input string, name string) {
	wh, moves := parseInput(input)
	fmt.Printf("%s p1: %d\n", name, part1(wh, moves))
}

func testPart2(input string, name string) {
	wh, moves := parseInput(input)
	fmt.Printf("%s p2: %d\n", name, part2(wh, moves))
}

func parseInput(input string) ([][]byte, string) {
	lines := strings.Split(input, "\n")
	warehouse := make([][]byte, 0)
	i := 0
	for lines[i] != "" {
		warehouse = append(warehouse, []byte(lines[i]))
		i += 1
	}

	moves := strings.Join(lines[i+1:], "")

	return warehouse, moves
}

func part1(warehouse [][]byte, moves string) int {

	m := len(warehouse)
	n := len(warehouse[0])

	i, j := 0, 0
out:
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if warehouse[i][j] == '@' {
				break out
			}
		}
	}

	for mi := 0; mi < len(moves); mi++ {
		// printMatrix(warehouse)

		move := moves[mi]
		dir := m2d[move]

		iNext := i + dir[0]
		jNext := j + dir[1]

		if warehouse[iNext][jNext] == '.' {
			warehouse[i][j] = '.'
			warehouse[iNext][jNext] = '@'
			i, j = iNext, jNext
		} else if warehouse[iNext][jNext] == 'O' {
			// try to move
			pi := iNext
			pj := jNext
			for warehouse[pi][pj] != '.' && warehouse[pi][pj] != '#' {
				pi += dir[0]
				pj += dir[1]
			}

			if warehouse[pi][pj] == '.' {
				for pi != iNext || pj != jNext {
					warehouse[pi][pj] = warehouse[pi-dir[0]][pj-dir[1]]
					pi -= dir[0]
					pj -= dir[1]
				}

				warehouse[i][j] = '.'
				warehouse[iNext][jNext] = '@'
				i, j = iNext, jNext
			}
		}
	}

	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if warehouse[i][j] == 'O' {
				sum += 100*i + j
			}
		}
	}

	return sum
}

func part2(warehouse [][]byte, moves string) int {
	m := len(warehouse)
	n := len(warehouse[0])

	newWh := make([][]byte, len(warehouse))
	for i := 0; i < m; i++ {
		newWh[i] = make([]byte, 0)
		for j := 0; j < n; j++ {
			if warehouse[i][j] == '#' {
				newWh[i] = append(newWh[i], '#', '#')
			} else if warehouse[i][j] == 'O' {
				newWh[i] = append(newWh[i], '[', ']')
			} else if warehouse[i][j] == '.' {
				newWh[i] = append(newWh[i], '.', '.')
			} else if warehouse[i][j] == '@' {
				newWh[i] = append(newWh[i], '@', '.')
			}
		}
	}
	warehouse = newWh
	m = len(warehouse)
	n = len(warehouse[0])

	i, j := 0, 0
out:
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if warehouse[i][j] == '@' {
				break out
			}
		}
	}

	for mi := 0; mi < len(moves); mi++ {
		move := moves[mi]
		dir := m2d[move]

		iNext := i + dir[0]
		jNext := j + dir[1]

		if warehouse[iNext][jNext] == '.' {
			warehouse[i][j] = '.'
			warehouse[iNext][jNext] = '@'
			i, j = iNext, jNext
		} else if warehouse[iNext][jNext] == '[' || warehouse[iNext][jNext] == ']' {
			// try to push
			pi := iNext
			pj := jNext

			if move == '<' || move == '>' {
				for warehouse[pi][pj] != '.' && warehouse[pi][pj] != '#' {
					pi += dir[0]
					pj += dir[1]
				}

				if warehouse[pi][pj] == '.' {
					for pi != iNext || pj != jNext {
						warehouse[pi][pj] = warehouse[pi-dir[0]][pj-dir[1]]
						pi -= dir[0]
						pj -= dir[1]
					}

					warehouse[i][j] = '.'
					warehouse[iNext][jNext] = '@'
					i, j = iNext, jNext
				}
			} else {
				// column that must be free to go there
				cols := make([]int, 0)
				cols = append(cols, jNext)

				if tryPush(warehouse, iNext, dir[0], cols) {
					warehouse[i][j] = '.'
					warehouse[iNext][jNext] = '@'
					i, j = iNext, jNext
				}
			}
		}

		// time.Sleep(50 * time.Millisecond)
		// fmt.Print("\033[H\033[2J")
		// for i := 0; i < m; i++ {
		// 	fmt.Println(string(warehouse[i]))
		// }
	}

	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if warehouse[i][j] == '[' {
				sum += 100*i + j
			}
		}
	}

	return sum
}

func tryPush(warehouse [][]byte, row, rowOffset int, cols []int) bool {
	// wall
	if row < 1 {
		return false
	}

	// cols - list of columns that must be free on the next row to push current row boxes
	// nextCols - list of columns that must be free on the next-next row to push next row boxes
	//   1 2 3 4
	// 1 . . . .    cols = [1,2,3,4], nextCols = []
	// 2 [ ] [ ]	cols = [1,2,3,4], nextCols = [1,2,3,4]
	// 3 . [ ] .	cols = [3,4], nextCols = [1,2,3,4]
	// 4 . # . .
	// from bot to top
	nextCols := make([]int, 0)
	for j := 0; j < len(cols); j++ {
		col := cols[j]

		// boxes cannot be moved
		if warehouse[row][col] == '#' {
			return false
		}

		// given box can be moved
		if warehouse[row][col] == '.' {
			continue
		}

		// avoid duplicates
		if slices.Index(nextCols, col) == -1 {
			nextCols = append(nextCols, col)
		}

		// do not forget to append the 2nd part of the box
		if warehouse[row][col] == ']' {
			if slices.Index(nextCols, col-1) == -1 {
				nextCols = append(nextCols, col-1)
			}
		}

		// do not forget to append the 2nd part of the box
		if warehouse[row][col] == '[' {
			if slices.Index(nextCols, col+1) == -1 {
				nextCols = append(nextCols, col+1)
			}
		}
	}

	// if there is something has to be moved to free space
	// recursively check the next row
	if len(nextCols) > 0 {
		possible := tryPush(warehouse, row+rowOffset, rowOffset, nextCols)

		// post processing when recursion going back
		// if the next row has space
		// push boxes from the current row
		// boxes from the prev row will be pushed to the current row on the next iteration of recursion
		if possible {
			for j := 0; j < len(nextCols); j++ {
				col := nextCols[j]

				warehouse[row+rowOffset][col] = warehouse[row][col]
				warehouse[row][col] = '.'
			}

		}

		return possible
	}

	// len(nextCols) == 0 => all cols are '.' so we can push
	return true
}
