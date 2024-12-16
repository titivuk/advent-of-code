package main

import (
	"container/heap"
	"fmt"
	"io"
	"maps"
	"math"
	"os"
	"strings"

	"github.com/titivuk/advent-of-code/utils"
)

// combination of current position and moving direction
type State struct {
	pos       utils.Point
	direction utils.Point
}

type WeightedPoint struct {
	state  State
	weight int
	// remember path for every state
	path map[utils.Point]bool
}

type MinHeap []*WeightedPoint

func (pq MinHeap) Len() int {
	return len(pq)
}

func (pq MinHeap) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinHeap) Push(x any) {
	item := x.(*WeightedPoint)
	*pq = append(*pq, item)
}

func (pq *MinHeap) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	ex1 := `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`
	test(ex1, "ex1")

	ex2 := `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`
	test(ex2, "ex2")

	ex3 := `##########
#.......E#
#.##.#####
#..#.....#
##.#####.#
#S.......#
##########`
	test(ex3, "ex3")

	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	test(string(input), "input")
}

func test(input, name string) {
	score, tiles := part12(parseInput(input))
	fmt.Printf("%s: score - %d, tiles - %d\n", name, score, tiles)
}

func parseInput(input string) [][]byte {
	rows := strings.Split(input, "\n")

	matrix := make([][]byte, len(rows))
	for i := 0; i < len(rows); i++ {
		matrix[i] = []byte(rows[i])
	}

	return matrix
}

func part12(matrix [][]byte) (int, int) {
	m := len(matrix)
	n := len(matrix[0])

	// dijkstra implementation but vertex distance and visits are tracked for every direction (at most 4 directions for every cell)

	// distance for each vertex from every direction
	dist := make(map[State]int)

	// find start
	i, j := 0, 0
outer:
	for i = 0; i < m; i++ {
		for j = 0; j < n; j++ {
			if matrix[i][j] == 'S' {
				break outer
			}
		}
	}

	mh := &MinHeap{}
	heap.Init(mh)

	// add starting point
	state := State{
		pos:       utils.Point{X: j, Y: i},
		direction: utils.Point{X: 1, Y: 0},
	}
	heap.Push(mh, &WeightedPoint{
		path: map[utils.Point]bool{
			{X: j, Y: i}: true,
		},
		weight: 0,
		state:  state,
	})
	dist[state] = 0

	minScore := math.MaxInt
	p2 := map[utils.Point]bool{}
	for mh.Len() > 0 {
		p := heap.Pop(mh).(*WeightedPoint)

		// instead of visited we track via dist
		// if we've been in this vertex with given direction and current path is more costly
		// we can skip that branch
		if c, ok := dist[p.state]; ok && c < p.weight {
			continue
		}
		// set new min score
		dist[p.state] = p.weight

		// if 'E' found and path cost less then known score
		// new best found
		if matrix[p.state.pos.Y][p.state.pos.X] == 'E' && p.weight <= minScore {
			minScore = p.weight
			maps.Copy(p2, p.path)
		}

		dx, dy := p.state.direction.X, p.state.direction.Y
		dirs := [][3]int{
			{dy, dx, 1},      // forward
			{-dx, -dy, 1001}, // left
			{dx, dy, 1001},   // right
		}

		for _, d := range dirs {
			nx, ny := p.state.pos.X+d[1], p.state.pos.Y+d[0]
			weight := d[2]

			// hit the wall
			if matrix[ny][nx] == '#' {
				continue
			}

			nstate := State{
				pos:       utils.Point{X: nx, Y: ny},
				direction: utils.Point{X: d[1], Y: d[0]},
			}
			npath := maps.Clone(p.path)
			npath[nstate.pos] = true
			heap.Push(mh, &WeightedPoint{
				path:   npath,
				weight: dist[p.state] + weight,
				state:  nstate,
			})
		}
	}

	return minScore, len(p2)
}
