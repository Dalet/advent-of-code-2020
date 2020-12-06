package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read input.txt")
		return
	}
	defer file.Close()

	actualIDSum := 0
	minSeatID, maxSeatID := -1, -1
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		code := scanner.Bytes()

		row := calcCode(code[:8], 0, 127)
		col := calcCode(code[7:], 0, 7)
		seatID := row*8 + col

		if seatID > maxSeatID {
			maxSeatID = seatID
		}

		if minSeatID == -1 || seatID < minSeatID {
			minSeatID = seatID
		}

		actualIDSum += seatID
	}

	fmt.Println("Part One: highest seat id:", maxSeatID)

	// sum of minSeatID->maxSeatID  = "sum of 0->maxSeatID" minus "sum of 0->(minSeatID-1)"
	zeroToMaxSeatIDSum := maxSeatID * (maxSeatID + 1) / 2
	zeroToMinSeatIDSum := (minSeatID - 1) * minSeatID / 2
	expectedIDSum := zeroToMaxSeatIDSum - zeroToMinSeatIDSum

	mySeatID := expectedIDSum - actualIDSum
	fmt.Println("Part Two: my seat id:", mySeatID)
}

func calcCode(code []byte, start int, end int) int {
	for _, c := range code {
		middle := float64(start+end) / 2.0
		switch c {
		case 'F', 'L':
			end = int(math.Floor(middle))
		case 'B', 'R':
			start = int(math.Ceil(middle))
		}
	}

	return end
}
