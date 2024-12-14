package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
	"github.com/titivuk/advent-of-code/utils"
)

type robot struct {
	p utils.Point
	v utils.Point
}

func (r *robot) AdvancePosition(m, n, t int) {
	r.p = utils.Point{
		X: r.calculateCoordinate(r.p.X, n, r.v.X, t),
		Y: r.calculateCoordinate(r.p.Y, m, r.v.Y, t),
	}
}

func (r robot) calculateCoordinate(coord, length, speed, t int) int {
	S := speed * t
	offset := S % length
	newCoord := coord + offset
	if newCoord < 0 {
		// the same as length - |coord|
		newCoord = length + newCoord
	} else if newCoord >= length {
		newCoord %= length
	}

	return newCoord
}

func main() {
	f, err := os.Open("./2024/14/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	testPart1(string(input), "input", 103, 101, 100)
	fmt.Printf("p1: %d\n", part2(parseInput(string(input)), 103, 101))
}

func testPart1(input, name string, m, n, t int) {
	fmt.Printf("%s p1: %d\n", name, part1(parseInput(input), m, n, t))
}

func parseInput(input string) []robot {
	lines := strings.Split(input, "\n")
	robots := make([]robot, len(lines))

	for lIdx, row := range lines {
		i := 0

		// skip "p="
		i += 2

		// parse px
		l := i
		for i < len(row) && row[i] != ',' {
			i += 1
		}
		px, err := strconv.Atoi(row[l:i])
		if err != nil {
			panic(err)
		}

		// skip ','
		i += 1

		// parse py
		l = i
		for i < len(row) && utils.IsDigit(row[i]) {
			i += 1
		}
		py, err := strconv.Atoi(row[l:i])
		if err != nil {
			panic(err)
		}

		// skip " v="
		i += 3

		// parse vx
		l = i
		for i < len(row) && row[i] != ',' {
			i += 1
		}
		vx, err := strconv.Atoi(row[l:i])
		if err != nil {
			panic(err)
		}

		// skip ','
		i += 1

		// parse vy
		l = i
		for i < len(row) {
			i += 1
		}
		vy, err := strconv.Atoi(row[l:i])
		if err != nil {
			panic(err)
		}

		r := robot{
			p: utils.Point{X: px, Y: py},
			v: utils.Point{X: vx, Y: vy},
		}

		robots[lIdx] = r
	}

	return robots
}

func part1(robots []robot, m, n, t int) int {
	xmid := n / 2
	ymid := m / 2

	// q2 q1
	// q3 q4
	var q1, q2, q3, q4 int
	for _, sr := range robots {
		sr.AdvancePosition(m, n, t)
		nr := sr

		// q1
		if nr.p.X > xmid && nr.p.Y < ymid {
			q1 += 1
		}
		// q2
		if nr.p.X < xmid && nr.p.Y < ymid {
			q2 += 1
		}
		// q3
		if nr.p.X < xmid && nr.p.Y > ymid {
			q3 += 1
		}
		// q4
		if nr.p.X > xmid && nr.p.Y > ymid {
			q4 += 1
		}
	}

	safetyFactor := q1 * q2 * q3 * q4

	return safetyFactor
}

func part2(robots []robot, m, n int) int {
	// thanks to https://www.reddit.com/r/adventofcode/comments/1hdvhvu/comment/m1zws1g/

	// https://en.wikipedia.org/wiki/Chinese_remainder_theorem
	// N = r1(mod a1)
	// N = r2(mod a2)
	// =>
	// N = r1 + a*a1
	// N = r2 + b*a2
	// a1 and a2 are co-prime and a1=n and a2=m
	// r1 and r2 are the most clustered x and y accordingly

	// find min variance
	x := 0
	var xVal float64 = math.MaxFloat64
	y := 0
	var yVal float64 = math.MaxFloat64
	for i := 0; i < m*n; i++ {
		xs := make([]float64, len(robots))
		ys := make([]float64, len(robots))

		for j := 0; j < len(robots); j++ {
			xs[j] = float64(robots[j].p.X)
			ys[j] = float64(robots[j].p.Y)

			robots[j].AdvancePosition(m, n, 1)
		}

		if xv, err := stats.Variance(xs); err != nil {
			panic(err)
		} else if xv < xVal {
			x = i
			xVal = xv
		}

		if yv, err := stats.Variance(ys); err != nil {
			panic(err)
		} else if yv < yVal {
			y = i
			yVal = yv
		}

	}

	// N = x + a*n
	// N = y + b*m
	// need to solve the system of the equations above
	// I did not find modulos operations in go STD
	// so here brutefoce solution (a and b are integers)
	N := 0
	a := 0
	b := 0
	for {
		l := x + a*n
		r := y + b*m

		if l == r {
			N = l
			break
		}

		if l < r {
			a += 1
		} else {
			b += 1
		}
	}

	return N
}
