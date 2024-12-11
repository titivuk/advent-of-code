package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	f, err := os.Open("./2024/8/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, _ := io.ReadAll(f)

	fmt.Println("part1", part1(parseInput(string(input))))
	fmt.Println("part2", part2(parseInput(string(input))))
}

func parseInput(inputStr string) [][]byte {
	rows := strings.Split(inputStr, "\n")

	input := make([][]byte, len(rows))
	for i, r := range rows {
		row := make([]byte, len(r))
		for j := 0; j < len(r); j++ {
			row[j] = r[j]
		}
		input[i] = row
	}

	return input
}

func part1(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	result := 0

	antennas := make(map[byte][]point)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != '.' {
				antennas[matrix[i][j]] = append(antennas[matrix[i][j]], point{i, j})
			}
		}
	}

	for _, points := range antennas {
		for i := 0; i < len(points)-1; i++ {
			a1 := points[i]
			for j := i + 1; j < len(points); j++ {
				a2 := points[j]

				dx := a2.x - a1.x
				dy := a2.y - a1.y

				antinode := point{a2.x + dx, a2.y + dy}
				if inRange(m, n, antinode.x, antinode.y) && matrix[antinode.x][antinode.y] != '#' {
					matrix[antinode.x][antinode.y] = '#'
					result += 1
				}

				antinode = point{a1.x - dx, a1.y - dy}
				if inRange(m, n, antinode.x, antinode.y) && matrix[antinode.x][antinode.y] != '#' {
					matrix[antinode.x][antinode.y] = '#'
					result += 1
				}
			}
		}

	}

	return result
}

func part2(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	result := 0

	antennas := make(map[byte][]point)

	// x AND y ARE REVERSED
	// i.e. x = i, y = j
	// too lazy to change

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != '.' {
				antennas[matrix[i][j]] = append(antennas[matrix[i][j]], point{i, j})
			}
		}
	}

	var a1, a2, antinode point
	for _, points := range antennas {
		for i := 0; i < len(points)-1; i++ {
			a1 = points[i]

			for j := i + 1; j < len(points); j++ {
				a2 = points[j]

				dx := a2.x - a1.x
				dy := a2.y - a1.y

				antinode.x = a2.x
				antinode.y = a2.y
				for inRange(m, n, antinode.x, antinode.y) {
					if matrix[antinode.x][antinode.y] != '#' {
						matrix[antinode.x][antinode.y] = '#'
						result += 1
					}

					antinode.x += dx
					antinode.y += dy
				}

				antinode.x = a1.x
				antinode.y = a1.y
				for inRange(m, n, antinode.x, antinode.y) {
					if matrix[antinode.x][antinode.y] != '#' {
						matrix[antinode.x][antinode.y] = '#'
						result += 1
					}

					antinode.x -= dx
					antinode.y -= dy
				}
			}
		}
	}

	return result
}

func inRange(m, n, i, j int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}
