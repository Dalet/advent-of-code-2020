package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	checkErr(err)
	defer file.Close()

	inputNumbers := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		checkErr(err)

		inputNumbers = append(inputNumbers, num)
	}

	printResult := func(howManyNumbers int) {
		foundNumbers, _ := findNumbersAddingUpTo(2020, inputNumbers, howManyNumbers)
		fmt.Println("The numbers", foundNumbers, "sum to 2020. Multiplied all together:", multiply(foundNumbers))
	}

	fmt.Println("PART ONE:")
	printResult(2)

	fmt.Println("PART TWO:")
	printResult(3)
}

func findNumbersAddingUpTo(result int, numbers []int, howManyNumbers int) (addingNumbers []int, found bool) {
	if howManyNumbers == 1 {
		for _, n := range numbers {
			if n == result {
				return []int{n}, true
			}
		}
		return nil, false
	}

	for _, n := range numbers {
		foundNumbers, found := findNumbersAddingUpTo(result-n, numbers, howManyNumbers-1)
		if found {
			return append(foundNumbers, n), true
		}
	}

	return nil, false
}

func multiply(nums []int) (result int) {
	result = nums[0]
	for _, num := range nums[1:] {
		result *= num
	}
	return
}

func checkErr(e interface{}) {
	if e != nil {
		panic(e)
	}
}
