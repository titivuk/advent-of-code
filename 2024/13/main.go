package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/titivuk/advent-of-code/utils"
)

type behavior struct {
	a, b  utils.Point
	prize utils.Point
}

func main() {
	f, err := os.Open("2024/13/input.txt")
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
	fmt.Printf("part 1: %s - %d\n", name, part12(parseInput(input, false)))
}

func testPart2(input string, name string) {
	fmt.Printf("part 2: %s - %d\n", name, part12(parseInput(input, true)))
}

func part12(behaviors []behavior) int {
	totalTokens := 0
	for _, bh := range behaviors {
		// solve a system of 2 equations
		// ax + by = X
		// cx + dy = Y
		a := float64(bh.a.X)
		b := float64(bh.b.X)
		X := float64(bh.prize.X)

		c := float64(bh.a.Y)
		d := float64(bh.b.Y)
		Y := float64(bh.prize.Y)

		x := (d*X - b*Y) / (a*d - b*c)
		y := (X - a*x) / b

		// only integer solutions are valid
		if math.Mod(x, 1.0) == 0 && math.Mod(y, 1.0) == 0 {
			totalTokens += int(x*3 + y)
		}
	}

	return totalTokens
}

func parseInput(input string, addToPrize bool) []behavior {
	behaviors := make([]behavior, 0)

	rows := strings.Split(input, "\n")
	i := 0
	for i < len(rows)-2 {
		aRow := rows[i]
		i += 1
		bRow := rows[i]
		i += 1
		prizeRow := rows[i]
		i += 2

		aButton := parseRow(aRow)
		bButton := parseRow(bRow)
		prize := parseRow(prizeRow)

		if addToPrize {
			prize.X += 10000000000000
			prize.Y += 10000000000000
		}
		behaviors = append(behaviors, behavior{a: aButton, b: bButton, prize: prize})
	}

	return behaviors
}

func parseRow(row string) utils.Point {
	i := 1
	for i < len(row) && !utils.IsDigit(row[i]) {
		i += 1
	}

	l := i
	for i < len(row) && utils.IsDigit(row[i]) {
		i += 1
	}

	x, err := strconv.Atoi(row[l:i])
	if err != nil {
		panic(err)
	}

	for i < len(row) && !utils.IsDigit(row[i]) {
		i += 1
	}

	l = i
	for i < len(row) && utils.IsDigit(row[i]) {
		i += 1
	}

	y, err := strconv.Atoi(row[l:i])
	if err != nil {
		panic(err)
	}

	return utils.Point{x, y}
}
