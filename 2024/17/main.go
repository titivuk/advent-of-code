package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

var (
	instructionPointer = 0
	A, B, C            uint64
	operands           = map[byte]func() uint64{
		'0': func() uint64 { return 0 },
		'1': func() uint64 { return 1 },
		'2': func() uint64 { return 2 },
		'3': func() uint64 { return 3 },
		'4': func() uint64 { return A },
		'5': func() uint64 { return B },
		'6': func() uint64 { return C },
	}
	instructions = map[byte]func(operand byte){
		'0': func(operand byte) {
			opVal := operands[operand]()
			A /= uint64(math.Pow(2, float64(opVal)))
		},
		'1': func(operand byte) {
			literalOperand, err := strconv.Atoi(string(operand))
			if err != nil {
				panic(err)
			}

			B = B ^ uint64(literalOperand)
		},
		'2': func(operand byte) {
			opVal := operands[operand]()
			B = opVal % 8
		},
		'3': func(operand byte) {
			if A == 0 {
				return
			}

			literalOperand, err := strconv.Atoi(string(operand))
			if err != nil {
				panic(err)
			}

			instructionPointer = literalOperand
		},
		'4': func(operand byte) {
			B = B ^ C
		},
		'5': func(operand byte) {
			opVal := operands[operand]()
			output = append(output, fmt.Sprint(opVal%8))
		},
		'6': func(operand byte) {
			opVal := operands[operand]()
			B = A / uint64(math.Pow(2, float64(opVal)))
		},
		'7': func(operand byte) {
			opVal := operands[operand]()
			C = A / uint64(math.Pow(2, float64(opVal)))
		},
	}
	output []string
	minA   uint64 = math.MaxUint64
)

func main() {
	// ex1 := "2,6"
	// C = 9
	// testPart1(ex1, "ex1")

	// ex2 := "5,0,5,1,5,4"
	// A = 10
	// testPart1(ex2, "ex2")

	// ex3 := "0,1,5,4,3,0"
	// A = 2024
	// testPart1(ex3, "ex3")

	ex4 := "0,3,5,4,3,0"
	A = 229
	testPart1(ex4, "ex4")

	input := "2,4,1,1,7,5,1,5,4,0,5,5,0,3,3,0"
	A = 64854237
	B = 0
	C = 0
	testPart1(input, "input")
	fmt.Println("input p2", part2(input))
}

func testPart1(program, name string) {
	fmt.Printf("%s p1: \"%s\"\n", name, string(part1(program)))
}

func part1(program string) []byte {
	defer func() {
		instructionPointer = 0
		A = 0
		B = 0
		C = 0
		output = nil
	}()

	for instructionPointer < len(program)-2 {
		opcode := program[instructionPointer]
		operand := program[instructionPointer+2]

		ip := instructionPointer
		instructions[opcode](operand)

		if opcode != '3' || ip == instructionPointer {
			instructionPointer += 4
		}
	}
	return []byte(strings.Join(output, ""))
}

func part2(program string) uint64 {
	defer func() {
		instructionPointer = 0
		A = 0
		B = 0
		C = 0
		output = nil
	}()

	rp := []byte(program)
	slices.Reverse(rp)

	dfs(0, 0, program)

	return minA
}

func dfs(a uint64, depth int, program string) uint64 {
	expectedOutput := []byte(strings.Join(strings.Split(program, ","), ""))

	for i := 0; i < 8; i += 1 {
		nextA := a*8 + uint64(i)
		A = nextA
		output := part1(program)

		if !slices.Equal(output, expectedOutput[len(expectedOutput)-1-depth:]) {
			continue
		}

		if depth == len(expectedOutput)-1 {
			if minA > nextA {
				minA = nextA
			}
			break
		}

		dfs(nextA, depth+1, program)
	}

	return 0
}
