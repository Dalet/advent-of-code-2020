package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type password struct {
	password string
	policy   passwordPolicy
}

type passwordPolicy struct {
	char      byte
	firstInt  int
	secondInt int
}

func main() {
	file, err := os.Open("input.txt")
	checkErr(err)
	defer file.Close()

	passwords := make([]password, 0)
	regex := regexp.MustCompile(`^(\d+)-(\d+) (\S): (.+)$`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// parse input
		match := regex.FindStringSubmatch(line)
		min, err := strconv.Atoi(match[1])
		checkErr(err)
		max, err := strconv.Atoi(match[2])
		checkErr(err)
		char := match[3][0]
		passwordStr := match[4]

		passwords = append(passwords, password{
			password: passwordStr,
			policy: passwordPolicy{
				firstInt:  min,
				secondInt: max,
				char:      char,
			},
		})
	}

	validPasswordsPartOneCount := 0
	validPasswordsPartTwoCount := 0
	for _, password := range passwords {
		if isValidPasswordPartOne(password) {
			validPasswordsPartOneCount++
		}

		if isValidPasswordPartTwo(password) {
			validPasswordsPartTwoCount++
		}
	}

	fmt.Println("Valid passwords count for part one:", validPasswordsPartOneCount)
	fmt.Println("Valid passwords count for part two:", validPasswordsPartTwoCount)
}

func isValidPasswordPartOne(password password) bool {
	charCount := 0
	for _, char := range password.password {
		if byte(char) == password.policy.char {
			charCount++
		}
	}

	return charCount >= password.policy.firstInt && charCount <= password.policy.secondInt
}

func isValidPasswordPartTwo(password password) bool {
	firstChar := password.password[password.policy.firstInt-1]
	secondChar := password.password[password.policy.secondInt-1]

	firstCharMatch := firstChar == password.policy.char
	secondCharMatch := secondChar == password.policy.char
	return (firstCharMatch || secondCharMatch) && firstCharMatch != secondCharMatch
}

func checkErr(e interface{}) {
	if e != nil {
		panic(e)
	}
}
