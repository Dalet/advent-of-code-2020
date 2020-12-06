package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"

	"../util"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't read input.txt")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(util.SplitOnBlankLine)

	validPassportsCountPartOne := 0
	validPassportsCountPartTwo := 0
	for scanner.Scan() {
		hasRequiredFields, hasValidRequiredFields := isValidPassport(scanner.Bytes())
		if hasRequiredFields {
			validPassportsCountPartOne++
			if hasValidRequiredFields {
				validPassportsCountPartTwo++
			}
		}
	}

	fmt.Println("Part One:", validPassportsCountPartOne, "valid passports")
	fmt.Println("Part Two:", validPassportsCountPartTwo, "valid passports")
}

func isValidPassport(passportStr []byte) (hasRequiredFields bool, hasValidRequiredFields bool) {
	fields := bytes.Fields(passportStr)

	hasValidRequiredFields = true
	missingFields := getRequiredPassportFields()
	for _, keyValue := range fields {
		i := bytes.IndexByte(keyValue, ':')
		field := string(keyValue[:i])
		value := keyValue[i+1:]

		validate, found := missingFields[field]
		if !found {
			continue
		} else {
			delete(missingFields, field)
		}

		if !validate(value) {
			hasValidRequiredFields = false
		}
	}

	return len(missingFields) == 0, hasValidRequiredFields
}

type validator func([]byte) bool

func getRequiredPassportFields() map[string]validator {
	return map[string]validator{
		"byr": func(value []byte) bool {
			return len(value) == 4 && isInRange(value, 1920, 2002)
		},
		"iyr": func(value []byte) bool {
			return len(value) == 4 && isInRange(value, 2010, 2020)
		},
		"eyr": func(value []byte) bool {
			return len(value) == 4 && isInRange(value, 2020, 2030)
		},
		"hgt": func(value []byte) bool {
			firstNonDigitIndex := bytes.IndexFunc(value, func(r rune) bool {
				return !unicode.IsDigit(r)
			})
			if firstNonDigitIndex <= 0 {
				return false
			}

			height := value[:firstNonDigitIndex]
			heightUnit := value[firstNonDigitIndex:]

			switch {
			case bytes.Equal(heightUnit, []byte("in")):
				return isInRange(height, 59, 76)
			case bytes.Equal(heightUnit, []byte("cm")):
				return isInRange(height, 150, 193)
			default:
				return false
			}
		},
		"hcl": func(value []byte) bool {
			return hclRegex.Match(value)
		},
		"ecl": func(value []byte) bool {
			_, found := allowedEcl[string(value)]
			return found
		},
		"pid": func(value []byte) bool {
			if len(value) != 9 {
				return false
			}

			nonDigitIndex := bytes.IndexFunc(value, func(r rune) bool {
				return !unicode.IsDigit(r)
			})
			return nonDigitIndex == -1
		},
	}
}

func isInRange(value []byte, min int, max int) bool {
	n, err := strconv.Atoi(string(value))
	return err == nil && n >= min && n <= max
}

var hclRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)

var allowedEcl = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}
