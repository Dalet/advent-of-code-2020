package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"

	"../util"
)

var questionsAnswered [26]bool

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read input.txt")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(util.SplitOnBlankLine)

	totalQuestionsAnswered := 0
	for ; scanner.Scan(); resetQuestionsAnswered() {
		group := scanner.Bytes()

		for _, c := range group {
			if unicode.IsLetter(rune(c)) && !questionsAnswered[c-'a'] {
				questionsAnswered[c-'a'] = true
				totalQuestionsAnswered++
			}
		}
	}

	fmt.Println("Part One:", totalQuestionsAnswered, "questions answered \"yes\" for each group")
}

func resetQuestionsAnswered() {
	for i := 0; i < 26; i++ {
		questionsAnswered[i] = false
	}
}
