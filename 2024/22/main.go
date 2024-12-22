package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./2024/22/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	sum := part1(string(input))
	fmt.Println(sum)
}

func part1(input string) int {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, l := range lines {
		num, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}

		cur := num
		for i := 0; i < 2000; i += 1 {
			next := prune(mix(cur, cur*64))
			next = prune(mix(next, next/32))
			next = prune(mix(next, next*2048))

			cur = next
		}

		sum += cur
	}

	return sum
}

func mix(secret, num int) int {
	result := num ^ secret
	return result
}

func prune(secret int) int {
	result := secret % 16777216
	return result
}
