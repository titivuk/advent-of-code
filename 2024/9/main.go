package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./2024/9/input.txt")
	if err != nil {
		panic(err)
	}

	fBuf, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1", part1(parseDiskMap(string(fBuf))))
	fmt.Println("part2", part2(parseDiskMap(string(fBuf))))
}

func part1(diskMap []int) int {
	// two pointers technique
	l := 1
	r := len(diskMap) - 1

	for l < r {
		if diskMap[l] != -1 {
			l += 1
			continue
		}

		if diskMap[r] == -1 {
			r -= 1
			continue
		}

		diskMap[l] = diskMap[r]
		diskMap[r] = -1

		l += 1
		r -= 1
	}

	i := 0
	checksum := 0
	for diskMap[i] != -1 {
		checksum += i * diskMap[i]
		i += 1
	}

	return checksum
}

func part2(diskMap []int) int {
	// file block id
	fId := len(diskMap) - 1
	for fId >= 0 {
		// skip empty spaces
		if diskMap[fId] == -1 {
			fId -= 1
			continue
		}

		// calculate file len
		fileLen := 0
		for fId-fileLen >= 0 && diskMap[fId-fileLen] == diskMap[fId] {
			fileLen += 1
		}

		// space id
		spId := 0
		for spId < fId {
			// skip files
			if diskMap[spId] != -1 {
				spId += 1
				continue
			}

			// calculate free space
			freeSpace := 0
			for spId+freeSpace < fId && diskMap[spId+freeSpace] == -1 {
				freeSpace += 1
			}

			if fileLen <= freeSpace {
				// move file blocks to the left
				for j := 0; j < fileLen; j++ {
					diskMap[spId+j] = diskMap[fId-fileLen+j+1]
					diskMap[fId-fileLen+j+1] = -1
				}

				break
			}

			// check next free space
			spId += freeSpace
		}

		fId -= fileLen
	}

	checksum := 0
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] != -1 {
			checksum += i * diskMap[i]
		}
	}

	return checksum
}

func parseDiskMap(diskMapStr string) []int {
	diskMap := make([]int, 0)

	fileId := 0
	i := 0
	for i < len(diskMapStr)-1 {
		fileLen, err := strconv.Atoi(string(diskMapStr[i]))
		if err != nil {
			panic(err)
		}

		for j := 0; j < fileLen; j++ {
			diskMap = append(diskMap, fileId)
		}
		i += 1

		space, err := strconv.Atoi(string(diskMapStr[i]))
		if err != nil {
			panic(err)
		}

		for j := 0; j < space; j++ {
			diskMap = append(diskMap, -1)
		}
		i += 1
		fileId += 1
	}

	if len(diskMapStr)%2 == 1 {
		fileLen, err := strconv.Atoi(string(diskMapStr[len(diskMapStr)-1]))
		if err != nil {
			panic(err)
		}

		for j := 0; j < fileLen; j++ {
			diskMap = append(diskMap, fileId)
		}
		i += 1
	}

	return diskMap
}
