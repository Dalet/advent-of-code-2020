package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read input.txt")
		return
	}
	defer file.Close()

	input := make([][]byte, 0)
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		lineBuffer := scanner.Bytes()
		lineCopy := make([]byte, len(lineBuffer))
		copy(lineCopy, lineBuffer)
		input = append(input, lineCopy)
	}

	fmt.Println("Part One:", countTrees(input, 3, 1), "trees encountered")

	partTwoAnswer := countTrees(input, 1, 1) *
		countTrees(input, 3, 1) *
		countTrees(input, 5, 1) *
		countTrees(input, 7, 1) *
		countTrees(input, 1, 2)
	fmt.Println("Part Two: product of trees encountered:", partTwoAnswer)
}

func countTrees(input [][]byte, slopeX int, slopeY int) int {
	treeCount := 0
	x, y := slopeX, slopeY

	for y < len(input) {
		char := input[y][x]
		if char == '#' {
			treeCount++
		}

		x = (x + slopeX) % len(input[y])
		y += slopeY
	}

	return treeCount
}
