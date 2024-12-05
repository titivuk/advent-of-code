package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./2024/5/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rules, updates := parseInput(f)

	fmt.Println("part1", part1(rules, updates))
	fmt.Println("part2", part2(rules, updates))
}

func parseInput(r io.Reader) (map[int][]int, [][]int) {
	scanner := bufio.NewScanner(r)

	rulesPerPage := make(map[int][]int)
	var l string
	for scanner.Scan() {
		l = scanner.Text()

		if l == "" {
			break
		}

		pages := strings.Split(l, "|")
		lp, err := strconv.Atoi(pages[0])
		if err != nil {
			panic(err)
		}
		rp, err := strconv.Atoi(pages[1])
		if err != nil {
			panic(err)
		}

		rulesPerPage[lp] = append(rulesPerPage[lp], rp)
	}

	updates := make([][]int, 0)
	for scanner.Scan() {
		l = scanner.Text()

		pages := strings.Split(l, ",")
		updatePages := make([]int, len(pages))
		for i, p := range pages {
			pv, err := strconv.Atoi(p)
			if err != nil {
				panic(err)
			}

			updatePages[i] = pv
		}
		updates = append(updates, updatePages)
	}

	return rulesPerPage, updates
}

func part1(rules map[int][]int, updates [][]int) int {
	sum := 0

	for _, u := range updates {
		if validUpdate(u, rules) {
			sum += u[len(u)/2]
		}
	}

	return sum
}

func part2(rules map[int][]int, updates [][]int) int {
	sum := 0

	for _, u := range updates {
		if !validUpdate(u, rules) {
			orderedUpdate := sortUpdate(u, rules)
			sum += orderedUpdate[len(orderedUpdate)/2]
		}
	}

	return sum
}

func validUpdate(update []int, rules map[int][]int) bool {
	seen := make(map[int]bool)

	for _, x := range update {
		yPages := rules[x]

		for _, y := range yPages {
			// if any page `y` that has to be after page `x` seen
			// order is violated
			if seen[y] {
				return false
			}
		}

		seen[x] = true
	}

	return true
}

func sortUpdate(update []int, rules map[int][]int) []int {
	slices.SortFunc(update, func(a, b int) int {
		if r, ok := rules[b]; ok {
			if slices.Index(r, a) != -1 {
				// a should come before b
				return -1
			}
		}

		// b should come before a
		return 1
	})

	return update
}
