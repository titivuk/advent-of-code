package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./2024/2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	count1 := 0
	count2 := 0
	for scanner.Scan() {
		report := scanner.Text()
		isSafe1 := isSafeReportPart1(report)
		isSafe2 := isSafeReportPart2(report)
		if isSafe1 {
			count1++
		}
		if isSafe2 {
			count2++
		}
	}

	fmt.Println("p1", count1)
	fmt.Println("p2", count2)
}

func isSafeReportPart1(report string) bool {
	lvls := strings.Split(report, " ")
	if len(lvls) == 1 {
		return true
	}

	return checkLvls(lvls)
}

func isSafeReportPart2(report string) bool {
	lvls := strings.Split(report, " ")
	if len(lvls) == 1 {
		return true
	}

	if checkLvls(lvls) {
		return true
	}

	for i := 0; i < len(lvls); i++ {
		lvlsVariant := make([]string, 0, len(lvls)-1)
		lvlsVariant = append(lvlsVariant, lvls[:i]...)
		lvlsVariant = append(lvlsVariant, lvls[i+1:]...)

		if len(lvlsVariant) == 1 {
			return true
		}

		if checkLvls(lvlsVariant) {
			return true
		}
	}

	return false
}

func checkLvls(lvls []string) bool {
	prevDiff := 0
	for i := 1; i < len(lvls); i++ {
		prev, err := strconv.Atoi(lvls[i-1])
		if err != nil {
			panic(err)
		}
		cur, err := strconv.Atoi(lvls[i])
		if err != nil {
			panic(err)
		}

		diff := cur - prev
		if diff < 0 && prevDiff > 0 || diff > 0 && prevDiff < 0 {

			return false
		}

		absDiff := math.Abs(float64(diff))
		if absDiff < 1 || 3 < absDiff {
			return false
		}

		prevDiff = diff
	}

	return true
}
