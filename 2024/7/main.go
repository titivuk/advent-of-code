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
	f, err := os.Open("./2024/7/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	equations := make(map[int][]int)
	for _, l := range strings.Split(string(input), "\n") {
		parts := strings.SplitN(l, ": ", 2)

		testValue, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		numsStr := strings.Split(parts[1], " ")
		nums := make([]int, len(numsStr))
		for i, ns := range numsStr {
			n, err := strconv.Atoi(ns)
			if err != nil {
				panic(err)
			}
			nums[i] = n
		}

		equations[testValue] = nums
	}

	fmt.Println("part1", part12(equations, false))
	fmt.Println("part2", part12(equations, true))
}

func part12(equations map[int][]int, concat bool) int {
	var solveEq func(result, curResult int, nums []int, i int, concat bool) bool

	solveEq = func(testValue, curValue int, nums []int, i int, concat bool) bool {
		if i >= len(nums) {
			return testValue == curValue
		}

		if curValue > testValue {
			return false
		}

		return solveEq(testValue, curValue+nums[i], nums, i+1, concat) ||
			solveEq(testValue, curValue*nums[i], nums, i+1, concat) ||
			(concat && solveEq(testValue, concatInt(curValue, nums[i]), nums, i+1, concat))
	}

	sum := 0
	for testValue, nums := range equations {
		if solveEq(testValue, nums[0], nums, 1, concat) {
			sum += testValue
		}
	}

	return sum
}

func concatInt(a, b int) int {
	bLog10 := int(math.Log10(float64(b)))
	result := a*int(math.Pow10(bLog10+1)) + b

	return result
}
