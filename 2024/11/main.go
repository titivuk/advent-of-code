package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./2024/11/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1", part12(parseInput(string(input)), 25))
	fmt.Println("part2", part12(parseInput(string(input)), 75))
}

func parseInput(stonesStr string) map[int]int {
	stones := make(map[int]int)

	for _, ss := range strings.Split(stonesStr, " ") {
		stone, err := strconv.Atoi(ss)
		if err != nil {
			panic(err)
		}

		stones[stone] += 1
	}

	return stones
}

func part12(stones map[int]int, blinks int) int {
	for i := 0; i < blinks; i++ {
		stonesAfterBlink := make(map[int]int)

		for s, count := range stones {
			if s == 0 {
				stonesAfterBlink[1] += count
			} else if countDigits(s)%2 == 0 {
				a, b := splitInt(s)

				stonesAfterBlink[a] += count
				stonesAfterBlink[b] += count
			} else {
				stonesAfterBlink[s*2024] += count
			}
		}

		stones = stonesAfterBlink
	}

	sum := 0
	for _, count := range stones {
		sum += count
	}

	return sum
}

func countDigits(val int) int {
	return int(math.Log10(float64(val))) + 1
}

func splitInt(val int) (int, int) {
	digits := countDigits(val)
	divider := int(math.Pow10(digits / 2))

	a := val / divider
	b := val % divider

	return a, b
}
