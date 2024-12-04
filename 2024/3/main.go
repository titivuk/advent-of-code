package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	mul  = "mul"
	do   = "do()"
	dont = "don't()"
)

func main() {
	f, err := os.Open("./2024/3/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("part 1", part1(string(input)))
	fmt.Println("part 2", part2(string(input)))
}

func part1(input string) int {
	i := 0
	sum := 0
	for i < len(input) {
		if input[i] == 'm' {
			// mul(x,y) - min required length = 8
			if i+8 > len(input) {
				break
			}

			mulNext := expectSequence(input, i, mul)
			if mulNext {
				i += len(mul)
			} else {
				i += 1
				continue
			}

			lbraceNext := expectChar(input, i, '(')
			i += 1
			if !lbraceNext {
				continue
			}

			if !isDigit(input[i]) {
				i += 1
				continue
			}

			// parse left number
			l := i
			for i < len(input) && isDigit(input[i]) {
				i += 1
			}
			lnum, err := strconv.Atoi(input[l:i])
			if err != nil {
				panic(err)
			}

			commaNext := expectChar(input, i, ',')
			i += 1
			if !commaNext {
				continue
			}

			// parse right number
			l = i
			for i < len(input) && isDigit(input[i]) {
				i += 1
			}
			rnum, err := strconv.Atoi(input[l:i])
			if err != nil {
				panic(err)
			}

			if !expectChar(input, i, ')') {
				i += 1
				continue
			}

			sum += lnum * rnum
		}

		i++
	}

	return sum
}

func part2(input string) int {
	i := 0
	sum := 0
	enabled := true
	for i < len(input) {
		switch input[i] {
		case 'm':
			if !enabled {
				i += 1
				continue
			}

			mulNext := expectSequence(input, i, mul)
			if mulNext {
				i += 3
			} else {
				i += 1
				continue
			}

			lbraceNext := expectChar(input, i, '(')
			i += 1
			if !lbraceNext {
				continue
			}

			if !isDigit(input[i]) {
				i += 1
				continue
			}

			// parse left number
			l := i
			for i < len(input) && isDigit(input[i]) {
				i += 1
			}
			lnum, err := strconv.Atoi(input[l:i])
			if err != nil {
				panic(err)
			}

			commaNext := expectChar(input, i, ',')
			i += 1
			if !commaNext {
				continue
			}

			// parse right number
			l = i
			for i < len(input) && isDigit(input[i]) {
				i += 1
			}
			rnum, err := strconv.Atoi(input[l:i])
			if err != nil {
				panic(err)
			}

			if !expectChar(input, i, ')') {
				i += 1
				continue
			}

			sum += lnum * rnum

			i += 1
		case 'd':
			if expectSequence(input, i, do) {
				enabled = true
				i += len(do)
			} else if expectSequence(input, i, dont) {
				enabled = false
				i += len(dont)
			} else {
				i += 1
			}
		default:
			i += 1
		}
	}

	return sum
}

func isDigit(ch byte) bool {
	return 48 <= ch && ch <= 57
}

func expectSequence(input string, start int, seq string) bool {
	if len(input) <= start+len(seq)-1 {
		return false
	}

	for i := 0; i < len(seq); i++ {
		if seq[i] != input[start+i] {
			return false
		}
	}

	return true
}

func expectChar(input string, i int, ch byte) bool {
	if len(input) <= i {
		return false
	}

	return input[i] == ch
}
