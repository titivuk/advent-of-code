package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	directions = [4][2]int{
		{-1, 0}, // top
		{0, 1},  // right
		{1, 0},  // bot
		{0, -1}, // left
	}
)

type point struct {
	x, y int
}

func main() {
	f, err := os.Open("./2024/10/input.txt")
	if err != nil {
		panic(err)
	}

	var hikeMap [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		row := make([]int, len(line))

		for i := 0; i < len(line); i++ {
			ch := line[i]
			row[i], _ = strconv.Atoi(string(ch))
		}

		hikeMap = append(hikeMap, row)
	}

	fmt.Println(part12(hikeMap))
}

func part12(hikeMap [][]int) (int, int) {
	m := len(hikeMap)
	n := len(hikeMap[0])

	sumScore := 0
	sumRating := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if hikeMap[i][j] == 0 {
				visited := make(map[point]bool)
				score, rating := checkTrail(hikeMap, visited, i, j)

				sumScore += score
				sumRating += rating
			}
		}
	}

	return sumScore, sumRating
}

func checkTrail(hikeMap [][]int, visited map[point]bool, i, j int) (int, int) {
	if hikeMap[i][j] == 9 {
		rating := 1
		score := 0

		p := point{i, j}
		if !visited[p] {
			visited[p] = true
			score = 1
		}

		return score, rating
	}

	m := len(hikeMap)
	n := len(hikeMap[0])

	sumScore := 0
	sumRating := 0
	for _, d := range directions {
		di := i + d[0]
		dj := j + d[1]

		if inRange(di, dj, m, n) && hikeMap[i][j]+1 == hikeMap[di][dj] {
			score, rating := checkTrail(hikeMap, visited, di, dj)
			sumScore += score
			sumRating += rating
		}
	}

	return sumScore, sumRating
}

func inRange(i, j, m, n int) bool {
	return 0 <= i && i < m && 0 <= j && j < n
}
