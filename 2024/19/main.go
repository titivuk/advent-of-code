package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	ex1 := `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`
	testPart12(ex1, "ex1")

	f, err := os.Open("./2024/19/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	testPart12(string(input), "input")
}

func testPart12(input, name string) {
	designs, patterns := parseInput(input)
	possibleDesigns, possibleCombinations := part12(designs, patterns)
	fmt.Printf("%s p1: %d; p2: %d \n", name, possibleDesigns, possibleCombinations)
}

func parseInput(input string) ([]string, map[string]struct{}) {
	patterns := map[string]struct{}{}
	designs := make([]string, 0)

	lines := strings.Split(input, "\n")
	for _, p := range strings.Split(lines[0], ", ") {
		patterns[p] = struct{}{}
	}

	for i := 2; i < len(lines); i++ {
		designs = append(designs, lines[i])
	}

	return designs, patterns
}

func part12(designs []string, patterns map[string]struct{}) (int, int) {
	seen := map[string]int{}

	possibleDesigns := 0
	possibleCombinations := 0
	for _, d := range designs {
		if c := designPossible(d, patterns, 0, seen); c > 0 {
			possibleDesigns += 1
			possibleCombinations += c
		}
	}

	return possibleDesigns, possibleCombinations
}

func designPossible(design string, patterns map[string]struct{}, i int, seen map[string]int) int {
	if i >= len(design) {
		return 1
	}

	if seen[design[i:]] > 0 {
		return seen[design[i:]]
	}

	count := 0
	for p := range patterns {
		inext := i + len(p)
		if inext > len(design) {
			continue
		}

		if design[i:inext] == p {
			count += designPossible(design, patterns, inext, seen)
		}
	}

	seen[design[i:]] += count

	return count
}
