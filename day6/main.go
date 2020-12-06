package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"../util"
)

var questionsAnswered [26]int

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read input.txt")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(util.SplitOnBlankLine)

	questionsAnsweredByAnyGroupMember := 0
	questionsAnsweredYesByEntireGroup := 0
	for ; scanner.Scan(); resetQuestionsAnswered() {
		groupStr := scanner.Bytes()
		memberCount := 0

		for _, c := range groupStr {
			if c == '\n' {
				memberCount++
			} else if unicode.IsLetter(rune(c)) {
				questionsAnswered[c-'a']++
			}
		}

		for _, n := range questionsAnswered {
			if n >= 1 {
				questionsAnsweredByAnyGroupMember++
			}
			if n == memberCount {
				questionsAnsweredYesByEntireGroup++
			}
		}
	}

	fmt.Println("Part One:", questionsAnsweredByAnyGroupMember, "questions answered \"yes\" for each group")
	fmt.Println("Part Two:", questionsAnsweredYesByEntireGroup, "questions answered \"yes\" by everyone in each group")
}

func resetQuestionsAnswered() {
	for i := 0; i < 26; i++ {
		questionsAnswered[i] = 0
	}
}
