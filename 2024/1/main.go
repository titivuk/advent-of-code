package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	p1 := part1()
	fmt.Println("p1:", p1)

	p2 := part2()
	fmt.Println("p2:", p2)
}

func part1() int {
	left := make([]int, 0, 1000)
	right := make([]int, 0, 1000)

	f, err := os.Open("./2024/1/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Fields(line)

		lNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left = append(left, lNum)

		rNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		right = append(right, rNum)
	}

	slices.Sort(left)
	slices.Sort(right)

	totalDist := 0
	for i := 0; i < len(left); i++ {
		totalDist += int(math.Abs(float64(left[i] - right[i])))
	}

	return totalDist
}

func part2() int {
	left := make([]int, 0, 1000)
	right := make(map[int]int)

	f, err := os.Open("./2024/1/input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		nums := strings.Fields(line)

		lNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		left = append(left, lNum)

		rNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		right[rNum] += 1
	}

	score := 0
	for i := 0; i < len(left); i++ {
		score += left[i] * right[left[i]]
	}

	return score
}
